package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Chat struct {
		Id          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Type        string    `json:"type"`
		LastMessage *Message  `json:"lastMessage,omitempty"`
	}
)

func NewChat(chat *entity.Chat, LastMessage *Message) *Chat {
	chatModel := Chat{
		Id:   chat.Id,
		Name: chat.Name,
		Type: chat.Type,
	}
	if LastMessage != nil {
		chatModel.LastMessage = LastMessage
	}
	return &chatModel
}
