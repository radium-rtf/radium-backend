package section

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/section/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/section/internal/destroy"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

func New(r *chi.Mux, pg *db.Query, manager auth.TokenManager) {
	sectionRepo := postgres.NewSectionRepo(pg)
	useCase := usecase.NewSectionUseCase(sectionRepo)

	r.Route("/v1/section", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(manager))
			r.Post("/", create.New(useCase))
			r.Delete("/{id}", destroy.New(useCase))
		})
	})
}
