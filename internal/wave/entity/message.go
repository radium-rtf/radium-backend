package entity

import (
	"github.com/google/uuid"
	radium "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/uptrace/bun"
)

type (
	Message struct {
		bun.BaseModel `bun:"table:wave.messages"`
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
		bun.BaseModel `bun:"table:wave.read_by"`

		MessageId uuid.UUID    `bun:",pk"`
		Message   *Message     `bun:"rel:belongs-to,join:message_id=id"`
		UserId    uuid.UUID    `bun:",pk"`
		User      *radium.User `bun:"rel:belongs-to,join:user_id=id"`
	}

	Reaction struct {
		bun.BaseModel `bun:"table:wave.reactions"`
		DBModel

		MessageId uuid.UUID
		Message   *Message `bun:"rel:belongs-to,join:message_id=id"`
		UserId    uuid.UUID
		User      *radium.User `bun:"rel:belongs-to,join:user_id=id"`
		Reaction  string
	}
)
