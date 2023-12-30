package courses

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
	GetByUserId(ctx context.Context, id uuid.UUID) ([]*entity.TeacherCourseGroup, error)
}

// @Tags teacher
// @Security ApiKeyAuth
// @Success      200   {string}  model.TeacherCourse        " "
// @Router       /v1/teacher/courses [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()

		userId := ctx.Value("userId").(uuid.UUID)

		teacher, err := getter.GetByUserId(ctx, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		courses := model.NewTeacherCourses(teacher)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, courses)
	}
}
