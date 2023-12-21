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

func NewSMTPSender(username, pass, host string, port int,
	verificationCodeLength int, from string) *SMTPSender {
	dialer := gomail.NewDialer(host, port, username, pass)

	return &SMTPSender{
		dialer:                 dialer,
		isAvailable:            true,
		verificationCodeLength: verificationCodeLength,
		otp:                    otp.NewOTPGenerator(),
		from:                   from,
	}
}

type verificationData struct {
	Name, Code string
}

func (s *SMTPSender) SendVerificationEmail(email, name, code string) error {
	if !s.isAvailable {
		return errors.New("отправка сообщений недоступна")
	}

	body, err := generateVerificationEmail(verificationData{name, code})
	if err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Поддтверждение почты радиум")
	msg.SetBody("text/html", body)

	if err := s.dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("error while sending email: %s", err.Error())
	}

	return nil
}
