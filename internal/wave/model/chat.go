package model

import (
	"github.com/google/uuid"
)

type (
	Chat struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Type string    `json:"type"`
	}
)

func NewChat(Id uuid.UUID, Name, Type string) Chat {
	return Chat{
		Id:   Id,
		Name: Name,
		Type: Type,
	}
}
