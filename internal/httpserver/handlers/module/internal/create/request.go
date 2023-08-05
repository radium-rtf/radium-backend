package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Module struct {
	CourseId uuid.UUID `json:"courseId"`
	Name     string    `json:"name"`
	Order    float64   `json:"order"`
}

func (m Module) toModule() *entity.Module {
	return &entity.Module{
		Name:     m.Name,
		CourseId: m.CourseId,
		Order:    m.Order,
		Slug:     translit.RuEn(m.Name),
	}
}
