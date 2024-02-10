package destroy

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
	"strconv"
)

type deleter interface {
	Delete(ctx context.Context, id, editorId uuid.UUID, isSoft bool) error
}

// @Tags course
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path      string  true  "id"
// @Param   is_soft     query     boolean    false  "по умолчанию soft"
// @Success 200
// @Router /v1/course/{id} [delete]
func New(deleter deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			id     uuid.UUID
			isSoft bool
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		isSoft, err = strconv.ParseBool(r.URL.Query().Get("is_soft"))
		if err != nil {
			isSoft = true
		}

		err = deleter.Delete(ctx, id, userId, isSoft)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
