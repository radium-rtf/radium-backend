package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Course struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Slug string    `json:"slug"`

		ShortDescription string `json:"shortDescription"`
		Description      string `json:"description"`
		Logo             string `json:"logo"`
		Banner           string `json:"banner"`
		Authors          []User `json:"authors"`
		Links            []Link `json:"links"`

		Score    uint      `json:"score"`
		MaxScore uint      `json:"maxScore"`
		Modules  []*Module `json:"modules"` // TODO: скрыть для людей, у которых нет доступа к курсу
	}
)

func NewCourses(courses []*entity.Course) []*Course {
	c := make([]*Course, 0, len(courses))
	for _, course := range courses {
		c = append(c, NewCourse(course, map[uuid.UUID]*entity.Answer{}))
	}
	return c
}

func NewCourse(course *entity.Course, answers map[uuid.UUID]*entity.Answer) *Course {
	links := make([]Link, 0, len(course.Links))
	for _, link := range course.Links {
		links = append(links, Link{Name: link.Name, Link: link.Link})
	}

	authors := make([]User, 0, len(course.Authors))
	for _, author := range course.Authors {
		authors = append(authors,
			User{Id: author.Id, Name: author.Name, Avatar: author.Avatar, Email: author.Email})
	}

	modules, score, maxScore := NewModules(course.Modules, answers)

	return &Course{
		Id:               course.Id,
		Name:             course.Name,
		Slug:             course.Slug,
		ShortDescription: course.ShortDescription,
		Description:      course.Description,
		Logo:             course.Logo,
		Banner:           course.Banner,
		Authors:          authors,
		Links:            links,
		MaxScore:         maxScore,
		Score:            score,
		Modules:          modules,
	}
}
