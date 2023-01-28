package email

import (
	"errors"
	"strings"

	gosendmail "github.com/ondrejsika/gosendmail/lib"
	"github.com/sikalabs/tergum/version"
)

type EmailBackend struct {
	SmtpHost string `yaml:"SmtpHost"`
	SmtpPort string `yaml:"SmtpPort"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	From     string `yaml:"From"`
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
	var finalSubject string
	if strings.HasPrefix(subject, "[") {
		finalSubject = "[tergum]" + subject
	} else {
		finalSubject = "[tergum] " + subject
	}
	finalBody := body + "\n\n--\ntergum, " + version.Version
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
