package signup

import (
	"context"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/validator"
	"net/http"
)

type signUp interface {
	SignUp(ctx context.Context, user *entity.User) (model.Tokens, error)
}

// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param       request body SignUp true "SignUp"
// @Success     201 {object} model.Tokens
// @Router      /v1/auth/signup [post]
func New(signUp signUp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request SignUp
			ctx     = r.Context()
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		err = validator.Struct(request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		user := request.toUser()
		tokens, err := signUp.SignUp(ctx, user)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, tokens)
	}
}
