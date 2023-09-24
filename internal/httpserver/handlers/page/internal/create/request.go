package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Page struct {
	ModuleId uuid.UUID `json:"moduleId"`
	Name     string    `json:"name" validate:"required,min=1,max=20"`
	Order    float64   `json:"order" validate:"numeric"`
}

func (r Page) toPage() *entity.Page {
	return &entity.Page{
		DBModel:  entity.DBModel{Id: uuid.New()},
		Name:     r.Name,
		ModuleId: r.ModuleId,
		Order:    r.Order,
		Slug:     translit.Make(r.Name),
	}
}
