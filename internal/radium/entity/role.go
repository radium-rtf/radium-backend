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
		IsCoauthor    bool
		IsAdmin       bool
	}
)

func GetAllRoles(userId uuid.UUID) *Roles {
	return &Roles{
		UserId:     userId,
		IsAuthor:   true,
		IsCoauthor: true,
		IsTeacher:  true,
		IsAdmin:    true,
	}
}
