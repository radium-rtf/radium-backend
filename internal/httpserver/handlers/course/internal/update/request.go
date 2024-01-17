package update

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Course struct {
	Name             string `json:"name" validate:"max=128"`
	ShortDescription string `json:"shortDescription" validate:"max=400"`
	Description      string `json:"description" validate:"max=3000"`
	Logo             string `json:"logo" validate:"url"`
	Banner           string `json:"banner" validate:"url"`
	Slug             string `json:"slug,omitempty" validate:"max=11"`
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
