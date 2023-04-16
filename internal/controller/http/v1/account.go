package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"net/http"
)

type accountRoutes struct {
	uc usecase.AccountUseCase
}

func newAccountRoutes(h chi.Router, useCase usecase.AccountUseCase, signingKey string) {
	routes := accountRoutes{uc: useCase}
	h.Route("/account", func(r chi.Router) {
		r.Use(authRequired(signingKey))
		r.Get("/", handler(routes.account).HTTP)
		r.Patch("/name", handler(routes.name).HTTP)
		r.Patch("/password", handler(routes.password).HTTP)
	})
}

// @Tags  	    account
// @Accept      json
// @Produce     json
// @Security ApiKeyAuth
// @Success     200 {object} entity.UserDto
// @Router      /account [get]
func (r accountRoutes) account(w http.ResponseWriter, request *http.Request) *appError {
	userId := request.Context().Value("userId")
	user, err := r.uc.Account(request.Context(), userId.(string))
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, user)
	return nil
}

// @Tags  	    account
// @Accept      json
// @Produce     json
// @Security ApiKeyAuth
// @Param       request body entity.UserName true "Новое имя"
// @Success     200
// @Router      /account/name [patch]
func (r accountRoutes) name(w http.ResponseWriter, request *http.Request) *appError {
	var name entity.UserName
	err := render.DecodeJSON(request.Body, &name)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	userId := request.Context().Value("userId").(string)
	err = r.uc.UpdateName(request.Context(), userId, name)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

// @Tags  	    account
// @Accept      json
// @Produce     json
// @Security ApiKeyAuth
// @Param       request body entity.PasswordUpdate true "PasswordUpdate"
// @Success     200
// @Router      /account/password [patch]
func (r accountRoutes) password(w http.ResponseWriter, request *http.Request) *appError {
	var passwordUpdate entity.PasswordUpdate
	err := render.DecodeJSON(request.Body, &passwordUpdate)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	userId := request.Context().Value("userId").(string)
	err = r.uc.UpdatePassword(request.Context(), userId, passwordUpdate)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}
