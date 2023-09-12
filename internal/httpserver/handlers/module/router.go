package module

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module/internal/destroy"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	useCase := useCases.Module

	r.Route("/v1/module", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(useCases.Deps.TokenManager))
			r.Post("/", create.New(useCase))
			r.Delete("/{id}", destroy.New(useCase))
		})
	})
}
