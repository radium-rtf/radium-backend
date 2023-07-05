package mapper

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type Module struct {
	p Page
}

func (m Module) PostToModule(module entity.ModulePost) *entity.Module {
	return &entity.Module{
		Name:     module.Name,
		CourseId: module.CourseId,
		Slug:     translit.RuEn(module.Name),
	}
}

func (m Module) ModulesToDto(modules []*entity.Module) []*entity.ModuleDto {
	dto := make([]*entity.ModuleDto, 0, len(modules))
	for _, module := range modules {
		dto = append(dto, m.ToDto(module))
	}
	return dto
}

func (m Module) ToDto(module *entity.Module) *entity.ModuleDto {
	return &entity.ModuleDto{
		Id:    module.Id,
		Name:  module.Name,
		Slug:  translit.RuEn(module.Name),
		Pages: m.p.Pages(module.Pages, map[uuid.UUID]*entity.Answer{}),
	}
}
