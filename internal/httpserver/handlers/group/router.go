package group

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group/internal/get"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group/internal/getbyid"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group/internal/invite"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

func New(r *chi.Mux, pg *db.Query, manager auth.TokenManager) {
	groupRepo := postgres.NewGroupRepo(pg)
	useCase := usecase.NewGroupUseCase(groupRepo)

	r.Route("/v1/group", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/{id}", getbyid.New(useCase))
			r.Get("/", get.New(useCase))

			r.Group(func(r chi.Router) {
				r.Use(mwAuth.Required(manager))
				r.Post("/", create.New(useCase))
				r.Patch("/invite/{inviteCode}", invite.New(useCase))
			})
		})
	})
}
