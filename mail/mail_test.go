package mail

import (
	"testing"
)

func TestSendMessage(t *testing.T) {

	sender := GetSender()
	sender.Configure("host", "username", "password", "port")

	message := CreateMessage()
	message.AddTo("example@gmail.com")
	message.AddSubject("TEST SENDING MESSAGE")
	message.AddBody("Test sending message.")

	err := sender.Send(message)

	if err != nil {
		t.Errorf("Message not sent, error encountered: %s", err.Error())
	}

}

func TestSendMessageWithCC(t *testing.T) {

	sender := GetSender()
	sender.Configure("host", "username", "password", "port")

	message := CreateMessage()
	message.AddTo("example@gmail.com")
	message.AddCC("example@outlook.com")
	message.AddSubject("TEST SENDING MESSAGE WITH CC")
	message.AddBody("Test sending message with cc.")

	err := sender.Send(message)

	if err != nil {
		t.Errorf("Message not sent, error encountered: %s", err.Error())
	}

}

func TestSendMessageWithBCC(t *testing.T) {

	sender := GetSender()
	sender.Configure("host", "username", "password", "port")

	message := CreateMessage()
	message.AddTo("example@gmail.com")
	message.AddCC("example@outlook.com")
	message.AddBCC("example@proton.me")
	message.AddSubject("TEST SENDING MESSAGE WITH BCC")
	message.AddBody("Test sending message with bcc.")

	err := sender.Send(message)

	if err != nil {
		t.Errorf("Message not sent, error encountered: %s", err.Error())
	}
}

func TestSendMessageWithAttachments(t *testing.T) {

	sender := GetSender()
	sender.Configure("host", "username", "password", "port")

	message := CreateMessage()
	message.AddTo("example@gmail.com")
	message.AddSubject("TEST SIMPLE MESSAGE WITH ATTACHMENTS")
	message.AddBody("Test sending message with attachments.")

	err := message.AttachFile("..path/to/file")

	if err != nil {
		t.Errorf("Attachment not added, error encountered: %s", err.Error())
	}

	err = sender.Send(message)

	if err != nil {
		t.Errorf("Message not sent, error encountered: %s", err.Error())
	}
}
