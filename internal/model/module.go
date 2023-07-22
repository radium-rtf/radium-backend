package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type (
	Module struct {
		Id    uuid.UUID `json:"id"`
		Slug  string    `json:"slug"`
		Name  string    `json:"name"`
		Pages []*Page   `json:"pages"`
	}
)

func NewModule(module *entity.Module) *Module {
	return &Module{
		Id:    module.Id,
		Name:  module.Name,
		Slug:  translit.RuEn(module.Name),
		Pages: NewPages(module.Pages, map[uuid.UUID]*entity.Answer{}),
	}
}

func NewModules(modules []*entity.Module) []*Module {
	dto := make([]*Module, 0, len(modules))
	for _, module := range modules {
		dto = append(dto, NewModule(module))
	}
	return dto
}
