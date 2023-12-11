package app

import (
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/lib/session"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	pg "github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/email"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

func newDependencies(storage filestorage.Storage, cfg *config.Config, db *postgres.Postgres) usecase.Dependencies {
	repositories := pg.NewRepositories(db)
	tokenManager := auth.NewManager(cfg.Auth.SigningKey)
	passwordHasher := hash.NewPasswordHasher(cfg.Auth.PasswordSaltSha1, cfg.Auth.PasswordCostBcrypt)
	smtp := email.NewSMTPSender(cfg.Email, cfg.Password, cfg.Smtp.Host, cfg.Smtp.Port,
		"pkg/email/verification.html", cfg.Smtp.LengthVerificationCode, cfg.Smtp.Username)
	return usecase.Dependencies{
		Repos:                  repositories,
		TokenManager:           tokenManager,
		Storage:                storage,
		PasswordHasher:         passwordHasher,
		Smtp:                   smtp,
		Session:                session.New(tokenManager, cfg.AccessTokenTTL, cfg.RefreshTokenTTL),
		LengthVerificationCode: cfg.LengthVerificationCode,
	}
}
