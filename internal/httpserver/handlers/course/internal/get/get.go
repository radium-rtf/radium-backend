package get

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetCourses(ctx context.Context) ([]*entity.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Success      200   {object} model.Course        "ok"
// @Router       /v1/course [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId, _ := ctx.Value("userId").(uuid.UUID)

		courses, err := getter.GetCourses(ctx)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewCourses(courses, userId)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
