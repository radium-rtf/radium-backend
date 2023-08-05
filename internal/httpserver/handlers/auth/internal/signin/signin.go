package signin

import (
	"context"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type signIn interface {
	SignIn(ctx context.Context, email, password string) (model.Tokens, error)
}

// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param       request body SignIn true "SignIn"
// @Success     200 {object} model.Tokens
// @Router      /v1/auth/signin [post]
func New(in signIn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx     = r.Context()
			request SignIn
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		tokens, err := in.SignIn(ctx, request.Email, request.Password)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, tokens)
	}
}
