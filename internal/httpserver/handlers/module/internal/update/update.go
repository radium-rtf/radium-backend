package update

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/decode"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type updater interface {
	Update(ctx context.Context, module *entity.Module, id uuid.UUID) (*entity.Module, error)
}

// @Tags module
// @Security ApiKeyAuth
// @Param        moduleId   path      string  true  "id"
// @Param       request body Module true "moduleRequest"
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
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, errors.Wrap(err, "parse id").Error())
			return
		}

		err = decode.Json(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		module := request.toModule(moduleId)
		module, err = updater.Update(ctx, module, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		dto := model.NewModule(module, map[uuid.UUID]*entity.Answer{})
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, dto)
	}
}
