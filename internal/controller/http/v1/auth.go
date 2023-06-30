package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/auth"
)

type authRoutes struct {
	uc           usecase.AuthUseCase
	emailPattern *regexp.Regexp
}

func authRequired(signingKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		manager, err := auth.NewManager(signingKey)
		if err != nil {
			panic(err)
		}
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")
			userId, err := manager.ExtractUserId(tokenHeader)
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte(err.Error()))
			}
			ctx := context.WithValue(request.Context(), "userId", userId)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}

func authToken(signingKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		manager, err := auth.NewManager(signingKey)
		if err != nil {
			panic(err)
		}
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")
			userId, _ := manager.ExtractUserId(tokenHeader)
			ctx := context.WithValue(request.Context(), "userId", userId)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}

func newAuthRoutes(h chi.Router, useCase usecase.AuthUseCase) {
	emailPattern, _ := regexp.Compile("[a-zA-Z.]@urfu.(me|ru)")
	routes := authRoutes{uc: useCase, emailPattern: emailPattern}
	h.Route("/auth", func(r chi.Router) {
		r.Post("/signIn", handler(routes.signIn).HTTP)
		r.Post("/signUp", handler(routes.signUp).HTTP)
		r.Post("/refresh", handler(routes.refresh).HTTP)
		r.Post("/verify", handler(routes.verify).HTTP)
	})
}

// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param       request body entity.SignIn true "SignIn"
// @Success     200 {object} entity.Tokens
// @Router      /auth/signIn [post]
func (r authRoutes) signIn(w http.ResponseWriter, request *http.Request) *appError {
	signIn := entity.SignIn{}
	err := json.NewDecoder(request.Body).Decode(&signIn)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	in, err := r.uc.SignIn(request.Context(), signIn)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, in)
	return nil
}

// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param       request body entity.SignUp true "SignUp"
// @Success     201 {object} entity.Tokens
// @Router      /auth/signUp [post]
func (r authRoutes) signUp(w http.ResponseWriter, request *http.Request) *appError {
	signUp := entity.SignUp{}
	err := json.NewDecoder(request.Body).Decode(&signUp)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	if ok := r.emailPattern.MatchString(signUp.Email); !ok {
		return newAppError(errors.New("email должен удовлетвориять [a-zA-Z.]@urfu.(me|ru)"), http.StatusBadRequest)
	}
	tokens, err := r.uc.SignUp(request.Context(), signUp)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, tokens)
	return nil
}

// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param       request body entity.RefreshToken true "RefreshToken"
// @Success     200 {object} entity.Tokens
// @Router      /auth/refresh [post]
func (r authRoutes) refresh(w http.ResponseWriter, request *http.Request) *appError {
	var refreshToken entity.RefreshToken
	err := json.NewDecoder(request.Body).Decode(&refreshToken)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	token, err := r.uc.RefreshToken(request.Context(), refreshToken.RefreshToken)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, token)
	return nil
}

func (r authRoutes) verify(w http.ResponseWriter, request *http.Request) *appError {
	var verificationCode entity.VerificationCode
	err := json.NewDecoder(request.Body).Decode(&verificationCode)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	result, err := r.uc.VerifyEmail(request.Context(), verificationCode.VerificationCode)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, result)
	return nil
}
