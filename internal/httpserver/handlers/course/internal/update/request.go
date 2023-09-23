package update

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Course struct {
	Name             string `json:"name" validate:"required,min=1,max=45"`
	ShortDescription string `json:"shortDescription" validate:"required,min=1,max=200"`
	Description      string `json:"description" validate:"max=3000"`
	Logo             string `json:"logo" validate:"url"`
	Banner           string `json:"banner" validate:"url"`
}

func (c Course) toCourse(id uuid.UUID) *entity.Course {
	return &entity.Course{
		DBModel:          entity.DBModel{Id: id},
		Name:             c.Name,
		Slug:             translit.RuEn(c.Name),
		ShortDescription: c.ShortDescription,
		Description:      c.Description,
		Logo:             c.Logo,
		Banner:           c.Banner,
	}
}