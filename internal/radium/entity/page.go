package entity

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
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

	LastVisitedPage struct {
		bun.BaseModel `bun:"table:last_visited_page"`

		UserId uuid.UUID `bun:",pk"`
		User   *User     `bun:"rel:belongs-to,join:user_id=id"`

		PageId uuid.UUID
		Page   *Page `bun:"rel:belongs-to,join:course_id=id"`

		CourseId uuid.UUID `bun:",pk"`
		Course   *Course   `bun:"rel:belongs-to,join:course_id=id"`

		UpdatedAt time.Time `bun:",nullzero"`
	}
)

func (m *LastVisitedPage) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
