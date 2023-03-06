package v1

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"net/http"
	"regexp"
	"strings"
)

type authRoutes struct {
	uc           usecase.AuthUseCase
	emailPattern *regexp.Regexp
}

func authRequired(signingKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		manager, err := auth.NewManager(signingKey)
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if err != nil {
				writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			token := strings.Split(request.Header.Get("Authorization"), " ")[1]
			userId, err := manager.Parse(token)
			if err != nil {
				writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}
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
	err = json.NewEncoder(w).Encode(in)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
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
	err = json.NewEncoder(w).Encode(tokens)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
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
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}
