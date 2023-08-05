package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/account"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/answer"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/auth"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/course"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/file"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/group"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/module"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/page"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/review"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/section"
	"github.com/radium-rtf/radium-backend/internal/httpserver/handlers/teacher"
	session2 "github.com/radium-rtf/radium-backend/internal/lib/session"
	authutil "github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

// @title       без юлерна
// @version     1.0
// @host        localhost:8080
// @BasePath    /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(h *chi.Mux, pg *db.Query, storage filestorage.Storage, cfg *config.Config) {
	hasher := hash.NewSHA1Hasher(cfg.PasswordSalt)
	tokenManager := authutil.NewManager(cfg.SigningKey)
	session := session2.New(tokenManager, cfg.AccessTokenTTL, cfg.RefreshTokenTTL)

	answer.New(h, pg, tokenManager)

	course.New(h, pg, tokenManager)
	module.New(h, pg, tokenManager)
	page.New(h, pg, tokenManager)
	section.New(h, pg, tokenManager)

	group.New(h, pg, tokenManager)

	teacher.New(h, pg, tokenManager)
	review.New(h, pg, tokenManager)

	auth.New(h, pg, hasher, session)
	account.New(h, pg, tokenManager, hasher)

	file.New(h, storage, tokenManager)
}
