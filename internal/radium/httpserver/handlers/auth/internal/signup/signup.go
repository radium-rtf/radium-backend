package signup

import (
	"context"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type signUp interface {
	SignUp(ctx context.Context, user *entity.User) (*entity.UnverifiedUser, error)
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

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		user := request.toUser()
		unverifiedUser, err := signUp.SignUp(ctx, user)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		response := Response{
			Email:     unverifiedUser.Email,
			ExpiresAt: unverifiedUser.ExpiresAt,
		}
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, response)
	}
}
