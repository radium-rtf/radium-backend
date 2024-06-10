package role

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/role/internal/deletecoauthor"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/role/internal/postadmin"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/role/internal/postauthor"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/role/internal/postcoauthor"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/role/internal/postteacher"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/role/internal/updaterole"
	mwAuth "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Role

	r.Route("/v1/role", func(r chi.Router) {
		r.Use(mwAuth.Required(useCases.Deps.TokenManager))
		r.Post("/teacher", postteacher.New(useCase))
		r.Post("/author", postauthor.New(useCase))

		r.Group(func(r chi.Router) {
			r.Use(role.Admin(useCases.Deps.TokenManager))
			r.Post("/admin", postadmin.New(useCase))
			r.Patch("/{id}", updaterole.New(useCase))
		})

		r.Group(func(r chi.Router) {
			r.Use(role.Author(useCases.Deps.TokenManager))
			r.Post("/coauthor", postcoauthor.New(useCase))
			r.Delete("/coauthor/{id}/{courseId}", deletecoauthor.New(useCase))
		})
	})
}
