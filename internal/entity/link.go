package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Link struct {
		bun.BaseModel `bun:"table:links"`
		DBModel
		Name string `validate:"required,min=1,max=15"`
		Link string `validate:"required,url"`

		Course   *Course `bun:"rel:belongs-to,join:course_id=id"`
		CourseId uuid.UUID
	}
)
