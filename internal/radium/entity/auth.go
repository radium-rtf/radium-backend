package entity

import (
	"github.com/uptrace/bun"
	"time"

	"github.com/google/uuid"
)

type (
	Session struct {
		bun.BaseModel `bun:"table:sessions"`
		RefreshToken  uuid.UUID `bun:",pk"`
		ExpiresIn     time.Time
		UserId        uuid.UUID
		User          *User `bun:"rel:belongs-to,join:user_id=id"`
	}
)
