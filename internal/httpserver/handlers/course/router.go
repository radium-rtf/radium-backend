package course

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/course/internal/create"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/course/internal/destroy"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/course/internal/get"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/course/internal/getbyid"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/course/internal/getbyslug"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/course/internal/join"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

func New(r *chi.Mux, useCases usecase.UseCases) {
	useCase := useCases.Course
	answerUseCase := useCases.Answer

	r.Route("/v1/course", func(r chi.Router) {
		r.Get("/", get.New(useCase))

		r.Group(func(r chi.Router) {
			r.Use(mwAuth.UserId(useCases.Deps.TokenManager))
			r.Get("/{courseId}", getbyid.New(useCase, answerUseCase))
			r.Get("/slug/{slug}", getbyslug.New(useCase, answerUseCase))
		})

		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(useCases.Deps.TokenManager))
			r.Post("/", create.New(useCase))
			r.Patch("/join/{courseId}", join.New(useCase))
			r.Delete("/{id}", destroy.New(useCase))
		})
	})
}
