package teacher

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/teacher/internal/courses"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/teacher/internal/create"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Teacher
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/teacher", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(tokenManager))
			r.Post("/", create.New(useCase))

			r.Group(func(r chi.Router) {
				r.Use(role.Teacher(tokenManager))
				r.Get("/courses", courses.New(useCase))
			})
		})
	})
}
