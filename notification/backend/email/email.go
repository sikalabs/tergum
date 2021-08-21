package email

import (
	"errors"

	gosendmail "github.com/ondrejsika/gosendmail/lib"
)

type EmailBackend struct {
	SmtpHost string
	SmtpPort string
	Email    string
	Password string
	From     string
}

func (b EmailBackend) Validate() error {
	if b.SmtpHost == "" {
		return errors.New("EmailBackend backend requires SmtpHost")
	}
	if b.SmtpPort == "" {
		return errors.New("EmailBackend backend requires SmtpPort")
	}
	if b.From == "" {
		return errors.New("EmailBackend backend requires From")
	}
	return nil
}

func (b EmailBackend) SendMail(
	to string,
	subject string,
	body string,
) error {
	finalSubject := "[tergum] " + subject
	finalBody := body + "\n\n--\ntergum"
	rawMessage := []byte("To: " + to + "\r\n" +
		"From: " + b.From + "\r\n" +
		"Subject: " + finalSubject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Disposition: inline\r\n" +
		"\r\n" +
		finalBody + "\r\n")
	err := gosendmail.GoRawSendMail(
		b.SmtpHost,
		b.SmtpPort,
		b.From,
		b.Password,
		to,
		rawMessage,
	)
	return err
}
