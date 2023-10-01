package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Link struct {
		bun.BaseModel `bun:"table:links"`
		DBModel
		Name     string
		Link     string
		CourseId uuid.UUID
	}
)
