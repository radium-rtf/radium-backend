package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"slices"
)

type (
	Course struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Slug string    `json:"slug"`

		ShortDescription string  `json:"shortDescription"`
		Description      string  `json:"description"`
		Logo             string  `json:"logo"`
		Banner           string  `json:"banner"`
		Authors          []*User `json:"authors"`
		Coauthors        []*User `json:"coauthors"`
		Links            []Link  `json:"links"`

		IsPublished bool `json:"isPublished"`
		IsStudent   bool `json:"isStudent"`

		Score    uint      `json:"score"`
		MaxScore uint      `json:"maxScore"`
		Modules  []*Module `json:"modules"` // TODO: скрыть для людей, у которых нет доступа к курсу
	}
)

func NewCourses(courses []*entity.Course, userId uuid.UUID) []*Course {
	c := make([]*Course, 0, len(courses))
	for _, course := range courses {
		c = append(c, NewCourse(course, map[uuid.UUID][]*entity.Answer{}, userId))
	}
	return c
}

func NewCourse(course *entity.Course, answers map[uuid.UUID][]*entity.Answer, userId uuid.UUID) *Course {
	modules, score, maxScore := NewModules(course.Modules, answers)
	isStudent := slices.ContainsFunc(course.Students, func(user entity.User) bool {
		return user.Id == userId
	})
	return &Course{
		Id:               course.Id,
		Name:             course.Name,
		Slug:             course.Slug,
		ShortDescription: course.ShortDescription,
		Description:      course.Description,
		Logo:             course.Logo,
		Banner:           course.Banner,
		Authors:          NewUsers(course.Authors),
		Coauthors:        NewUsers(course.Coauthors),
		Links:            NewLinks(course.Links),
		IsPublished:      course.IsPublished,
		IsStudent:        isStudent,
		MaxScore:         maxScore,
		Score:            score,
		Modules:          modules,
	}
}
