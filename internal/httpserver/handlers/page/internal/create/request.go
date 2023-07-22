package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Request struct {
	ModuleId uuid.UUID `json:"moduleId"`
	Name     string    `json:"name"`
}

func (r Request) NewPostToPage() *entity.Page {
	return &entity.Page{
		DBModel:  entity.DBModel{Id: uuid.New()},
		Name:     r.Name,
		ModuleId: r.ModuleId,
		Slug:     translit.RuEn(r.Name),
	}
}
