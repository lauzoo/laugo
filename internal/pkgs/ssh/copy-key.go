package ssh

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"
	"path"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type Tools struct {
	tunnelConfig *SshTunnelConfig
}

func (t *Tools) CopyKey(host string, port int, username string, password string) {
	session, err := t.createSession(username, password, host, port)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	var cmd string
	var sshKey = getSSHKey()
	cmd = fmt.Sprintf("echo \"%s\"  >> ~/.ssh/authorized_keys", sshKey)
	session.Run(cmd)
}

func getSSHKey() (key string) {
	var err error
	var usr *user.User
	if usr, err = user.Current(); err != nil {
		panic(err)
	}

	var sshKeyPath string
	var sshKeyContent []byte
	sshKeyPath = path.Join(usr.HomeDir, ".ssh/id_rsa.pub")
	if sshKeyContent, err = ioutil.ReadFile(sshKeyPath); err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(sshKeyContent))
}

func (t *Tools) createSession(user, password, host string, port int) (session *ssh.Session, err error) {
	var client *ssh.Client
	client, err = t.connect(user, password, host, port)
	if err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}
func (t *Tools) connect(user, password, host string, port int) (client *ssh.Client, err error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			// todo(@liuliqiang): save host key to local?
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	return ssh.Dial("tcp", addr, clientConfig)
}
