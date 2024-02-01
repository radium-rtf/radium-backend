package getbyid

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type getter interface {
	GetById(ctx context.Context, id uuid.UUID) (*entity2.Course, error)
	GetByIdAndUser(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*entity2.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Param        courseId   path     string  true  "course id"
// @Success      200   {object} model.Course  "ok"
// @Router       /v1/course/{courseId} [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()
		userId, ok := ctx.Value("userId").(uuid.UUID)

		courseId, err := uuid.Parse(chi.URLParam(r, "courseId"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		if !ok {
			responseIfNotAuthorized(getter, courseId, w, r)
			return
		}

		course, err := getter.GetByIdAndUser(ctx, courseId, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewCourseWithUserGroups(course, map[uuid.UUID][]*entity2.Answer{}, userId)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}

func responseIfNotAuthorized(getter getter, courseId uuid.UUID, w http.ResponseWriter, r *http.Request) {
	course, err := getter.GetById(r.Context(), courseId)
	if err != nil {
		resp.Error(r, w, err)
		return
	}

	c := model.NewCourse(course, map[uuid.UUID][]*entity2.Answer{}, uuid.UUID{})
	render.Status(r, http.StatusOK)
	render.JSON(w, r, c)
}
