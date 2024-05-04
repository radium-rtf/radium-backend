package update

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type Page struct {
	Name string `json:"name" validate:"required,min=1,max=48"`
}

func (r Page) toPage(pageId uuid.UUID) *entity2.Page {
	return &entity2.Page{
		DBModel: entity2.DBModel{Id: pageId},
		Name:    r.Name,
	}
}
