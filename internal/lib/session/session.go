package session

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"time"
)

type Session struct {
	tokenManager          auth.TokenManager
	accessTTL, refreshTTL time.Duration
}

func New(tokenManager auth.TokenManager, accessTTL, refreshTTL time.Duration) Session {
	return Session{tokenManager: tokenManager, accessTTL: accessTTL, refreshTTL: refreshTTL}
}

func (s Session) Create(userId uuid.UUID) (entity.Session, model.Tokens, error) {
	var (
		tokens  model.Tokens
		session entity.Session
		err     error
	)

	expiresIn := time.Now().Add(s.accessTTL)
	tokens.ExpiresIn = expiresIn
	tokens.AccessToken, err = s.tokenManager.NewJWT(userId, expiresIn)
	if err != nil {
		return session, tokens, err
	}

	tokens.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return session, tokens, err
	}

	session = entity.Session{
		RefreshToken: tokens.RefreshToken,
		ExpiresIn:    time.Now().Add(s.refreshTTL),
		UserId:       userId,
	}

	return session, tokens, err
}

func (s Session) Refresh(id uuid.UUID, refreshToken string) (model.Tokens, time.Time, error) {
	var (
		tokens model.Tokens
		err    error
	)
	expiresIn := time.Now().Add(s.accessTTL)
	tokens.ExpiresIn = expiresIn
	tokens.RefreshToken = refreshToken

	tokens.AccessToken, err = s.tokenManager.NewJWT(id, expiresIn)
	if err != nil {
		return tokens, time.Time{}, err
	}

	return tokens, time.Now().Add(s.refreshTTL), err
}
