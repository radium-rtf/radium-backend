package entity

import "errors"

var (
	SlideNotFoundErr = errors.New("слайды курса не найдены")
)

type (
	Slide struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		NameEng  string `json:"nameEng"`
		ModuleId uint   `json:"module"`
	}

	SlideDto struct {
		Id      uint   `json:"id"`
		NameEng string `json:"name_eng"`
		Name    string `json:"name"`
	}

	SlideRequest struct {
		CourseId      uint   `json:"course_id"`
		ModuleNameEng string `json:"module_name_eng"`
		Name          string `json:"name"`
	}

	SlidesRequest struct {
		CourseId      uint   `json:"course_id"`
		ModuleNameEng string `json:"module_name_eng"`
	}

	SlideSectionsRequest struct {
		// CourseId      uint   `json:"course_id"`
		// ModuleNameEng string `json:"module_name_eng"`
		SlideId uint `json:"slide_id"`
	}

	SlideSections struct {
		Id       uint          `json:"id"`
		Name     string        `json:"name"`
		NameEng  string        `json:"name_eng"`
		Sections []interface{} `json:"sections"`
	}
)
