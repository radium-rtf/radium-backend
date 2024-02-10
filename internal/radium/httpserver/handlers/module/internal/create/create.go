package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, module *entity2.Module, editorId uuid.UUID) (*entity2.Module, error)
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
			resp.Error(r, w, err)
			return
		}

		module := request.toModule()
		module, err = creator.Create(ctx, module, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		dto := model.NewModule(module, map[uuid.UUID][]*entity2.Answer{})
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, dto)
	}
}
