package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Message struct {
		Id uuid.UUID `json:"id"`

		SenderId        uuid.UUID  `json:"senderId"`
		Chat            *Chat      `json:"chat,omitempty"`
		Content         Content    `json:"content"`
		ParentMessageId *uuid.UUID `json:"parentMessageId"`
		Type            string     `json:"type"`
		Pinned          bool       `json:"pinned"`
	}
)

func NewMessage(message *entity.Message) *Message {
	if message == nil {
		return nil
	}
	content := NewContent(message.Content)

	return &Message{
		Id:              message.Id,
		SenderId:        message.SenderId,
		Content:         content,
		ParentMessageId: message.ParentMessageId,
		Type:            message.Type,
		Pinned:          message.IsPinned,
	}
}

func NewMessages(messages []*entity.Message) []*Message {
	result := make([]*Message, 0, len(messages))
	for _, message := range messages {
		result = append(result, NewMessage(message))
	}
	return result
}
