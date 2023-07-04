package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

// NewRouter
// @title       без юлерна
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(h *chi.Mux, pg *db.Query, storage filestorage.Storage, cfg *config.Config) {

	authUseCase := usecase.NewAuthUseCase(pg, cfg)
	accountUseCase := usecase.NewAccountUseCase(pg, cfg)
	courseUseCase := usecase.NewCourseUseCase(pg)
	fileUseCase := usecase.NewFileUseCase(storage)
	// _ = usecase.NewGroupUseCase(pg)
	answerUseCase := usecase.NewAnswerUseCase(pg)
	moduleUseCase := usecase.NewModuleUseCase(pg)
	pageUseCase := usecase.NewPageUseCase(pg)
	sectionUseCase := usecase.NewSectionUseCase(pg)

	h.Route("/v1", func(v1 chi.Router) {
		v1.Mount("/auth", newAuthRoutes(authUseCase))
		v1.Mount("/account", newAccountRoutes(accountUseCase, cfg.SigningKey))

		v1.Mount("/course", newCourseRoutes(courseUseCase, cfg.SigningKey))

		v1.Mount("/upload", newFileRoutes(fileUseCase, cfg.SigningKey))
		v1.Mount("/module", newModuleRoutes(moduleUseCase, cfg.SigningKey))

		// 	newGroupRoutes(v1, groupUseCase, cfg.SigningKey)
		v1.Mount("/answer", newAnswerRoutes(answerUseCase, cfg.SigningKey))
		v1.Mount("/page", newPageRoutes(pageUseCase, cfg.SigningKey))
		v1.Mount("/section", newSectionRoutes(sectionUseCase, cfg.SigningKey))
	})
}
