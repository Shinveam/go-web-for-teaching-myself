package ssh_connector

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
)

func Connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth    []ssh.AuthMethod
		addr    string
		client  *ssh.Client
		session *ssh.Session
		err     error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	client, err = ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

