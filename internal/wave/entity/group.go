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

		Name      string
		AvatarUrl string
		OwnerId   uuid.UUID
		Owner     *radium.User   `bun:"rel:belongs-to,join:owner_id=id"`
		Members   []*radium.User `bun:"m2m:wave.group_member,join:Group=User"`
		// Settings  GroupSettings  `bun:"rel:has-one"`
	}

	GroupMember struct {
		bun.BaseModel `bun:"table:wave.group_member"`

		GroupId uuid.UUID `bun:",pk"`
		Group   *Group    `bun:"rel:belongs-to,join:group_id=id"`

		UserId uuid.UUID    `bun:",pk"`
		User   *radium.User `bun:"rel:belongs-to,join:user_id=id"`
		Admin  bool
		Role   string
	}

	GroupSettings struct {
		bun.BaseModel `bun:"table:wave.group_settings"`
		DBModel
	}
)
