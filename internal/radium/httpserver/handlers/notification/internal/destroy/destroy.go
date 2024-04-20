package destroy

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	resp "github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type deleter interface {
	Delete(ctx context.Context, ids []uuid.UUID, userId uuid.UUID) (int64, error)
}

// @Tags notification
// @Security ApiKeyAuth
// @Accept json
// @Param request body uuid.UUID true "массив uuid`ов уведомлений для удаления"
// @Success 200 {object} resp.Success       "количество удачных операций, чужие уведомления удалять нельзя"
// @Router /v1/notifications [delete]
func New(reader deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ids    uuid.UUIDs
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &ids)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		success, err := reader.Delete(ctx, ids, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := resp.Success{Success: success}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)
	}
}
