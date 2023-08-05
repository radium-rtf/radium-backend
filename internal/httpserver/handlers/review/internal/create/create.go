package create

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type creator interface {
	CreateAnswerReview(ctx context.Context, review *entity.AnswerReview) (*entity.AnswerReview, error)
	CreateCodeReview(ctx context.Context, review *entity.CodeReview) (*entity.CodeReview, error)
}

// @Tags review
// @Security ApiKeyAuth
// @Accept json
// @Param request body Answer true "score - от 0 до 1"
// @Success 201 {object} entity.AnswerReview "created"
// @Router /v1/review/answer [post]
func NewAnswerReview(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Answer
			ctx     = r.Context()
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		review := request.toReview()
		review, err = creator.CreateAnswerReview(ctx, review)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, review)
	}
}

// @Tags review
// @Security ApiKeyAuth
// @Accept json
// @Param request body Code true "score - от 0 до 1"
// @Success 201 {object} entity.CodeReview "created"
// @Router /v1/review/code [post]
func NewCodeReview(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Code
			ctx     = r.Context()
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		review := request.toReview()
		review, err = creator.CreateCodeReview(ctx, review)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, review)
	}
}
