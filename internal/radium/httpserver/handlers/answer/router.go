package answer

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/answer/internal/create"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/answer/internal/getbygroup"
	mwAuth "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/role"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	r.Route("/v1/answer", func(r chi.Router) {
		r.Use(mwAuth.Required(useCases.Deps.TokenManager))
		r.Post("/", create.New(useCases.Answer))
	})

	r.Route("/v1/answers", func(r chi.Router) {
		r.Use(mwAuth.Required(useCases.Deps.TokenManager))
		r.Use(role.Teacher(useCases.Deps.TokenManager))
		r.Get("/group/{groupId}", getbygroup.NewAnswer(useCases.Group))
	})
}
