package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/decode"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, module *entity.Module, editorId uuid.UUID) (*entity.Module, error)
}

// @Tags module
// @Security ApiKeyAuth
// @Param       request body Module true "moduleRequest"
// @Success      201   {object} model.Module       "created"
// @Router       /v1/module [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Module
			ctx     = r.Context()
			userId  = r.Context().Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		module := request.toModule()
		module, err = creator.Create(ctx, module, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		dto := model.NewModule(module, map[uuid.UUID][]*entity.Answer{})
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, dto)
	}
}
