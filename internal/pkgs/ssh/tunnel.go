package ssh

import (
	"fmt"
	"io"
	"net"

	"golang.org/x/crypto/ssh"
)

func (t *Tools) CreateTunnel(config *SshTunnelConfig) error {
	t.tunnelConfig = config
	listener, err := net.Listen("tcp", config.Local.String())
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go t.forward(conn)
	}
}

func (t *Tools) forward(localConn net.Conn) {
	var (
		remoteClient *ssh.Client
		err          error
	)
	remoteClient, err = t.connect(
		t.tunnelConfig.Remote.User,
		t.tunnelConfig.Remote.Pass,
		t.tunnelConfig.Remote.Host,
		t.tunnelConfig.Remote.Port)
	if err != nil {
		fmt.Printf("Server dial error: %s\n", err)
		return
	}

	remoteConn, err := remoteClient.Dial("tcp", t.tunnelConfig.Target.String())
	if err != nil {
		fmt.Printf("Remote dial error: %s\n", err)
		return
	}

	copyConn := func(writer, reader net.Conn) {
		_, err := io.Copy(writer, reader)
		if err != nil {
			fmt.Printf("io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}
