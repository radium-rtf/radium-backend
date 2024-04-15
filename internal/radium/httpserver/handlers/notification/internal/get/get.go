package get

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	entity "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type getter interface {
	Get(ctx context.Context, userId uuid.UUID) ([]entity.Notification, error)
}

// @Tags notification
// @Security ApiKeyAuth
// @Accept json
// @Success 200 {object} model.Notification " "
// @Router /v1/notification [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		notifications, err := getter.Get(ctx, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewNotifications(notifications)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)
	}
}
