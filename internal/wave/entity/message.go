package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Message struct {
		bun.BaseModel `bun:"table:messages"`
		DBModel

		SenderId        uuid.UUID
		ContentId       uuid.UUID
		Content         *Content `bun:"rel:belongs-to,join:content_id=id"`
		ParentMessageId uuid.UUID
		ParentMessage   *Message    `bun:"rel:belongs-to,join:parent_message_id=id"`
		ReadBy          []*ReadBy   `bun:"rel:has-many"`
		Reactions       []*Reaction `bun:"rel:has-many"`
		Type            string
	}

	ReadBy struct {
		bun.BaseModel `bun:"table:read_by"`

		MessageId uuid.UUID `bun:",pk"`
		Message   *Message  `bun:"rel:belongs-to,join:message_id=id"`
		UserId    uuid.UUID `bun:",pk"`
	}

	Reaction struct {
		bun.BaseModel `bun:"table:reactions"`
		DBModel

		MessageId uuid.UUID
		Message   *Message `bun:"rel:belongs-to,join:message_id=id"`
		UserId    uuid.UUID
		Reaction  string
	}
)
