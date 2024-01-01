package update

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Page struct {
	Name string `json:"name" validate:"required,min=1,max=40"`
}

func (r Page) toPage(pageId uuid.UUID) *entity.Page {
	return &entity.Page{
		DBModel: entity.DBModel{Id: pageId},
		Name:    r.Name,
	}
}
