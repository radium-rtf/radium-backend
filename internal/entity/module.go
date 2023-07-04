package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrModulesNotFound = errors.New("модули курса не найдены")
)

type (
	ModulePost struct {
		CourseId uuid.UUID `json:"courseId"`
		Name     string    `json:"name"`
	}

	Module struct {
		DBModel
		Slug     string    `json:"slug" gorm:"type:string; not null"`
		Name     string    `json:"name" gorm:"type:string; not null"`
		CourseId uuid.UUID `json:"courseId" gorm:"type:uuid; not null"`
		Pages    []*Page
	}

	ModuleDto struct {
		Id    uuid.UUID `json:"id"`
		Slug  string    `json:"slug"`
		Name  string    `json:"name"`
		Pages []PageDto `json:"pages"`
	}
)
