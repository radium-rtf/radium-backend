package entity

import "errors"

var (
	ModulesNotFoundErr = errors.New("модули курса не найдены")
)

type (
	ModuleRequest struct {
		CourseId uint   `json:"course_id"`
		Name     string `json:"name"`
	}

	Module struct {
		NameEng  string `json:"name_eng"`
		CourseId uint   `json:"course_id"`
		Name     string `json:"name"`
	}

	ModuleSlides struct {
		Id      uint       `json:"id"`
		Name    string     `json:"name"`
		NameEng string     `json:"name_eng"`
		Slides  []SlideDto `json:"slides"`
	}

	ModuleDto struct {
		NameEng string `json:"name_eng"`
		Name    string `json:"name"`
	}
)
