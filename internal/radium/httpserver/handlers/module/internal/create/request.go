package create

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/pkg/str"
)

type Module struct {
	CourseId uuid.UUID `json:"courseId"`
	Name     string    `json:"name" validate:"required,min=1,max=48"`
	Order    float64   `json:"order" validate:"numeric"`
}

func (m Module) toModule() *entity2.Module {
	return &entity2.Module{
		DBModel:  entity2.DBModel{Id: uuid.New()},
		Name:     m.Name,
		CourseId: m.CourseId,
		Order:    m.Order,
		Slug:     str.Random(11),
	}
}
