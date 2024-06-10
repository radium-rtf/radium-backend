package send

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
)

type MessageSend struct {
	ChatId          uuid.UUID     `json:"chatId" validate:"required"`
	Content         model.Content `json:"content" validate:"required"`
	ParentMessageId *uuid.UUID    `json:"parent_message_id"`
}
