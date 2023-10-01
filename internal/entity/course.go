package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Course struct {
		bun.BaseModel `bun:"table:courses"`
		DBModel

		Name             string
		Slug             string
		ShortDescription string
		Description      string
		Logo             string
		Banner           string

		Students []User `bun:"m2m:course_student,join:Course=User"`

		Authors []User `bun:"m2m:course_author,join:Course=User"`
		Links   []Link `bun:"rel:has-many,join:id=course_id"`

		Modules []*Module `bun:"rel:has-many,join:id=course_id"`
	}

	CourseAuthor struct {
		bun.BaseModel `bun:"table:course_author"`

		CourseId uuid.UUID `bun:",pk"`
		Course   *Course   `bun:"rel:belongs-to,join:course_id=id"`

		UserId uuid.UUID `bun:",pk"`
		User   *User     `bun:"rel:belongs-to,join:user_id=id"`
	}

	CourseStudent struct {
		bun.BaseModel `bun:"table:course_student"`

		CourseId uuid.UUID `bun:",pk"`
		Course   *Course   `bun:"rel:belongs-to,join:course_id=id"`

		UserId uuid.UUID `bun:",pk"`
		User   *User     `bun:"rel:belongs-to,join:user_id=id"`
	}
)

func (c Course) SectionsIds() []uuid.UUID {
	panic("not implemented")
}
