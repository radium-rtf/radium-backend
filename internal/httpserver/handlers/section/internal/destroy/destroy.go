package destroy

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type deleter interface {
	Delete(ctx context.Context, id uuid.UUID, isSoft bool) error
}

// @Tags section
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path      string  true  "id"
// @Param   is_soft     query     boolean    false  "по умолчанию soft"
// @Success 200
// @Router /v1/section/{id} [delete]
func New(deleter deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			id     uuid.UUID
			isSoft bool
			ctx    = r.Context()
		)

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, errors.Wrap(err, "parse id").Error())
			return
		}

		isSoft, err = strconv.ParseBool(r.URL.Query().Get("is_soft"))
		if err != nil {
			isSoft = true
		}

		err = deleter.Delete(ctx, id, isSoft)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
