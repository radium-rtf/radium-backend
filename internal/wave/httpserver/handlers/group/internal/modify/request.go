package modify

import "github.com/google/uuid"

type GroupMember struct {
	UserId uuid.UUID `json:"userId" validate:"required"`
	ChatId uuid.UUID `json:"chatId" validate:"required"`
}
