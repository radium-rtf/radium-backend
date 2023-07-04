package mapper

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Module struct {
}

func (m Module) PostToModule(module entity.ModulePost) *entity.Module {
	return &entity.Module{
		Name:     module.Name,
		CourseId: module.CourseId,
		Slug:     translit.RuEn(module.Name),
	}
}
