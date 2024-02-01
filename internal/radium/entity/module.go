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
		Course   *Course `bun:"rel:belongs-to,join:course_id=id"`

		Slug  string
		Name  string
		Order float64

		Pages []*Page `bun:"rel:has-many,join:id=module_id"`
	}
)
