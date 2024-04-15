package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var (
	ReviewNotification NotificationType = "review"
)

type (
	NotificationType string

	Notification struct {
		bun.BaseModel `bun:"table:notifications"`
		DBModel

		UserId uuid.UUID
		User   *User `bun:"rel:belongs-to,join:user_id=id"`

		AnswerId uuid.UUID `bun:",nullzero"`
		Answer   *Answer   `bun:"rel:belongs-to,join:answer_id=id"`

		Type NotificationType
		Read bool
	}
)
