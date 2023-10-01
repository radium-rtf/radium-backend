package repoerr

import "errors"

var (
	CourseNotFound      = errors.New("курс не найден")
	CourseAlreadyExists = errors.New("курс уже существует")
)
