package addstudent

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type adder interface {
	AddStudent(ctx context.Context, groupId uuid.UUID, userId uuid.UUID) (*entity.Group, error)
}

// @Tags group
// @Security ApiKeyAuth
// @Param       request body AddStudent true " "
// @Success      201   {string} model.Group "patched"
// @Router       /v1/group/{groupId}/students [patch]
func New(adder adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx     = r.Context()
			request AddStudent
		)

		groupId, err := uuid.Parse(chi.URLParam(r, "groupId"))

		if err != nil {
			resp.Error(r, w, err)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		group, err := adder.AddStudent(ctx, groupId, request.UserId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewGroup(group)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, dto)
	}
}
