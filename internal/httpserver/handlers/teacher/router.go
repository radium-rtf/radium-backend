package teacher

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/teacher/internal/courses"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/teacher/internal/create"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	useCase := useCases.Teacher

	r.Route("/v1/teacher", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(useCases.Deps.TokenManager))
			r.Post("/", create.New(useCase))
			r.Get("/courses", courses.New(useCase))
		})
	})
}
