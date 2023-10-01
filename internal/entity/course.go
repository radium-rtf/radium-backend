package entity

import (
	"errors"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var (
	ErrCourseNotFound = errors.New("курс не найден")
)

type (
	Course struct {
		bun.BaseModel `bun:"table:courses"`
		DBModel

		Name             string `gorm:"type:string; not null"`
		Slug             string `gorm:"type:string; not null"`
		ShortDescription string `gorm:"not null"`
		Description      string `gorm:"not null"`
		Logo             string `gorm:"not null"`
		Banner           string `json:"banner"`

		Authors  []User `bun:"m2m:course_author,join:Course=User"`
		Students []User `bun:"m2m:course_student,join:Course=User"`
		Links    []Link `bun:"rel:has-many,join:id=course_id"`
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
