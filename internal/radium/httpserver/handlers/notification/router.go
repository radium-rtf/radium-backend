package module

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/notification/internal/destroy"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/notification/internal/get"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/notification/internal/read"
	mwAuth "github.com/radium-rtf/radium-backend/internal/radium/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
)

func New(r chi.Router, useCases usecase.UseCases) {
	useCase := useCases.Notification
	tokenManager := useCases.Deps.TokenManager

	r.Route("/v1/notification", func(r chi.Router) {

		r.Use(mwAuth.Required(tokenManager))

		r.Get("/", get.New(useCase))
	})

	r.Route("/v1/notifications", func(r chi.Router) {

		r.Use(mwAuth.Required(tokenManager))

		r.Delete("/", destroy.New(useCase))
		r.Patch("/read", read.New(useCase))
	})
}
