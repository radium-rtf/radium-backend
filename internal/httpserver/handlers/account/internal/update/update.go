package update

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"net/http"
)

type updater interface {
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

// @Tags  	    account
// @Accept      json
// @Produce     json
// @Security ApiKeyAuth
// @Param       request body User true "Данные для обновления"
// @Success     200
// @Router      /v1/account [patch]
func New(updater updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request User
			ctx     = r.Context()
			userId  = r.Context().Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		user := request.ToUser(userId)
		result, err := updater.UpdateUser(ctx, user)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, model.NewUser(result))
	}
}
