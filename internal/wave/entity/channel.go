package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Channel struct {
		bun.BaseModel `bun:"table:channels"`
		DBModel
		Name string

		Settings    ChannelSettings      `bun:"rel:has-one"`
		Posts       []*Post              `bun:"rel:has-many"`
		Subscribers []*ChannelSubscriber `bun:"rel:has-many"`
		OwnerId     uuid.UUID
		Admins      []*ChannelAdmin `bun:"rel:has-many"`
		GroupId     uuid.UUID
	}

	ChannelSettings struct {
		bun.BaseModel `bun:"table:channel_settings"`
		DBModel
	}

	ChannelAdmin struct {
		bun.BaseModel `bun:"table:channel_admins"`

		ChannelId uuid.UUID `bun:",pk"`
		Channel   *Channel  `bun:"rel:belongs-to,join:channel_id=id"`

		UserId uuid.UUID `bun:",pk"`
	}

	ChannelSubscriber struct {
		bun.BaseModel `bun:"table:channel_subscribers"`

		ChannelId uuid.UUID `bun:",pk"`
		Channel   *Channel  `bun:"rel:belongs-to,join:channel_id=id"`

		UserId uuid.UUID `bun:",pk"`
	}
)
