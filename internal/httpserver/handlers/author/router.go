package author

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/author/internal/courses"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Author

	r.Route("/v1/author", func(r chi.Router) {
		r.Use(mwAuth.Required(useCases.Deps.TokenManager))

		r.Group(func(r chi.Router) {
			r.Use(role.CanEditCourse(useCases.Deps.TokenManager))
			r.Get("/courses", courses.New(useCase))
		})
	})
}
