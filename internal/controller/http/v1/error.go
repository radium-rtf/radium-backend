package v1

type appError struct {
	err  error
	code int
}

func newAppError(err error, code int) *appError {
	return &appError{err: err, code: code}
}
