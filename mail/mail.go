// Package mail allows you to easily send emails in your go applications.
// It is built using go's net/smtp package.
//
// Author [Chisomo Chiweza]: https://github.com/mwprogrammer
package mail

// Returns the sender being used to send email.
// Sender will need to be configured before sending messages and can be configured
// with Sender.Configure()
func GetSender() *Sender {
	return &Sender{}
}

// Returns an instance of a message object which can be used to construct emails.
func CreateMessage() *Message {
	return &Message{Attachments: make(map[string][]byte)}
}
