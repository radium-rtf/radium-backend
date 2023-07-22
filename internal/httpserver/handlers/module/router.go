package module

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module/internal/destroy"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

func New(r *chi.Mux, pg *db.Query, manager auth.TokenManager) {

	moduleRepo := postgres.NewModuleRepo(pg)
	useCase := usecase.NewModuleUseCase(moduleRepo)
	r.Route("/v1/module", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(manager))
			r.Post("/", create.New(useCase))
			r.Delete("/{id}", destroy.New(useCase))
		})

	})

}
