package course

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/create"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/createlink"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/destroy"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/destroylink"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/get"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/getbyid"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/getbyslug"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/join"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/publish"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/search"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/course/internal/update"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	role2 "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Course
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/course", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(auth.UserId(tokenManager))
			r.Get("/", get.New(useCase))
			r.Get("/{courseId}", getbyid.New(useCase))
			r.Get("/slug/{slug}", getbyslug.New(useCase))
			r.Get("/search", search.New(useCase))
		})

		r.Group(func(r chi.Router) {
			r.Use(auth.Required(tokenManager))

			r.Patch("/join/{courseId}", join.New(useCase))

			r.Group(func(r chi.Router) {
				r.Use(role2.Author(tokenManager))
				r.Post("/", create.New(useCase))
				r.Patch("/publish/{id}", publish.New(useCase))
				r.Delete("/{id}", destroy.New(useCase))
			})

			r.Group(func(r chi.Router) {
				r.Use(role2.CanEditCourse(tokenManager))
				r.Put("/{courseId}", update.New(useCase))
				r.Delete("/link/{id}", destroylink.New(useCase))
				r.Post("/link/{courseId}", createlink.New(useCase))
			})
		})
	})
}
