package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrPageNotFound = errors.New("страницы курса не найдены")
)

type (
	Page struct {
		DBModel
		Name     string     `json:"name" gorm:"type:string; not null"`
		Slug     string     `json:"slug" gorm:"type:string; not null"`
		ModuleId uuid.UUID  `json:"moduleId" gorm:"type:uuid; not null"`
		Sections []*Section `json:"sections"`
	}

	PageDto struct {
		Id       uuid.UUID     `json:"id"`
		Slug     string        `json:"slug"`
		Name     string        `json:"name"`
		Sections []*SectionDto `json:"sections"`
	}

	PagePost struct {
		ModuleId uuid.UUID `json:"moduleId"`
		Name     string    `json:"name"`
	}
)
