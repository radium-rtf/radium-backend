package join

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type connector interface {
	Join(ctx context.Context, userId, courseId uuid.UUID) (*entity.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Param        courseId   path      string  true  "course id"
// @Success      201   {object} model.Course "created"
// @Router       /v1/course/join/{courseId} [patch]
func New(connector connector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = ctx.Value("userId").(uuid.UUID)
		)

		courseId, err := uuid.Parse(chi.URLParam(r, "courseId"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		course, err := connector.Join(ctx, userId, courseId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, model.NewCourse(course, nil, userId))
	}
}
