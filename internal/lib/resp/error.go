package resp

import (
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/pkg/validator"
	"net/http"
)

type respError struct {
	Type    string `json:"type"`
	Message string `json:"message"`

	ValidationError *validator.ValidationError `json:"validationError"`
}

func Error(r *http.Request, w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	switch err.(type) {
	case validator.ValidationError:
		ValidationError(r, w, err.(validator.ValidationError))
	default:
		RawError(r, w, err)
	}
}

func ValidationError(r *http.Request, w http.ResponseWriter, err validator.ValidationError) {
	resp := respError{Type: "validation", ValidationError: &err}
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, resp)
}

func RawError(r *http.Request, w http.ResponseWriter, err error) {
	resp := respError{Type: "raw", Message: err.Error()}
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, resp)
}
