package role

import (
	"github.com/go-chi/chi/v5"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"

	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/role/internal/postauthor"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/role/internal/postcoauthor"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/role/internal/postteacher"

	"github.com/radium-rtf/radium-backend/internal/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	useCase := useCases.Role

	r.Route("/v1/role", func(r chi.Router) {
		r.Use(mwAuth.Required(useCases.Deps.TokenManager))
		r.Post("/teacher", postteacher.New(useCase))
		r.Post("/author", postauthor.New(useCase))

		r.Group(func(r chi.Router) {
			r.Use(role.Author(useCases.Deps.TokenManager))
			r.Post("/coauthor", postcoauthor.New(useCase))
		})
	})
}
