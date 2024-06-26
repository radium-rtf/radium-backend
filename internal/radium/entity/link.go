package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Link struct {
		bun.BaseModel `bun:"table:links"`
		DBModel
		Name string
		Link string

		Course   *Course `bun:"rel:belongs-to,join:course_id=id"`
		CourseId uuid.UUID
	}
)
