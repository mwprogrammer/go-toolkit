package models

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Mail struct {
	Sender      string
	Receipients []string
	Subject     string
	Body        string
	Attachments []string
}