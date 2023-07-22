package create

import (
	"context"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, review *entity.AnswerReview) (*entity.AnswerReview, error)
}

// @Tags review
// @Security ApiKeyAuth
// @Accept json
// @Param request body Request true "score - от 0 до 1"
// @Success 201 {object} entity.AnswerReview "created"
// @Router /v1/review/answer [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Request
			ctx     = r.Context()
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		review := request.ToReview()
		review, err = creator.Create(ctx, review)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, review)
	}
}
