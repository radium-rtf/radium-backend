package group

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/group/internal/addstudent"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/group/internal/create"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/group/internal/get"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/group/internal/getbyid"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/group/internal/invite"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/group/internal/report"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/group/internal/update"
	mwAuth "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Group

	r.Route("/v1/group", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/{id}", getbyid.New(useCase))
			r.Get("/", get.New(useCase))

			r.Group(func(r chi.Router) {
				r.Use(mwAuth.Required(useCases.Deps.TokenManager))
				r.Post("/", create.New(useCase))
				r.Patch("/invite/{inviteCode}", invite.New(useCase))

				//r.Group(func(r chi.Router) {
				//r.Use(role.Author(useCases.Deps.TokenManager))
				r.Patch("/{groupId}", update.New(useCase))
				r.Patch("/{groupId}/add/{studentId}", addstudent.New(useCase))
				//r.Delete("/{groupId}", destroy.New(useCase))
				//})

				r.Group(func(r chi.Router) {
					r.Use(role.Teacher(useCases.Deps.TokenManager))
					r.Get("/report/{groupId}/{courseId}", report.New(useCase))
				})
			})
		})
	})
}
