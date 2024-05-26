package model

import (
	"github.com/google/uuid"
)

type (
	Chat struct {
		Id          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Type        string    `json:"type"`
		LastMessage *Message  `json:"lastMessage,omitempty"`
	}
)

func NewChat(Id uuid.UUID, Name, Type string, LastMessage *Message) Chat {
	chat := Chat{
		Id:   Id,
		Name: Name,
		Type: Type,
	}
	if LastMessage != nil {
		chat.LastMessage = LastMessage
	}
	return chat
}
