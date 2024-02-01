package getbygroup

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"net/http"
)

type getter interface {
	GetWithAnswers(ctx context.Context, teacherId, groupId, courseId uuid.UUID) (*entity.Group, error)
}

// @Tags answer
// @Security ApiKeyAuth
// @Accept json
// @Success 200 {object} model.GroupAnswers "ok"
// @Param        groupId   path      string  true  "id"
// @Param   course_id     query     string    true  " "
// @Router /v1/answers/group/{groupId} [get]
func NewAnswer(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx       = r.Context()
			teacherId = ctx.Value("userId").(uuid.UUID)
		)

		groupId, err := uuid.Parse(chi.URLParam(r, "groupId"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, fmt.Errorf("groupId: %w", err).Error())
			return
		}

		courseId, err := uuid.Parse(r.URL.Query().Get("course_id"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, fmt.Errorf("courseId: %w", err).Error())
			return
		}

		group, err := getter.GetWithAnswers(ctx, teacherId, groupId, courseId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, model.NewGroupAnswers(group))
	}
}
