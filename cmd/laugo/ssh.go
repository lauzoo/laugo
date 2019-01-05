package main

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

	. "github.com/lauzoo/laugo/internal/pkgs/log"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

var (
	host     string
	port     int
	username string
	password string
)

func init() {
	sshCmd.PersistentFlags().StringVarP(&host, "host", "H", "127.0.0.1", "Host to login")
	sshCmd.PersistentFlags().IntVarP(&port, "port", "P", 22, "Port for host")
	sshCmd.PersistentFlags().StringVarP(&username, "username", "u", "root", "User for host")
	sshCmd.PersistentFlags().StringVarP(&password, "password", "p", "password", "Password for username")
	rootCmd.AddCommand(sshCmd)
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh utils",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			Log.Error("invalid args count, user -help for more info.")
			return
		}

		switch strings.ToUpper(args[0])[:7] {
		case "COPY-ID":
			sshCopyId(host, port, username, password)
		default:
			Log.Error("unsupport operation: " + args[0])
			return
		}
	},
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

func sshCopyId(host string, port int, username string, password string) {
	session, err := connect(username, password, host, port)
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

func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
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

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}
