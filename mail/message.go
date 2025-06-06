package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

// Represents the content that is then sent via email.
type Message struct {
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	Attachments map[string][]byte
}

// Adds a main receipient.
func (m *Message) AddTo(recipient string) {
	m.to = append(m.to, recipient)
}

// Adds a cc receipient.
func (m *Message) AddCC(cc_recipient string) {
	m.cc = append(m.cc, cc_recipient)
}

// Add a bcc receipient.
func (m *Message) AddBCC(bcc_recipient string) {
	m.bcc = append(m.bcc, bcc_recipient)
}

// Adds a subject.
func (m *Message) AddSubject(subject string) {
	m.subject = subject
}

// Adds a body.
func (m *Message) AddBody(body string) {
	m.body = body
}

// Attaches a file to the message. Currently supports pdfs,
// jpgs and pngs. More formats to be added and tested soon.
func (m *Message) AttachFile(path string) error {

	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	_, filename := filepath.Split(path)

	m.Attachments[filename] = file

	return nil
}

// Constructs the message before transforming it into an array of bytes as
// a pre-processing step before sending through the smtp client.
func (m *Message) ToBytes() ([]byte, []string) {

	buffer := bytes.NewBuffer(nil)

	header := make(textproto.MIMEHeader)
	header.Set("Subject", m.subject)
	header.Set("To", strings.Join(m.to, ", "))
	if len(m.cc) > 0 {
		header.Set("Cc", strings.Join(m.cc, ", "))
	}
	header.Set("MIME-Version", "1.0")

	allRecipients := make([]string, 0, len(m.to)+len(m.cc)+len(m.bcc))
	allRecipients = append(allRecipients, m.to...)
	allRecipients = append(allRecipients, m.cc...)
	allRecipients = append(allRecipients, m.bcc...)

	if len(m.Attachments) > 0 {

		multipartWriter := multipart.NewWriter(buffer)
		header.Set("Content-Type", fmt.Sprintf("multipart/mixed; boundary=%s", multipartWriter.Boundary()))

		for k, v := range header {
			buffer.WriteString(fmt.Sprintf("%s: %s\n", k, strings.Join(v, ", ")))
		}
		buffer.WriteString("\n")

		textPartHeaders := make(textproto.MIMEHeader)
		textPartHeaders.Set("Content-Type", "text/plain; charset=utf-8")

		bodyPart, err := multipartWriter.CreatePart(textPartHeaders)
		if err != nil {
			fmt.Println("Error creating body part:", err)
			return nil, nil
		}
		bodyPart.Write([]byte(m.body))

		for filename, data := range m.Attachments {

			attachmentPartHeaders := make(textproto.MIMEHeader)
			contentType := "application/octet-stream"

			if strings.HasSuffix(filename, ".txt") {
				contentType = "text/plain"
			} else if strings.HasSuffix(filename, ".pdf") {
				contentType = "application/pdf"
			} else if strings.HasSuffix(filename, ".jpg") || strings.HasSuffix(filename, ".jpeg") {
				contentType = "image/jpeg"
			} else if strings.HasSuffix(filename, ".png") {
				contentType = "image/png"
			}

			attachmentPartHeaders.Set("Content-Type", contentType)
			attachmentPartHeaders.Set("Content-Transfer-Encoding", "base64")
			attachmentPartHeaders.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))

			attachmentPart, err := multipartWriter.CreatePart(attachmentPartHeaders)
			if err != nil {
				fmt.Println("Error creating attachment part:", err)
				return nil, nil
			}

			encoder := base64.NewEncoder(base64.StdEncoding, attachmentPart)
			encoder.Write(data)
			encoder.Close()
		}

		multipartWriter.Close()

	} else {

		header.Set("Content-Type", "text/plain; charset=utf-8")

		for k, v := range header {
			buffer.WriteString(fmt.Sprintf("%s: %s\n", k, strings.Join(v, ", ")))
		}
		buffer.WriteString("\n")
		buffer.WriteString(m.body)
	}

	return buffer.Bytes(), allRecipients
}
