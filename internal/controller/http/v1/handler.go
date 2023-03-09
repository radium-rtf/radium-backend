package v1

import (
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request) *appError

func (f handler) HTTP(w http.ResponseWriter, r *http.Request) {
	if e := f(w, r); e != nil {
		w.WriteHeader(e.code)
		_, _ = w.Write([]byte(e.err.Error()))
	}
}
