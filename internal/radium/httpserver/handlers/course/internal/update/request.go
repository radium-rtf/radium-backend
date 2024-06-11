package update

import (
	"github.com/google/uuid"
	entity "github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type Course struct {
	Name             string `json:"name" validate:"max=64"`
	ShortDescription string `json:"shortDescription" validate:"max=512"`
	Description      string `json:"description" validate:"max=4096"`
	Logo             string `json:"logo" validate:"url"`
	Banner           string `json:"banner" validate:"url"`
	Slug             string `json:"slug,omitempty" validate:"max=64,slug"`
}

func (c Course) toCourse(id uuid.UUID) *entity.Course {
	return &entity.Course{
		DBModel:          entity.DBModel{Id: id},
		Name:             c.Name,
		ShortDescription: c.ShortDescription,
		Description:      c.Description,
		Logo:             c.Logo,
		Banner:           c.Banner,
		Slug:             c.Slug,
	}
}
