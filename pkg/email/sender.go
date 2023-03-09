package email

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"regexp"
)

type SendEmailInput struct {
	To      string
	Subject string
	Body    string
}

// type Sender interface {
// 	Send(input SendEmailInput) error
// }

func (e *SendEmailInput) GenerateEmailBodyFromHTML(templatePath string, data interface{}) error {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		// ЭЭЭЭЭ ну хз мб стоит сделать кастомный логгер еще хзхзхз
		log.Printf("Failed to read template file %s:%s", templatePath, err.Error())

		return err
	}

	buffer := new(bytes.Buffer)
	if err := t.Execute(buffer, data); err != nil {
		return err
	}

	e.Body = buffer.String()

	return nil
}

func (e *SendEmailInput) ValidateEmail() error {
	var emailRegex = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

	if !emailRegex.MatchString(e.To) {
		return errors.New("invalid email")
	}

	return nil
}

func (e *SendEmailInput) ValidateEmailDataBeforeSend() error {
	if e.To == "" {
		return errors.New("email has no recipients")
	}

	if e.Subject == "" || e.Body == "" {
		return errors.New("email has no subject/body")
	}

	if !ValidateEmail(e.To) {
		return errors.New("email is invalid")
	}

	return nil
}
