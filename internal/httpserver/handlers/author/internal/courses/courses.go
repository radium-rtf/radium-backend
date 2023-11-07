package courses

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetCoursesByAuthorId(ctx context.Context, id uuid.UUID) ([]*entity.Course, error)
}

// @Tags author
// @Security ApiKeyAuth
// @Success      200   {object} model.Course "ok"
// @Router       /v1/author/courses [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = ctx.Value("userId").(uuid.UUID)
		)

		courses, err := getter.GetCoursesByAuthorId(ctx, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		c := model.NewCourses(courses, userId)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
