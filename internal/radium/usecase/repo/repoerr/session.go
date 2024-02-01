package repoerr

import "errors"

var (
	SessionIsExpired = errors.New("session is expired")
)
