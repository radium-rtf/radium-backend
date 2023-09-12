package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Course struct {
	Name             string       `json:"name" validate:"required,min=1,max=45"`
	ShortDescription string       `json:"shortDescription" validate:"required,min=1,max=200"`
	Description      string       `json:"description" validate:"max=3000"`
	Logo             string       `json:"logo" validate:"url"`
	Banner           string       `json:"banner" validate:"url"`
	Links            []model.Link `json:"links" validate:"dive"`
}

func (r Course) toCourse(authorId uuid.UUID) *entity.Course {
	authors := make([]entity.User, 0, 1)
	authors = append(authors, entity.User{DBModel: entity.DBModel{Id: authorId}})

	links := make([]entity.Link, 0, len(r.Links))
	for _, v := range r.Links {
		links = append(links, entity.Link{Name: v.Name, Link: v.Link})
	}

	return &entity.Course{
		Name:             r.Name,
		Slug:             translit.RuEn(r.Name),
		ShortDescription: r.ShortDescription,
		Description:      r.Description,
		Logo:             r.Logo,
		Authors:          authors,
		Links:            links,
	}
}
