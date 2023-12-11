package email

import (
	"bytes"
	"html/template"
	"log"
)

type emailBodyGenerator struct {
	t *template.Template
}

func newEmailBodyGenerator(templatePath string) (emailBodyGenerator, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Failed to read template file %s:%s", templatePath, err.Error())
		return emailBodyGenerator{}, err
	}

	return emailBodyGenerator{t: t}, nil
}

func (g *emailBodyGenerator) generateEmailBodyFromHTML(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	if err := g.t.Execute(buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
