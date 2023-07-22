package page

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/page/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/page/internal/destroy"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/page/internal/getbyid"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

func New(r *chi.Mux, pg *db.Query, manager auth.TokenManager) {

	pageRepo := postgres.NewPageRepo(pg)
	answerRepo := postgres.NewAnswerRepo(pg)

	useCase := usecase.NewPageUseCase(pageRepo, answerRepo)

	r.Route("/v1/page", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mwAuth.UserId(manager))
			r.Get("/{id}", getbyid.New(useCase))
		})

		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(manager))
			r.Post("/", create.New(useCase))
			r.Delete("/{id}", destroy.New(useCase))
		})
	})
}
