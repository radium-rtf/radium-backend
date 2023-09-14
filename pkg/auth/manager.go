package auth

import (
	"errors"
	"fmt"
	"github.com/radium-rtf/radium-backend/internal/model"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager struct {
	signingKey string
}

const (
	isTeacher = "isTeacher"
	isAuthor  = "isAuthor"
)

func NewManager(signingKey string) TokenManager {
	return TokenManager{signingKey: signingKey}
}

func (m TokenManager) NewJWT(user model.User, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     jwt.NewNumericDate(expiresAt),
		"sub":     user.Id.String(),
		isTeacher: slices.Contains(user.Roles, model.TeacherRole),
		isAuthor:  slices.Contains(user.Roles, model.AuthorRole),
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m TokenManager) parse(accessToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error get user claims from token")
	}

	return claims, nil
}

func (m TokenManager) NewRefreshToken() (string, error) {
	UUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return UUID.String(), nil
}

func (m TokenManager) ExtractUserId(tokenHeader []string) (uuid.UUID, error) {
	claims, err := m.extractClaims(tokenHeader)
	if err != nil {
		return uuid.UUID{}, err
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, err
	}

	return uuid.MustParse(sub), nil
}

func (m TokenManager) extractClaims(tokenHeader []string) (jwt.MapClaims, error) {
	token, err := m.getToken(tokenHeader)
	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, err := m.parse(token)
	if err != nil {
		return jwt.MapClaims{}, err
	}

	return claims, nil
}

func (m TokenManager) extractClaim(tokenHeader []string, key string) (any, error) {
	claims, err := m.extractClaims(tokenHeader)
	if err != nil {
		return false, err
	}

	claim, ok := claims[key]
	if !ok {
		return false, errors.New("invalid claim")
	}
	return claim, nil
}

func (m TokenManager) ExtractIsTeacher(tokenHeader []string) (bool, error) {
	claim, err := m.extractClaim(tokenHeader, isTeacher)
	if err != nil {
		return false, err
	}
	return claim.(bool), nil
}

func (m TokenManager) ExtractIsAuthor(tokenHeader []string) (bool, error) {
	claim, err := m.extractClaim(tokenHeader, isAuthor)
	if err != nil {
		return false, err
	}
	return claim.(bool), nil
}

func (m TokenManager) getToken(tokenHeader []string) (string, error) {
	if len(tokenHeader) != 2 || tokenHeader[0] != "Bearer" {
		return "", errors.New("empty or corrupted Authorization header: Bearer <token>")
	}
	return tokenHeader[1], nil
}
