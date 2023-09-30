package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Roles struct {
		bun.BaseModel `bun:"table:roles"`
		UserId        uuid.UUID
		IsAuthor      bool
		IsTeacher     bool
	}
)
