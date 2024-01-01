package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/str"
)

type Module struct {
	CourseId uuid.UUID `json:"courseId"`
	Name     string    `json:"name" validate:"required,min=1,max=40"`
	Order    float64   `json:"order" validate:"numeric"`
}

func (m Module) toModule() *entity.Module {
	return &entity.Module{
		DBModel:  entity.DBModel{Id: uuid.New()},
		Name:     m.Name,
		CourseId: m.CourseId,
		Order:    m.Order,
		Slug:     str.Random(11),
	}
}
