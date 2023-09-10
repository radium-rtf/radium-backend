package role

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/role/internal/postteacher"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	role := useCases.Role

	r.Route("/v1/role", func(r chi.Router) {
		r.Post("/teacher", postteacher.New(role))
	})
}
