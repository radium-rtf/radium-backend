package page

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/page/internal/create"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/page/internal/destroy"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/page/internal/getbyid"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/page/internal/getbyslug"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/page/internal/order"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/page/internal/update"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	role2 "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Page
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/page", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(auth.UserId(tokenManager))
			r.Get("/{id}", getbyid.New(useCase))
			r.Get("/slug/{slug}", getbyslug.New(useCase))
		})

		r.Group(func(r chi.Router) {
			r.Use(auth.Required(tokenManager))

			r.Group(func(r chi.Router) {
				r.Use(role2.Author(tokenManager))
				r.Post("/", create.New(useCase))
			})

			r.Group(func(r chi.Router) {
				r.Use(role2.CanEditCourse(tokenManager))
				r.Delete("/{id}", destroy.New(useCase))
				r.Put("/{pageId}", update.New(useCase))
				r.Patch("/{id}/order", order.New(useCase))
			})
		})
	})
}
