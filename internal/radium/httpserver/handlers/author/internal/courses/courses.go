package courses

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type getter interface {
	GetCoursesByAuthorOrCoauthorId(ctx context.Context, id uuid.UUID) ([]*entity.Course, error)
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

		courses, err := getter.GetCoursesByAuthorOrCoauthorId(ctx, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewCourses(courses, userId)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
