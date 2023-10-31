package author

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/author/internal/courses"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	useCase := useCases.Author

	r.Route("/v1/author", func(r chi.Router) {
		r.Use(mwAuth.Required(useCases.Deps.TokenManager))
		r.Use(role.Author(useCases.Deps.TokenManager))
		r.Get("/courses", courses.New(useCase))
	})
}
