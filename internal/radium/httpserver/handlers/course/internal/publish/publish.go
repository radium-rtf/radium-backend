package publish

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

type publisher interface {
	Publish(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*entity2.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Accept       json
// @Param        id   path     string  true  "опубликовать или снять с публикации"
// @Success      200   {object} model.Course       " "
// @Router       /v1/course/publish/{id} [patch]
func New(publisher publisher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = ctx.Value("userId").(uuid.UUID)
		)

		courseId, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		course, err := publisher.Publish(ctx, courseId, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewCourse(course, map[uuid.UUID][]*entity2.Answer{}, userId)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
