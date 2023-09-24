package update

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Module struct {
	Name string `json:"name" validate:"required,min=1,max=20"`
}

func (m Module) toModule(moduleId uuid.UUID) *entity.Module {
	return &entity.Module{
		Name:    m.Name,
		Slug:    translit.Make(m.Name),
		DBModel: entity.DBModel{Id: moduleId},
	}
}
