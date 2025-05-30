package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Message struct {
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	Attachments map[string][]byte
}

func (m *Message) AddTo(recipient string) {
	m.to = append(m.to, recipient)
}

func (m *Message) AddCC(cc_recipient string) {
	m.cc = append(m.cc, cc_recipient)
}

func (m *Message) AddBCC(bcc_recipient string) {
	m.bcc = append(m.bcc, bcc_recipient)
}

func (m *Message) AddSubject(subject string) {
	m.subject = subject
}

func (m *Message) AddBody(body string) {
	m.body = body
}

func (m *Message) AttachFile(path string) error {

	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	_, filename := filepath.Split(path)

	m.Attachments[filename] = file

	return nil
}

func (m *Message) ToBytes() []byte {

	buffer := bytes.NewBuffer(nil)

	buffer.WriteString(fmt.Sprintf("Subject: %s\n", m.subject))
	buffer.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.to, ",")))

	if len(m.cc) > 0 {
		buffer.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.cc, ",")))
	}

	if len(m.bcc) > 0 {
		buffer.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.bcc, ",")))
	}

	buffer.WriteString("MIME-Version: 1.0\n")

	writer := multipart.NewWriter(buffer)
	boundary := writer.Boundary()

	if len(m.Attachments) > 0 {

		buffer.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buffer.WriteString(fmt.Sprintf("--%s\n", boundary))

	} else {
		buffer.WriteString("Content-type: text/plain; charset=utf-8\n")
	}

	buffer.WriteString(m.body)

	if len(m.Attachments) > 0 {

		for filename, data := range m.Attachments {

			buffer.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buffer.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(data)))
			buffer.WriteString("Content-Transfer-Encoding: base64\n")
			buffer.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", filename))

			base_encoding := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
			base64.StdEncoding.Encode(base_encoding, data)
			buffer.Write(base_encoding)
			buffer.WriteString(fmt.Sprintf("\n--%s", boundary))

		}

		buffer.WriteString("--")

	}

	return buffer.Bytes()

}
