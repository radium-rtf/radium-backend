package update

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Module struct {
	Name string `json:"name" validate:"required,min=1,max=40"`
}

func (m Module) toModule(moduleId uuid.UUID) *entity.Module {
	return &entity.Module{
		Name:    m.Name,
		DBModel: entity.DBModel{Id: moduleId},
	}
}
