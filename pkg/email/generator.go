package email

import (
	"bytes"
	"html/template"
)

type emailBodyGenerator struct {
	t *template.Template
}

var verification = template.Must(template.ParseFiles("web/template/verification.html"))

func generateVerificationEmail(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	if err := verification.Execute(buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
