package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/answer"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/auth"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/course"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/file"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/page"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/review"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/role"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/section"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/teacher"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

// @title       без юлерна
// @version     1.0
// @host        localhost:8080
// @BasePath    /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(h *chi.Mux, useCases usecase.UseCases) {
	answer.New(h, useCases)

	course.New(h, useCases)
	module.New(h, useCases)
	page.New(h, useCases)
	section.New(h, useCases)

	group.New(h, useCases)

	teacher.New(h, useCases)
	review.New(h, useCases)

	auth.New(h, useCases)
	account.New(h, useCases)
	role.New(h, useCases)

	file.New(h, useCases)
}
