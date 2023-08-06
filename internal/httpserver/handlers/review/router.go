package review

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/review/internal/create"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

func New(r *chi.Mux, pg *db.Query, manager auth.TokenManager) {
	reviewRepo := postgres.NewReviewRepo(pg)

	useCase := usecase.NewReviewUseCase(reviewRepo)

	r.Route("/v1/review", func(r chi.Router) {
		r.Use(mwAuth.Required(manager))
		r.Post("/", create.NewReview(useCase))
	})
}
