package getbyid

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetById(ctx context.Context, id uuid.UUID) (*entity.Course, error)
}

// @Tags course
// @Param        courseId   path     string  true  "course id"
// @Success      200   {object} model.Course  "ok"
// @Router       /v1/course/{courseId} [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()

		courseId, err := uuid.Parse(chi.URLParam(r, "courseId"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		course, err := getter.GetById(ctx, courseId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		c := model.NewCourse(course)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
