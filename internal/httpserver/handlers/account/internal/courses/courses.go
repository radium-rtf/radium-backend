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
	GetStudentCourses(ctx context.Context, id uuid.UUID) ([]*entity.Course, error)
	GetRecommendations(ctx context.Context, userId uuid.UUID, limit int) ([]*entity.Course, error)
}

// @Tags account
// @Security ApiKeyAuth
// @Success      200   {object} Courses "ok"
// @Router       /v1/account/courses [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = ctx.Value("userId").(uuid.UUID)
		)

		courses, err := getter.GetStudentCourses(ctx, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		recommendations, err := getter.GetRecommendations(ctx, userId, 10)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		response := Courses{
			My:              model.NewCourses(courses, userId),
			Recommendations: model.NewCourses(recommendations, userId),
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, response)
	}
}
