package email

import (
	"errors"
	"strings"

	gosendmail "github.com/ondrejsika/gosendmail/lib"
	"github.com/sikalabs/tergum/version"
)

type EmailBackend struct {
	SmtpHost string `yaml:"SmtpHost" json:"SmtpHost,omitempty"`
	SmtpPort string `yaml:"SmtpPort" json:"SmtpPort,omitempty"`
	Username string `yaml:"Username" json:"Username,omitempty"`
	Password string `yaml:"Password" json:"Password,omitempty"`
	From     string `yaml:"From" json:"From,omitempty"`
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
	if b.Username == "" {
		return errors.New("EmailBackend backend requires Username")
	}
	if b.Password == "" {
		return errors.New("EmailBackend backend requires Password")
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
	finalBody := body + "\n\n--\n<br/>tergum, " + version.Version + "\n<br/>https://github.com/sikalabs/tergum"
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
		b.Username,
		b.Password,
		b.From,
		to,
		rawMessage,
	)
	return err
}
