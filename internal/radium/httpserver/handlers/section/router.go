package section

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/section/internal/answer"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/section/internal/create"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/section/internal/destroy"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/section/internal/order"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/section/internal/update"
	mwAuth "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	role2 "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Section
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/section", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(tokenManager))

			r.Group(func(r chi.Router) {
				r.Use(role2.Author(tokenManager))
				r.Post("/", create.New(useCase))
			})

			r.Group(func(r chi.Router) {
				r.Use(role2.CanEditCourse(tokenManager))
				r.Delete("/{id}", destroy.New(useCase))
				r.Put("/{id}", update.New(useCase, useCases.Answer))
				r.Patch("/{id}/order", order.New(useCase))
				r.Get("/{id}/answer", answer.New(useCase))
			})
		})
	})
}
