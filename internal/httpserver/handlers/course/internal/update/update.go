package update

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"net/http"
)

type updater interface {
	Update(ctx context.Context, course *entity.Course, userId uuid.UUID) (*entity.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Accept       json
// @Param        courseId   path     string  true  "course id"
// @Param request body Course true "Данные о курсе"
// @Success      201   {object} model.Course       "created"
// @Router       /v1/course/{courseId} [put]
func New(updater updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Course
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		courseId, err := uuid.Parse(chi.URLParam(r, "courseId"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		err = decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		course := request.toCourse(courseId)
		course, err = updater.Update(ctx, course, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewCourse(course, map[uuid.UUID][]*entity.Answer{}, userId)
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, c)
	}
}
