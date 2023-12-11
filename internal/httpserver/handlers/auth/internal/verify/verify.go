package verify

import (
	"context"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/lib/decode"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type verify interface {
	VerifyEmail(ctx context.Context, email, verificationCode string) (model.Tokens, error)
}

// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param       request body Request true " "
// @Success     201 {object} model.Tokens
// @Router      /v1/auth/verify [post]
func New(verify verify) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Request
			ctx     = r.Context()
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
		}

		tokens, err := verify.VerifyEmail(ctx, request.Email, request.VerificationCode)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, tokens)
	}
}
