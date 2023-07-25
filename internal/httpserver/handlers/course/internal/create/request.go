package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Request struct {
	Name             string       `json:"name"`
	ShortDescription string       `json:"shortDescription"`
	Description      string       `json:"description"`
	Logo             string       `json:"logo"`
	Banner           string       `json:"banner"`
	Links            []model.Link `json:"links"`
}

func (r Request) ToCourse(authorId uuid.UUID) *entity.Course {
	authors := make([]entity.User, 0, 1)
	authors = append(authors, entity.User{DBModel: entity.DBModel{Id: authorId}})

	links := make([]entity.Link, 0, len(r.Links))
	for _, v := range r.Links {
		links = append(links, entity.Link{Name: v.Link, Link: v.Link})
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
