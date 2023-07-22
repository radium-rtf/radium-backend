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
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

func New(r *chi.Mux, pg *db.Query, manager auth.TokenManager) {
	courseRepo := postgres.NewCourseRepo(pg)
	useCase := usecase.NewCourseUseCase(courseRepo)

	r.Route("/v1/course", func(r chi.Router) {
		r.Get("/", get.New(useCase))
		r.Get("/{courseId}", getbyid.New(useCase))
		r.Get("/slug/{slug}", getbyslug.New(useCase))

		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(manager))
			r.Post("/", create.New(useCase))
			r.Patch("/join/{courseId}", join.New(useCase))
			r.Delete("/{id}", destroy.New(useCase))
		})
	})
}
