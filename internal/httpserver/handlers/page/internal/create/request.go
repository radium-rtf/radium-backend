package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Request struct {
	ModuleId uuid.UUID `json:"moduleId"`
	Name     string    `json:"name"`
	Order    float64   `json:"order"`
}

func (r Request) ToPage() *entity.Page {
	return &entity.Page{
		DBModel:  entity.DBModel{Id: uuid.New()},
		Name:     r.Name,
		ModuleId: r.ModuleId,
		Order:    r.Order,
		Slug:     translit.RuEn(r.Name),
	}
}
