package module

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module/internal/destroy"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module/internal/order"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module/internal/update"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Module
	tokenManager := useCases.Deps.TokenManager
	r.Route("/v1/module", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(tokenManager))

			r.Group(func(r chi.Router) {
				r.Use(role.Author(tokenManager))
				r.Post("/", create.New(useCase))
			})

			r.Group(func(r chi.Router) {
				r.Use(role.CanEditCourse(tokenManager))
				r.Delete("/{id}", destroy.New(useCase))
				r.Put("/{moduleId}", update.New(useCase))
				r.Patch("/{id}/order", order.New(useCase))
			})
		})
	})
}
