package modify

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
)

type MessageEdit struct {
	MessageId uuid.UUID     `json:"messageId" validate:"required"`
	Content   model.Content `json:"content" validate:"required"`
}

type MessageGeneric struct {
	MessageId uuid.UUID `json:"messageId" validate:"required"`
}
