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

func getLastMessage(messages []*entity.Message) *entity.Message {
	if len(messages) > 0 {
		return messages[0]
	}
	return nil
}

func NewChat(chat *entity.Chat) *Chat {
	if chat == nil {
		return nil
	}
	chatModel := Chat{
		Id:          chat.Id,
		Name:        chat.Name,
		Type:        chat.Type,
		LastMessage: NewMessage(getLastMessage(chat.Messages)),
	}
	return &chatModel
}
