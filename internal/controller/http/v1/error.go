package v1

import "github.com/pkg/errors"

type appError struct {
	err  error
	code int
}

func newAppError(err error, code int) *appError {
	return &appError{err: err, code: code}
}

func newAppErrorf(format string, err error, code int) *appError {
	return &appError{err: errors.Errorf(format, err), code: code}
}
