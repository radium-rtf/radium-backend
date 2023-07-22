package answer

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/answer/internal/create"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

func New(r *chi.Mux, pg *db.Query, manager auth.TokenManager) {
	sectionRepo := postgres.NewSectionRepo(pg)
	answerRepo := postgres.NewAnswerRepo(pg)

	useCase := usecase.NewAnswerUseCase(sectionRepo, answerRepo)

	r.Route("/v1/answer", func(r chi.Router) {
		r.Use(mwAuth.Required(manager))
		r.Post("/", create.New(useCase))
	})
}
