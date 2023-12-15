package review

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/review/internal/create"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Review
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/review", func(r chi.Router) {
		r.Use(mwAuth.Required(tokenManager))
		r.Use(role.Teacher(tokenManager))
		r.Post("/", create.NewReview(useCase))
	})
}
