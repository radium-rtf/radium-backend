package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

// NewRouter
// @title       без юлерна
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(h *chi.Mux, pg *postgres.Postgres, storage filestorage.Storage, cfg *config.Config) {
	authUseCase := usecase.NewAuthUseCase(pg, cfg)
	accountUseCase := usecase.NewAccountUseCase(pg, cfg)
	courseUseCase := usecase.NewCourseUseCase(pg, storage)
	groupUseCase := usecase.NewGroupUseCase(pg)
	moduleUseCase := usecase.NewModuleUseCase(pg)
	slideUseCase := usecase.NewSlideUseCase(pg)

	h.Route("/v1", func(v1 chi.Router) {
		newAuthRoutes(v1, authUseCase)
		newAccountRoutes(v1, accountUseCase, cfg.SigningKey)

		newCourseRoutes(v1, courseUseCase, cfg.SigningKey)
		newModuleRoutes(v1, moduleUseCase, cfg.SigningKey)

		newGroupRoutes(v1, groupUseCase, cfg.SigningKey)
		newSlideRoutes(v1, slideUseCase, cfg.SigningKey)
	})
}
