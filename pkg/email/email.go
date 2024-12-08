package email

import (
	"fmt"
	"net/smtp"
)

type EmailSender interface {
	SendVerificationCode(to string, code string) error
	WarningMessageIP(ip string) error
}

type SMTPEmailSender struct {
	SMTPServer string
	Port       string
	Username   string
	Password   string
}

func NewSMTPEmailSender(smtpServer, port, username, password string) *SMTPEmailSender {
	return &SMTPEmailSender{
		SMTPServer: smtpServer,
		Port:       port,
		Username:   username,
		Password:   password,
	}
}

func (e *SMTPEmailSender) SendVerificationCode(to string, code string) error {
	subject := "Верификация вашего аккаунта"
	body := fmt.Sprintf("Ваш код верификации: %s", code)
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPServer)
	addr := fmt.Sprintf("%s:%s", e.SMTPServer, e.Port)

	return smtp.SendMail(addr, auth, e.Username, []string{to}, message)
}

func (e *SMTPEmailSender) WarningMessageIP(ip string) error {
	//TODO Дописать код
	return nil
}
