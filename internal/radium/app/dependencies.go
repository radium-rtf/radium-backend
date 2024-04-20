package app

import (
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/session"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase"
	pg "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/email"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

func newDependencies(storage filestorage.Storage, cfg *config.Config, db *postgres.Postgres) usecase.Dependencies {
	authCFG := cfg.Radium.Auth

	repositories := pg.NewRepositories(db)
	tokenManager := auth.NewManager(cfg.JWT.SigningKey)
	passwordHasher := hash.NewPasswordHasher(authCFG.PasswordSaltSha1, authCFG.PasswordCostBcrypt)
	smtp := email.NewSMTPSender(cfg.Smtp.Email, cfg.Smtp.Password, cfg.Smtp.Host, cfg.Smtp.Port,
		authCFG.LengthVerificationCode, cfg.Smtp.Username)

	return usecase.Dependencies{
		Repos:                  repositories,
		TokenManager:           tokenManager,
		Storage:                storage,
		PasswordHasher:         passwordHasher,
		Smtp:                   smtp,
		Session:                session.New(tokenManager, authCFG.AccessTokenTTL, authCFG.RefreshTokenTTL),
		LengthVerificationCode: authCFG.LengthVerificationCode,
	}
}
