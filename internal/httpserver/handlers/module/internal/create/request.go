package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Request struct {
	CourseId uuid.UUID `json:"courseId"`
	Name     string    `json:"name"`
}

func (m Request) ToModule() *entity.Module {
	return &entity.Module{
		Name:     m.Name,
		CourseId: m.CourseId,
		Slug:     translit.RuEn(m.Name),
	}
}
