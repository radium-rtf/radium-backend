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
	Update(ctx context.Context, module *entity.Module, id uuid.UUID) (*entity.Module, error)
}

// @Tags module
// @Security ApiKeyAuth
// @Param        moduleId   path      string  true  "id"
// @Param       request body Module true " "
// @Success      201   {object} model.Module       "created"
// @Router       /v1/module/{moduleId} [put]
func New(updater updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Module
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		moduleId, err := uuid.Parse(chi.URLParam(r, "moduleId"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		err = decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		module := request.toModule(moduleId)
		module, err = updater.Update(ctx, module, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewModule(module, map[uuid.UUID][]*entity.Answer{})
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, dto)
	}
}
