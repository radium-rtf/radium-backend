package update

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Page struct {
	Name string `json:"name" validate:"required,min=1,max=20"`
}

func (r Page) toPage(pageId uuid.UUID) *entity.Page {
	return &entity.Page{
		DBModel: entity.DBModel{Id: pageId},
		Name:    r.Name,
		Slug:    translit.Make(r.Name),
	}
}
