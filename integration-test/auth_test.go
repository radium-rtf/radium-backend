package main

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	e := httpexpect.Default(t, basePath)
	tokens := signIn(e)
	if t.Failed() {
		return
	}

	refresh(e, tokens)
}

func signIn(e *httpexpect.Expect) *model.Tokens {
	type signIn struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json := signIn{Email: userEmail, Password: userPassword}

	var tokens model.Tokens
	e.POST("/v1/auth/signin").
		WithJSON(json).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Decode(&tokens)

	return &tokens
}

func refresh(e *httpexpect.Expect, tokens *model.Tokens) {
	type refresh struct {
		RefreshToken uuid.UUID `json:"refreshToken"`
	}
	json := refresh{RefreshToken: tokens.RefreshToken}

	var refreshTokens model.Tokens
	e.POST("/v1/auth/refresh").
		WithJSON(json).
		Expect().Status(http.StatusOK).
		JSON().
		Decode(&refreshTokens)
}

func newHttpExpectWithAuth(t *testing.T) *httpexpect.Expect {
	e := httpexpect.Default(t, basePath)
	tokens := signIn(e)
	if t.Failed() {
		return nil
	}
	auth := e.Builder(func(request *httpexpect.Request) {
		request.WithHeader("Authorization", "Bearer "+tokens.AccessToken)
	})
	return auth
}
