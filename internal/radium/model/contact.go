package model

import (
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type (
	Contact struct {
		Name string `json:"name" validate:"required,min=1,max=32"`
		Link string `json:"link" validate:"required,url"`
	}
)

func NewContacts(contacts []*entity.Contact) []*Contact {
	dtos := make([]*Contact, 0, len(contacts))
	for _, contact := range contacts {
		dtos = append(dtos, NewContact(contact))
	}
	return dtos
}

func NewContact(contact *entity.Contact) *Contact {
	if contact == nil {
		return nil
	}

	return &Contact{
		Name: contact.Name,
		Link: contact.Link,
	}
}
