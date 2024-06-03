package entity

import (
	"github.com/google/uuid"
	radium "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/uptrace/bun"
)

type (
	Group struct {
		bun.BaseModel `bun:"table:wave.groups"`
		DBModel

		Name           string
		AvatarUrl      string
		Members        []*radium.User `bun:"m2m:wave.group_member,join:Group=User"`
		Settings       GroupSettings  `bun:"rel:has-one"`
		Admins         []*radium.User `bun:"m2m:wave.group_admin,join:Group=User"`
		PinnedMessages []*Message     `bun:"m2m:wave.group_pinned,join:Group=Message"`
		Messages       []*Message     `bun:"m2m:wave.group_message,join:Group=Message"`
	}

	GroupMember struct {
		bun.BaseModel `bun:"table:wave.group_member"`

		GroupId uuid.UUID `bun:",pk"`
		Group   *Group    `bun:"rel:belongs-to,join:group_id=id"`

		UserId uuid.UUID    `bun:",pk"`
		User   *radium.User `bun:"rel:belongs-to,join:user_id=id"`
		Role   string
	}

	GroupAdmin struct {
		bun.BaseModel `bun:"table:wave.group_admin"`

		GroupId uuid.UUID `bun:",pk"`
		Group   *Group    `bun:"rel:belongs-to,join:group_id=id"`

		UserId uuid.UUID    `bun:",pk"`
		User   *radium.User `bun:"rel:belongs-to,join:user_id=id"`
	}

	GroupSettings struct {
		bun.BaseModel `bun:"table:wave.group_settings"`
		DBModel
	}
)
