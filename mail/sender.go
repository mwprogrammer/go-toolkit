package mail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
)

type settings struct {
	host     string
	username string
	password string
	port     string
}

var config *settings

type Sender struct {
	auth smtp.Auth
}

func (s *Sender) Configure(host, username, password, port string) {

	config = &settings{
		host:     host,
		username: username,
		password: password,
		port:     port,
	}

	s.auth = &Auth{username, password}

}

func (s *Sender) Send(email *Message) error {

	if config == nil {
		return errors.New("sender is not configured")
	}

	tlsConfig := tls.Config{
		ServerName: config.host,
	}

	connection, err := net.Dial("tcp", net.JoinHostPort(config.host, config.port))

	if err != nil {
		return err
	}

	client, err := smtp.NewClient(connection, config.host)

	if err != nil {
		fmt.Println("Could not create client.")
		return err
	}

	err = client.Auth(s.auth)

	if err != nil {
		fmt.Println("Could not configure authentication.")
		return err
	}

	err = client.StartTLS(&tlsConfig)

	if err != nil {
		fmt.Println("Could not start TLS")
		return err
	}

	return smtp.SendMail(fmt.Sprintf("%s:%s", config.host, config.port), s.auth, config.username, email.to, email.ToBytes())
}
