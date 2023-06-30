package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager struct {
	signingKey string
}

func NewManager(signingKey string) (TokenManager, error) {
	return TokenManager{signingKey: signingKey}, nil
}

func (m TokenManager) NewJWT(id uuid.UUID, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		Subject:   id.String(),
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m TokenManager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func (m TokenManager) NewRefreshToken() (string, error) {
	UUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return UUID.String(), nil
}

func (m TokenManager) ExtractUserId(tokenHeader []string) (uuid.UUID, error) {
	if len(tokenHeader) == 1 {
		return uuid.UUID{}, errors.New("empty or corrupted Authorization header")
	}
	token := tokenHeader[1]
	userId, err := m.Parse(token)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.MustParse(userId), nil
}
