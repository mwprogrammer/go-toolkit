package mail

import (
	"log/slog"

	"github.com/mwprogrammer/go-utilities/models"
	"gopkg.in/gomail.v2"
)

var dialer *gomail.Dialer

func Configure(config models.MailConfig) {

	dialer = gomail.NewDialer(
		config.Host, 
		config.Port, 
		config.Username, 
		config.Password)

}

func Send(mail models.Mail, logger *slog.Logger) bool {

	is_success := true

	message := gomail.NewMessage()

	message.SetHeader("From", mail.Sender)
	message.SetHeader("To", mail.Receipients...)
	message.SetHeader("Subject", mail.Subject)
	message.SetBody("text/html", mail.Body)

	for _, attachment := range mail.Attachments {
		message.Attach(attachment)
	}
	

	err := dialer.DialAndSend(message)

	if err != nil {
		logger.Error(err.Error())
		is_success = false
	}

	return is_success

}