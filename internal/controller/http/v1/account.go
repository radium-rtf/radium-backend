package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
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
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}
