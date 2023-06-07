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
		Id               uuid.UUID `gorm:"default:gen_random_uuid()"`
		Name             string    `gorm:"type:string"`
		Slug             string    `gorm:"type:string"`
		ShortDescription string
		Description      string
		Logo             string
		Banner           string  `json:"banner"`
		Authors          []User  `gorm:"many2many:course_authors"`
		Students         []*User `gorm:"many2many:course_students"`
		Links            []Link
		Modules          []Module
	}

	CourseDto struct {
		Id   uuid.UUID `json:"id" gorm:"default:gen_random_uuid()"`
		Name string    `json:"name" gorm:"type:string"`
		Slug string    `json:"slug"`

		ShortDescription string      `json:"shortDescription"`
		Description      string      `json:"description"`
		Logo             string      `json:"logo"`
		Banner           string      `json:"banner"`
		Authors          []UserDto   `json:"authors"`
		Links            []Link      `json:"links"`
		Modules          []ModuleDto `json:"modules"` // TODO: скрыть для людей, у которых нет доступа к курсу
	}

	// CourseTitle struct {
	// 	Id            uint      `json:"id"`
	// 	Name          string    `json:"name"`
	// 	Description   string    `json:"description"`
	// 	Author        UserDto   `json:"author"`
	// 	Links         []LinkDto `json:"links"`
	// 	Collaborators []UserDto `json:"collaborators"`
	// 	Logo          string    `json:"logo"`
	// 	Type          string    `json:"type"`
	// }

	// CourseModules struct {
	// 	Id      uint        `json:"id"`
	// 	Name    string      `json:"name"`
	// 	Modules []ModuleDto `json:"modules"`
	// 	Logo    string      `json:"logo"`
	// 	Type    string      `json:"type"`
	// }
)

func NewCourse(c CourseRequest) *Course {
	authorsRes := make([]User, 0)
	for _, v := range c.Authors {
		authorsRes = append(authorsRes, User{Id: v})
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
