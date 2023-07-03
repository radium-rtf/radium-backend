package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCourseNotFound = errors.New("курс не найден")
)

type (
	CourseRequest struct {
		Name             string      `json:"name"`
		ShortDescription string      `json:"shortDescription"`
		Description      string      `json:"description"`
		Logo             string      `json:"logo"`
		Banner           string      `json:"banner"`
		Authors          []uuid.UUID `json:"authors"`
		Links            []Link      `json:"links"`
	}

	Course struct {
		DBModel
		Name             string  `gorm:"type:string; not null"`
		Slug             string  `gorm:"type:string; not null"`
		ShortDescription string  `gorm:"not null"`
		Description      string  `gorm:"not null"`
		Logo             string  `gorm:"not null"`
		Banner           string  `json:"banner"`
		Authors          []User  `gorm:"many2many:course_authors"`
		Students         []*User `gorm:"many2many:course_students"`
		Links            []Link
		Modules          []*Module
	}

	CourseDto struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name" gorm:"type:string"`
		Slug string    `json:"slug"`

		ShortDescription string      `json:"shortDescription"`
		Description      string      `json:"description"`
		Logo             string      `json:"logo"`
		Banner           string      `json:"banner"`
		Authors          []UserDto   `json:"authors"`
		Links            []LinkDto   `json:"links"`
		Modules          []ModuleDto `json:"modules"` // TODO: скрыть для людей, у которых нет доступа к курсу
	}
)

func NewCourse(c CourseRequest) *Course {
	authorsRes := make([]User, 0)
	for _, v := range c.Authors {
		authorsRes = append(authorsRes, User{DBModel: DBModel{Id: v}})
	}
	return &Course{
		Name:             c.Name,
		ShortDescription: c.ShortDescription,
		Description:      c.Description,
		Logo:             c.Logo,
		Authors:          authorsRes,
		Links:            c.Links,
	}
}
