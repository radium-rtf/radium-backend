package entity

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCourseNotFound = errors.New("курс не найден")
)

type (
	Course struct {
		DBModel
		Name             string   `gorm:"type:string; not null"`
		Slug             string   `gorm:"type:string; not null"`
		ShortDescription string   `gorm:"not null"`
		Description      string   `gorm:"not null"`
		Logo             string   `gorm:"not null"`
		Banner           string   `json:"banner"`
		Authors          []User   `gorm:"many2many:course_authors;"`
		Students         []*User  `gorm:"many2many:course_students;"`
		Groups           []*Group `gorm:"many2many:group_course"`
		Links            []Link
		Modules          []*Module
	}
)

func (c Course) SectionsIds() []uuid.UUID {
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
