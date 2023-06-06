package entity

import (
	"errors"

	"github.com/google/uuid"
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
		Slug  string `json:"slug"`
		Name  string `json:"name"`
		Pages []Page `json:"pages"`
	}
)
