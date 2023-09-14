package section

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/section/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/section/internal/destroy"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	useCase := useCases.Section
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/section", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(tokenManager))
			r.Delete("/{id}", destroy.New(useCase))

			r.Group(func(r chi.Router) {
				r.Use(role.Author(tokenManager))
				r.Post("/", create.New(useCase))
			})
		})
	})
}
