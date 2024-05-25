package get

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type tokenGetter interface {
	GetClientToken(ctx context.Context, userId uuid.UUID) (string, error)
}

// @Tags user
// @Security ApiKeyAuth
// @Success      200   {object} []model.CentrifugoToken        " "
// @Router       /v1/user/token [get]
func NewToken(getter tokenGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		token, err := getter.GetClientToken(ctx, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewCentrifugoToken(token)

		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
