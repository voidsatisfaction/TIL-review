package mailManager

import (
	"fmt"
	"net/smtp"
)

type MailManager struct {
	username string
	password string
	smtpHost string
	smtpPort string
	auth     smtp.Auth
}

func New(username, password, smtpHost, smtpPort string) *MailManager {
	plainAuth := smtp.PlainAuth("", username, password, smtpHost)

	return &MailManager{
		username,
		password,
		smtpHost,
		smtpPort,
		plainAuth,
	}
}

func (mm *MailManager) SendHTML(from string, to []string, subject, htmlString string) error {
	const MIME string = "MIME-version: 1.0;"
	const CONTENTTYPE string = "Content-Type: text/html; "
	const CHARSET string = "charset=\"UTF-8\";"
	subject = fmt.Sprintf("Subject: %s", subject)
	header := fmt.Sprintf("%s\n%s%s", MIME, CONTENTTYPE, CHARSET)

	return mm.send(from, to, subject, header, htmlString)
}

func (mm *MailManager) send(from string, to []string, subject, header, body string) error {
	smtpServerAddress := mm.smtpServerAddress()
	smtpBody := fmt.Sprintf("%s\n%s\n\n%s", subject, header, body)

	err := smtp.SendMail(smtpServerAddress, mm.auth, from, to, []byte(smtpBody))
	if err != nil {
		return err
	}
	return nil
}

func (mm *MailManager) smtpServerAddress() string {
	return fmt.Sprintf("%s:%s", mm.smtpHost, mm.smtpPort)
}
