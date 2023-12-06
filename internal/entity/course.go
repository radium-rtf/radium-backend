package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"slices"
)

type (
	Course struct {
		bun.BaseModel `bun:"table:courses"`
		DBModel

		Name             string `validate:"required,min=1,max=45"`
		Slug             string `validate:"required,min=1,max=45"`
		ShortDescription string `validate:"required,min=1,max=400"`
		Description      string `validate:"required,max=3000"`
		Logo             string `validate:"required,url"`
		Banner           string `validate:"required,url"`
		IsPublished      bool

		Students []User `bun:"m2m:course_student,join:Course=User"`

		Authors   []*User `bun:"m2m:course_author,join:Course=User"`
		Coauthors []*User `bun:"m2m:course_coauthor,join:Course=User"`

		Links []*Link `bun:"rel:has-many,join:id=course_id" validate:"dive"`

		Modules []*Module `bun:"rel:has-many,join:id=course_id"`

		Groups []*Group `bun:"m2m:group_course,join:Course=Group"`
	}

	CourseAuthor struct {
		bun.BaseModel `bun:"table:course_author"`

		CourseId uuid.UUID `bun:",pk"`
		Course   *Course   `bun:"rel:belongs-to,join:course_id=id"`

		UserId uuid.UUID `bun:",pk"`
		User   *User     `bun:"rel:belongs-to,join:user_id=id"`
	}

	CourseCoauthor struct {
		bun.BaseModel `bun:"table:course_coauthor"`

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
	// todo: должно быть подзапросом
	sectionsIds := make([]uuid.UUID, 0, 10)
	for _, module := range c.Modules {
		for _, page := range module.Pages {
			for _, section := range page.Sections {
				sectionsIds = append(sectionsIds, section.Id)
			}
		}
	}
	return sectionsIds
}

func (c Course) CanEdit(editorId uuid.UUID) bool {
	canEdit := slices.ContainsFunc(c.Authors, func(user *User) bool {
		return editorId == user.Id
	})
	canEdit = canEdit || slices.ContainsFunc(c.Coauthors, func(user *User) bool {
		return editorId == user.Id
	})
	return canEdit
}
