package update

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/decode"
	"github.com/radium-rtf/radium-backend/internal/model"
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
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		err = decode.Json(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		course := request.toCourse(courseId)
		course, err = updater.Update(ctx, course, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		c := model.NewCourse(course, map[uuid.UUID]*entity.Answer{})
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, c)
	}
}
