package usecase

import (
	"context"
	"errors"
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"strings"
	"time"
)

type AuthUseCase struct {
	userRepo              repo.UserRepo
	sessionRepo           repo.SessionRepo
	hasher                hash.PasswordHasher
	tokenManager          auth.TokenManager
	accessTTL, refreshTTL time.Duration
}

func NewAuthUseCase(pg *postgres.Postgres, cfg *config.Config) AuthUseCase {
	tokenManager, _ := auth.NewManager(cfg.SigningKey)
	passwordHasher := hash.NewSHA1Hasher(cfg.PasswordSalt)
	return AuthUseCase{
		userRepo: repo.NewUserRepo(pg), sessionRepo: repo.NewSessionRepo(pg),
		hasher: passwordHasher, tokenManager: tokenManager,
		accessTTL: cfg.Auth.AccessTokenTTL, refreshTTL: cfg.Auth.RefreshTokenTTL,
	}
}

func (uc AuthUseCase) SignIn(ctx context.Context, signIn entity.SignIn) (entity.Tokens, error) {
	var tokens entity.Tokens
	user, err := uc.userRepo.GetByEmail(ctx, signIn.Email)
	if err != nil {
		return tokens, err
	}
	if !uc.hasher.Equals(user.Password, signIn.Password) {
		return tokens, errors.New("неверный пароль")
	}
	return uc.createSession(ctx, user.Id)
}

func (uc AuthUseCase) SignUp(ctx context.Context, signIn entity.SignUp) (entity.Tokens, error) {
	var tokens entity.Tokens
	password, err := uc.hasher.Hash(signIn.Password)
	if err != nil {
		return tokens, err
	}
	signIn.Password = password
	username := strings.Split(signIn.Email, "@")[0]
	err = uc.userRepo.Create(ctx, signIn, username)
	if err != nil {
		return tokens, err
	}
	user, err := uc.userRepo.GetByEmail(ctx, signIn.Email)
	if err != nil {
		return tokens, err
	}
	return uc.createSession(ctx, user.Id)
}

func (uc AuthUseCase) RefreshToken(ctx context.Context, refreshToken string) (entity.Tokens, error) {
	user, err := uc.userRepo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return entity.Tokens{}, err
	}
	return uc.refreshSession(ctx, user.Id, refreshToken)
}

func (uc AuthUseCase) createSession(ctx context.Context, id uint) (entity.Tokens, error) {
	var (
		tokens entity.Tokens
		err    error
	)
	expiresIn := time.Now().Add(uc.accessTTL)
	tokens.ExpiresIn = expiresIn
	tokens.AccessToken, err = uc.tokenManager.NewJWT(id, expiresIn)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken, err = uc.tokenManager.NewRefreshToken()
	if err != nil {
		return tokens, err
	}
	session := entity.Session{
		RefreshToken: tokens.RefreshToken,
		ExpiresIn:    time.Now().Add(uc.refreshTTL),
		UserId:       id,
	}

	err = uc.sessionRepo.Create(ctx, session)
	return tokens, err
}

func (uc AuthUseCase) refreshSession(ctx context.Context, id uint, refreshToken string) (entity.Tokens, error) {
	var (
		tokens entity.Tokens
		err    error
	)
	expiresIn := time.Now().Add(uc.accessTTL)
	tokens.ExpiresIn = expiresIn
	tokens.RefreshToken = refreshToken
	tokens.AccessToken, err = uc.tokenManager.NewJWT(id, expiresIn)
	if err != nil {
		return tokens, err
	}

	err = uc.sessionRepo.Update(ctx, refreshToken, time.Now().Add(uc.refreshTTL))
	return tokens, err
}
