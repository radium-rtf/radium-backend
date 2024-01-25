package session

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/auth"
	"github.com/radium-rtf/radium-backend/internal/model"
	"time"
)

type Session struct {
	tokenManager          auth.TokenManager
	accessTTL, refreshTTL time.Duration
}

func New(tokenManager auth.TokenManager, accessTTL, refreshTTL time.Duration) Session {
	return Session{tokenManager: tokenManager, accessTTL: accessTTL, refreshTTL: refreshTTL}
}

func (s Session) Create(user *model.User) (entity.Session, model.Tokens, error) {
	var (
		tokens  = model.Tokens{User: user}
		session entity.Session
		err     error
	)

	expiresIn := time.Now().Add(s.accessTTL)
	tokens.ExpiresIn = expiresIn
	tokens.AccessToken, err = s.tokenManager.NewJWT(user, expiresIn)
	if err != nil {
		return session, tokens, err
	}

	tokens.RefreshToken = s.tokenManager.NewRefreshToken()
	if err != nil {
		return session, tokens, err
	}

	session = entity.Session{
		RefreshToken: tokens.RefreshToken,
		ExpiresIn:    time.Now().Add(s.refreshTTL),
		UserId:       user.Id,
	}

	return session, tokens, err
}

func (s Session) Refresh(user *model.User, refreshToken uuid.UUID) (model.Tokens, time.Time, error) {
	var (
		tokens model.Tokens
		err    error
	)

	tokens.User = user
	expiresIn := time.Now().Add(s.accessTTL)
	tokens.ExpiresIn = expiresIn
	tokens.RefreshToken = refreshToken

	tokens.AccessToken, err = s.tokenManager.NewJWT(user, expiresIn)
	if err != nil {
		return tokens, time.Time{}, err
	}

	return tokens, time.Now().Add(s.refreshTTL), err
}
