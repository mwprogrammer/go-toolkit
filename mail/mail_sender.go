package mail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
)

// Represents an entity that connects to the specified email
// server and is used to send emails.
type Sender struct {
	auth smtp.Auth
}

// Configure the email server's details.
func (s *Sender) Configure(host, username, password, port string) {

	config = &settings{
		host:     host,
		username: username,
		password: password,
		port:     port,
	}

	s.auth = &Auth{username, password}

}

// Sends an email. Only servers that support TLS can use this.
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

	defer connection.Close()

	client, err := smtp.NewClient(connection, config.host)

	if err != nil {
		fmt.Println("Could not create client.")
		return err
	}

	defer client.Close()

	err = client.StartTLS(&tlsConfig)

	if err != nil {
		fmt.Println("Could not start TLS")
		return err
	}

	err = client.Auth(s.auth)

	if err != nil {
		fmt.Println("Could not configure authentication.")
		return err
	}

	err = client.Mail(config.username)

	if err != nil {
		fmt.Println("Could not set sender address:", err)
		return err
	}

	content, recipients := email.ToBytes()

	for _, recipient := range recipients {

		err = client.Rcpt(recipient)

		if err != nil {
			fmt.Println("Could not set recipient address:", err, recipient)
			return err
		}

	}

	wc, err := client.Data()

	if err != nil {
		fmt.Println("Could not start data transfer:", err)
		return err
	}

	defer wc.Close()

	_, err = wc.Write(content)

	if err != nil {
		fmt.Println("Could not write email body:", err)
		return err
	}

	client.Quit()

	return nil
}
