package email

import (
	"errors"
	"fmt"
	"github.com/radium-rtf/radium-backend/pkg/otp"

	"github.com/go-gomail/gomail"
)

type SMTPSender struct {
	dialer                 *gomail.Dialer
	bodyGenerator          emailBodyGenerator
	verificationCodeLength int
	isAvailable            bool
	otp                    *otp.OTPGenerator
	from                   string
}

func NewSMTPSender(username, pass, host string, port int, templatePath string,
	verificationCodeLength int, from string) *SMTPSender {
	dialer := gomail.NewDialer(host, port, username, pass)

	bodyGenerator, err := newEmailBodyGenerator(templatePath)
	return &SMTPSender{
		dialer:                 dialer,
		bodyGenerator:          bodyGenerator,
		isAvailable:            err == nil,
		verificationCodeLength: verificationCodeLength,
		otp:                    otp.NewOTPGenerator(),
		from:                   from,
	}
}

func (s *SMTPSender) SendVerificationEmail(email, code string) error {
	if !s.isAvailable {
		return errors.New("отправка сообщений недоступна")
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Поддтверждение почты радиум")
	msg.SetBody("text/html", code)

	if err := s.dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("error while sending email: %s", err.Error())
	}

	return nil
}
