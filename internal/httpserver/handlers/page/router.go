package page

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/page/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/page/internal/destroy"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/page/internal/getbyid"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	useCase := useCases.Page
	answerUseCase := useCases.Answer

	r.Route("/v1/page", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mwAuth.UserId(useCases.Deps.TokenManager))
			r.Get("/{id}", getbyid.New(useCase, answerUseCase))
		})

		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(useCases.Deps.TokenManager))
			r.Post("/", create.New(useCase))
			r.Delete("/{id}", destroy.New(useCase))
		})
	})
}
