package deletecoauthor

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

type deleter interface {
	DeleteCoAuthor(ctx context.Context, id, courseId, deleterId uuid.UUID) error
}

// @Tags role
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path      string  true  "id"
// @Param        courseId   path      string  true  "courseId"
// @Success 200
// @Router /v1/role/coauthor/{id}/{courseId} [delete]
func New(deleter deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = ctx.Value("userId").(uuid.UUID)
		)

		courseId, err := uuid.Parse(chi.URLParam(r, "courseId"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, fmt.Errorf("courseId: %w", err).Error())
			return
		}

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, fmt.Errorf("courseId: %w", err).Error())
			return
		}

		err = deleter.DeleteCoAuthor(ctx, id, courseId, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusOK)
	}
}
