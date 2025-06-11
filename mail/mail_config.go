package mail

import (
	"errors"
	"net/smtp"
)

// A private object to keep the sender's settings.
type settings struct {
	host     string
	username string
	password string
	port     string
}

var config *settings

// Auth implements the Auth interface and allows for the sender to use
// LoginAuth rather than PlainAuth which is the default in smtp.
//
// Email providers like outlook do not work with PlainAuth.
type Auth struct {
	username, password string
}

func (a *Auth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *Auth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}
