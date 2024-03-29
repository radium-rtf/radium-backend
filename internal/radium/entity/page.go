package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Page struct {
		bun.BaseModel `bun:"table:pages"`
		DBModel

		ModuleId uuid.UUID
		Module   *Module `bun:"rel:belongs-to,join:module_id=id"`

		Name  string
		Slug  string
		Order float64

		Sections []*Section `bun:"rel:has-many,join:id=page_id"`
	}
)
