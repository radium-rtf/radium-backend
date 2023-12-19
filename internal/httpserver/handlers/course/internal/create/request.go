package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Course struct {
	Name             string       `json:"name" validate:"max=128"`
	ShortDescription string       `json:"shortDescription" validate:"max=400"`
	Description      string       `json:"description" validate:"max=3000"`
	Logo             string       `json:"logo" validate:"url"`
	Banner           string       `json:"banner" validate:"url"`
	Links            []model.Link `json:"links" validate:"dive"`
}

func (r Course) toCourse(authorId uuid.UUID) *entity.Course {
	courseId := uuid.New()
	authors := make([]*entity.User, 0, 1)
	authors = append(authors, &entity.User{DBModel: entity.DBModel{Id: authorId}})

	links := make([]*entity.Link, 0, len(r.Links))
	for _, v := range r.Links {
		link := &entity.Link{
			DBModel:  entity.DBModel{Id: uuid.New()},
			Name:     v.Name,
			Link:     v.Link,
			CourseId: courseId,
		}
		links = append(links, link)
	}

	slug := translit.Make(r.Name)
	if slug == "" {
		slug = courseId.String()
	}

	return &entity.Course{
		DBModel:          entity.DBModel{Id: courseId},
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
