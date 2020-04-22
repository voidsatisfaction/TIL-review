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

func (mm *MailManager) Send(from string, to []string, body string) error {
	smtpServerAddress := mm.smtpServerAddress()
	err := smtp.SendMail(smtpServerAddress, mm.auth, from, to, []byte(body))
	if err != nil {
		return err
	}
	return nil
}

func (mm *MailManager) smtpServerAddress() string {
	return fmt.Sprintf("%s:%s", mm.smtpHost, mm.smtpPort)
}
