package update

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type Module struct {
	Name string `json:"name" validate:"required,min=1,max=48"`
}

func (m Module) toModule(moduleId uuid.UUID) *entity2.Module {
	return &entity2.Module{
		Name:    m.Name,
		DBModel: entity2.DBModel{Id: moduleId},
	}
}
