package entity

import (
	"github.com/google/uuid"
	radium "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/uptrace/bun"
)

type (
	Channel struct {
		bun.BaseModel `bun:"table:channels"`
		DBModel
		Name string

		Settings    ChannelSettings `bun:"rel:has-one"`
		Posts       []*Post         `bun:"rel:has-many"`
		Subscribers []*radium.User  `bun:"m2m:channel_subscribers,join:Channel=User"`
		OwnerId     uuid.UUID
		Admins      []*radium.User `bun:"m2m:channel_admins,join:Channel=User"`
		GroupChatId uuid.UUID
		GroupChat   *GroupChat `bun:"rel:belongs-to,join:group_id=id"`
	}

	ChannelSettings struct {
		bun.BaseModel `bun:"table:channel_settings"`
		DBModel
	}

	ChannelAdmin struct {
		bun.BaseModel `bun:"table:channel_admins"`

		ChannelId uuid.UUID `bun:",pk"`
		Channel   *Channel  `bun:"rel:belongs-to,join:channel_id=id"`

		UserId uuid.UUID    `bun:",pk"`
		User   *radium.User `bun:"rel:belongs-to,join:user_id=id"`
	}

	ChannelSubscriber struct {
		bun.BaseModel `bun:"table:channel_subscribers"`

		ChannelId uuid.UUID `bun:",pk"`
		Channel   *Channel  `bun:"rel:belongs-to,join:channel_id=id"`

		UserId uuid.UUID    `bun:",pk"`
		User   *radium.User `bun:"rel:belongs-to,join:user_id=id"`
	}
)
