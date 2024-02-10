package usecase

import "errors"

var (
	cantEditCourse = errors.New("только авторы и соавторы могут редактировать курс")
)
