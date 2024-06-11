package group

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/group/internal/create"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/group/internal/get"
	"github.com/radium-rtf/radium-backend/internal/wave/httpserver/handlers/group/internal/modify"
	"github.com/radium-rtf/radium-backend/internal/wave/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Group
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/group", func(r chi.Router) {
		r.Use(auth.Required(tokenManager))
		r.Post("/create", create.NewGroup(useCase))
		r.Route("/member", func(r chi.Router) {
			r.Post("/", modify.NewAdd(useCase))
			r.Delete("/", modify.NewRemove(useCase))
		})
		r.Get("/{chatId}", get.New(useCase))
	})
}
