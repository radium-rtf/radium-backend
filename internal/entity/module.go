package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

var (
	ErrModulesNotFound = errors.New("модули курса не найдены")
)

type (
	ModuleRequest struct {
		CourseId uuid.UUID `json:"courseId"`
		Name     string    `json:"name"`
	}

	Module struct {
		Id       uuid.UUID `json:"id" gorm:"default:gen_random_uuid()"`
		Slug     string    `json:"slug" gorm:"type:string"`
		Name     string    `json:"name" gorm:"type:string"`
		CourseId uuid.UUID `json:"courseId"`
		Pages    []Page
	}

	ModuleDto struct {
		Id    uuid.UUID `json:"id"`
		Slug  string    `json:"slug"`
		Name  string    `json:"name"`
		Pages []PageDto `json:"pages"`
	}
)

func NewModuleRequestToModule(module ModuleRequest) Module {
	return Module{
		Name:     module.Name,
		CourseId: module.CourseId,
		Slug:     translit.RuEn(module.Name),
	}
}
