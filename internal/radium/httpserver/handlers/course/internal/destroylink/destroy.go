package destroylink

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type deleter interface {
	DeleteLink(ctx context.Context, id, editorId uuid.UUID) error
}

// @Tags course-link
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path      string  true  "id"
// @Success 200
// @Router /v1/course/link/{id} [delete]
func New(deleter deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			id     uuid.UUID
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		err = deleter.DeleteLink(ctx, id, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
