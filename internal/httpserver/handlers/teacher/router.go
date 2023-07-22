package teacher

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/teacher/internal/courses"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/teacher/internal/create"
	mwAuth "github.com/radium-rtf/radium-backend/internal/httpserver/middleware/auth"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

func New(r *chi.Mux, pg *db.Query, manager auth.TokenManager) {
	teacherRepo := postgres.NewTeacherRepo(pg)
	useCase := usecase.NewTeacherUseCase(teacherRepo)

	r.Route("/v1/teacher", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mwAuth.Required(manager))
			r.Post("/", create.New(useCase))
			r.Get("/courses", courses.New(useCase))
		})
	})
}
