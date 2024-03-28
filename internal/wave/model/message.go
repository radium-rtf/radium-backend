package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Message struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewMessage(message *entity.Message) Message {
	return Message{
		Id: message.Id,
	}
}
