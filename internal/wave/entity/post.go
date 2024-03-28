package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Post struct {
		bun.BaseModel `bun:"table:posts"`
		DBModel

		ChannelId uuid.UUID
		Channel   *Channel `bun:"rel:belongs-to,join:channel_id=id"`
		Content   Content
	}
)
