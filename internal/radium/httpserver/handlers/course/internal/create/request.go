package create

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/str"
)

type Course struct {
	Name             string       `json:"name" validate:"max=64"`
	ShortDescription string       `json:"shortDescription" validate:"max=512"`
	Description      string       `json:"description" validate:"max=4096"`
	Logo             string       `json:"logo" validate:"url"`
	Banner           string       `json:"banner" validate:"url"`
	Links            []model.Link `json:"links" validate:"dive"`
}

func (r Course) toCourse(authorId uuid.UUID) *entity2.Course {
	courseId := uuid.New()
	authors := make([]*entity2.User, 0, 1)
	authors = append(authors, &entity2.User{DBModel: entity2.DBModel{Id: authorId}})

	links := make([]*entity2.Link, 0, len(r.Links))
	for _, v := range r.Links {
		link := &entity2.Link{
			Link:     v.Link,
			Name:     v.Name,
			CourseId: courseId,
			DBModel:  entity2.DBModel{Id: uuid.New()},
		}
		links = append(links, link)
	}

	slug := str.Random(11)

	return &entity2.Course{
		DBModel:          entity2.DBModel{Id: courseId},
		Name:             r.Name,
		Banner:           r.Banner,
		Slug:             slug,
		ShortDescription: r.ShortDescription,
		Description:      r.Description,
		Logo:             r.Logo,
		Authors:          authors,
		Links:            links,
	}
}
