package entity

import "errors"

var (
	ModulesNotFoundErr = errors.New("модули курса не найдены")
)

type (
	ModuleRequest struct {
		Name string `json:"name"`
	}

	Module struct {
		Id       string `json:"id"`
		CourseId int    `json:"course_id"`
		Name     string `json:"name"`
	}

	ModuleDto struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
)
