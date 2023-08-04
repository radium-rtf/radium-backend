package refresh

import (
	"context"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type refresh interface {
	RefreshToken(ctx context.Context, refreshToken string) (model.Tokens, error)
}

// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param       request body Request true "RefreshToken"
// @Success     200 {object} model.Tokens
// @Router      /v1/auth/refresh [post]
func New(refresh refresh) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Request
			ctx     = r.Context()
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		token, err := refresh.RefreshToken(ctx, request.RefreshToken)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, token)
	}
}