package email

import (
	"errors"
	"fmt"

	"github.com/go-gomail/gomail"
)

type SMTPSender struct {
	from string
	pass string
	host string
	port int
}

func NewSMTPSender(from, pass, host string, port int) (*SMTPSender, error) {
	if !ValidateEmail(from) {
		return nil, errors.New("invalid SMTP email")
	}

	return &SMTPSender{
		from: from,
		pass: pass,
		host: host,
		port: port,
	}, nil
}

func (s *SMTPSender) Send(input SendEmailInput) error {
	if err := input.ValidateEmailDataBeforeSend(); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", input.To)
	msg.SetHeader("Subject", input.Subject)
	msg.SetBody("text/html", input.Body)

	dialer := gomail.NewDialer(s.host, s.port, s.from, s.pass)
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("error while sending email: %s", err.Error())
	}

	return nil
}
