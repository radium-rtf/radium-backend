package repoerr

import "errors"

var (
	SessionIsExpired = errors.New("сессия истекла")
	SessionNotFound  = errors.New("сессия не найдена")
)
