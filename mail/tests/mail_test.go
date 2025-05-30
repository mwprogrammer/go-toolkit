package tests

import (
	"testing"

	"github.com/mwprogrammer/go-utilities/mail"
)

func TestSendMessage(t *testing.T) {

	sender := mail.GetSender()
	sender.Configure("smtp.office365.com", "no.reply@nbs.mw", "Reports@21", "587")

	message := mail.CreateMessage()
	message.AddTo("chisomo.chiweza@nbs.mw")
	message.AddSubject("TEST SENDING MESSAGE")
	message.AddBody("Test sending message.")

	err := sender.Send(message)

	t.Errorf("Message not sent, error encountered: %s", err.Error())
}

func TestSendMessageWithCC(t *testing.T) {

	sender := mail.GetSender()
	sender.Configure("smtp.office365.com", "no.reply@nbs.mw", "Reports@21", "587")

	message := mail.CreateMessage()
	message.AddTo("chisomo.chiweza@nbs.mw")
	message.AddCC("liana.chagunda@nbs.mw")
	message.AddSubject("TEST SENDING MESSAGE WITH CC")
	message.AddBody("Test sending message with cc.")

	err := sender.Send(message)

	t.Errorf("Message not sent, error encountered: %s", err.Error())
}

func TestSendMessageWithBCC(t *testing.T) {

	sender := mail.GetSender()
	sender.Configure("smtp.office365.com", "no.reply@nbs.mw", "Reports@21", "587")

	message := mail.CreateMessage()
	message.AddTo("chisomo.chiweza@nbs.mw")
	message.AddCC("liana.chagunda@nbs.mw")
	message.AddBCC("ekariorama.magaleta@nbs.mw")
	message.AddSubject("TEST SENDING MESSAGE WITH BCC")
	message.AddBody("Test sending message with bcc.")

	err := sender.Send(message)

	t.Errorf("Message not sent, error encountered: %s", err.Error())
}

func TestSendMessageWithAttachments(t *testing.T) {

	sender := mail.GetSender()
	sender.Configure("smtp.office365.com", "no.reply@nbs.mw", "Reports@21", "587")

	message := mail.CreateMessage()
	message.AddTo("chisomo.chiweza@nbs.mw")
	message.AddSubject("TEST SIMPLE MESSAGE WITH ATTACHMENTS")
	message.AddBody("Test sending message with attachments.")

	err := message.AttachFile("tests/attachments/2025_Computer_Science.pdf")
	t.Errorf("Attachment not added, error encountered: %s", err.Error())

	sender.Send(message)

	err = sender.Send(message)

	t.Errorf("Message not sent, error encountered: %s", err.Error())
}
