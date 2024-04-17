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
		Subscribers []*radium.User  `bun:"m2m:wave.channel_subscriber,join:Channel=User"`
		OwnerId     uuid.UUID       `bun:",notnull"`
		Owner       *radium.User    `bun:"rel:belongs-to,join:owner_id=id"`
		Admins      []*radium.User  `bun:"m2m:wave.channel_admin,join:Channel=User"`
		Groups      []*Group        `bun:"m2m:wave.channel_group,join:Channel=Group"`
	}

	ChannelSettings struct {
		bun.BaseModel `bun:"table:wave.channel_settings"`
		DBModel

		ChannelId uuid.UUID `bun:",pk"`
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

	ChannelGroup struct {
		bun.BaseModel `bun:"table:wave.channel_group"`

		ChannelId uuid.UUID `bun:",pk"`
		Channel   *Channel  `bun:"rel:belongs-to,join:channel_id=id"`

		GroupId uuid.UUID `bun:",pk"`
		Group   *Group    `bun:"rel:belongs-to,join:group_id=id"`
	}
)
