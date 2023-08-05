package updatepass

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

type updater interface {
	UpdatePassword(ctx context.Context, id uuid.UUID, current, new string) error
}

// @Tags  	    account
// @Accept      json
// @Produce     json
// @Security ApiKeyAuth
// @Param       request body Password true "Password"
// @Success     200
// @Router      /v1/account/password [patch]
func New(updater updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Password
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		err = updater.UpdatePassword(ctx, userId, request.Current, request.New)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
