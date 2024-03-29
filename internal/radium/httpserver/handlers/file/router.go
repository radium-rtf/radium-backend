package file

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/file/internal/upload"
	mwAuth "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	r.Route("/v1", func(r chi.Router) {
		r.Use(mwAuth.Required(useCases.Deps.TokenManager))
		r.Post("/upload", upload.New(useCases.File))
	})
}
