package create

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/pkg/str"
)

type Page struct {
	ModuleId uuid.UUID `json:"moduleId"`
	Name     string    `json:"name" validate:"required,min=1,max=40"`
	Order    float64   `json:"order" validate:"numeric"`
}

func (r Page) toPage() *entity2.Page {
	return &entity2.Page{
		DBModel:  entity2.DBModel{Id: uuid.New()},
		Name:     r.Name,
		ModuleId: r.ModuleId,
		Order:    r.Order,
		Slug:     str.Random(11),
	}
}
