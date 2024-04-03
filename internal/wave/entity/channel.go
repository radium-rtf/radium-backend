package entity

import (
	"github.com/google/uuid"
	radium "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/uptrace/bun"
)

type (
	Channel struct {
		bun.BaseModel `bun:"table:wave.channels"`
		DBModel
		Name string

		Settings    ChannelSettings `bun:"rel:has-one"`
		Posts       []*Post         `bun:"rel:has-many"`
		Subscribers []*radium.User  `bun:"m2m:wave.channel_subscriber,join:Channel=User"`
		OwnerId     uuid.UUID
		Admins      []*radium.User `bun:"m2m:wave.channel_admin,join:Channel=User"`
		GroupId     uuid.UUID
		Group       *Group `bun:"rel:belongs-to,join:group_id=id"`
	}

	ChannelSettings struct {
		bun.BaseModel `bun:"table:wave.channel_settings"`
		DBModel
	}

	ChannelAdmin struct {
		bun.BaseModel `bun:"table:wave.channel_admin"`

		ChannelId uuid.UUID `bun:",pk"`
		Channel   *Channel  `bun:"rel:belongs-to,join:channel_id=id"`

		UserId uuid.UUID    `bun:",pk"`
		User   *radium.User `bun:"rel:belongs-to,join:user_id=id"`
	}

	ChannelSubscriber struct {
		bun.BaseModel `bun:"table:wave.channel_subscriber"`

		ChannelId uuid.UUID `bun:",pk"`
		Channel   *Channel  `bun:"rel:belongs-to,join:channel_id=id"`

		UserId uuid.UUID    `bun:",pk"`
		User   *radium.User `bun:"rel:belongs-to,join:user_id=id"`
	}
)
