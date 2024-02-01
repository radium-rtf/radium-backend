package refresh

import (
	"context"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type refresh interface {
	RefreshToken(ctx context.Context, refreshToken uuid.UUID) (model.Tokens, error)
}

// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param       request body Refresh true "RefreshToken"
// @Success     200 {object} model.Tokens
// @Router      /v1/auth/refresh [post]
func New(refresh refresh) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Refresh
			ctx     = r.Context()
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		token, err := refresh.RefreshToken(ctx, request.RefreshToken)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, token)
	}
}
