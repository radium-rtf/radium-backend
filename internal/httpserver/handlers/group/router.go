package group

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group/internal/get"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group/internal/getbyid"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group/internal/invite"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group/internal/report"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	useCase := useCases.Group

	r.Route("/v1/group", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/{id}", getbyid.New(useCase))
			r.Get("/", get.New(useCase))

			r.Group(func(r chi.Router) {
				r.Use(mwAuth.Required(useCases.Deps.TokenManager))
				r.Post("/", create.New(useCase))
				r.Get("/report/{groupId}/{courseId}", report.New(useCase))
				r.Patch("/invite/{inviteCode}", invite.New(useCase))
			})
		})
	})
}
