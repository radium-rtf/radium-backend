package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Module struct {
		bun.BaseModel `bun:"table:modules"`
		DBModel

		CourseId uuid.UUID

		Slug  string
		Name  string
		Order float64
	}
)
