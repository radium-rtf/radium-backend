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
	Authors          []uuid.UUID  `json:"authors"`
	Links            []model.Link `json:"links"`
}

func (r Request) ToCourse() *entity.Course {
	authorsRes := make([]entity.User, 0)
	for _, v := range r.Authors {
		authorsRes = append(authorsRes, entity.User{DBModel: entity.DBModel{Id: v}})
	}
	linksRes := make([]entity.Link, 0)
	for _, v := range r.Links {
		linksRes = append(linksRes, entity.Link{Name: v.Link, Link: v.Link})
	}
	return &entity.Course{
		Name:             r.Name,
		Slug:             translit.RuEn(r.Name),
		ShortDescription: r.ShortDescription,
		Description:      r.Description,
		Logo:             r.Logo,
		Authors:          authorsRes,
		Links:            linksRes,
	}
}
