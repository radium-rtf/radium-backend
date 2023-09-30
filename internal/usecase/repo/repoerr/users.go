package repoerr

import "errors"

var (
	UserNotFound      = errors.New("пользователь не найден")
	UserAlreadyExists = errors.New("пользователь уже существует")
)
