package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrPageNotFound = errors.New("страницы курса не найдены")
)

type (
	// TODO: добавить секции
	Page struct {
		Id       uuid.UUID `json:"id" gorm:"default:gen_random_uuid()"`
		Name     string    `json:"name" gorm:"type:string"`
		Slug     string    `json:"slug" gorm:"type:string"`
		ModuleId uuid.UUID `json:"moduleId"`
	}

	PageDto struct {
		Id   uuid.UUID `json:"id"`
		Slug string    `json:"slug"`
		Name string    `json:"name"`
	}

	PageRequest struct {
		CourseId uuid.UUID `json:"courseId"`
		Slug     string    `json:"slug"`
		ModuleId uuid.UUID `json:"moduleId"`
		Name     string    `json:"name"`
	}

	// SlidesRequest struct {
	// CourseId      uint   `json:"course_id"`
	// ModuleNameEng string `json:"module_name_eng"`
	// }

	// SlideSectionsRequest struct {
	// CourseId      uint   `json:"course_id"`
	// ModuleNameEng string `json:"module_name_eng"`
	// SlideId uint `json:"slide_id"`
	// }

	// SlideSections struct {
	// 	Id       uint          `json:"id"`
	// 	Name     string        `json:"name"`
	// 	NameEng  string        `json:"name_eng"`
	// 	Sections []interface{} `json:"sections"`
	// }
)
