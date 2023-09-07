package report

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetReportByCourse(ctx context.Context, userId, courseId, groupId uuid.UUID) (*model.Report, error)
}

// @Tags group
// @Security ApiKeyAuth
// @Param        courseId   path      string  true  "courseId"
// @Param        groupId   path      string  true  "groupId"
// @Success      200   {object}  model.Report        " "
// @Router       /v1/group/report/{groupId}/{courseId} [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = r.Context().Value("userId").(uuid.UUID)
		)

		courseId, err := uuid.Parse(chi.URLParam(r, "courseId"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		groupId, err := uuid.Parse(chi.URLParam(r, "groupId"))
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		report, err := getter.GetReportByCourse(ctx, userId, courseId, groupId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, report)
	}
}
