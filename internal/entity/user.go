package entity

import (
	"database/sql"
	"github.com/uptrace/bun"
)

type (
	User struct {
		bun.BaseModel `bun:"table:users"`
		DBModel

		Avatar   sql.NullString
		Email    string
		Name     string
		Password string

		Roles *Roles `bun:"rel:has-one,join:id=user_id"`
	}
)
