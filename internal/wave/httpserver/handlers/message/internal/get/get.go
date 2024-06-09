package get

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type getter interface {
	GetMessagesFrom(
		ctx context.Context, chatId uuid.UUID, page, pageSize int, sort, order string,
	) ([]*entity.Message, error)
}

// @Tags message
// @Security ApiKeyAuth
// @Param        chatId   path      string  true  "ID группы/диалога"
// @Param        page     query     int     false "Номер страницы"  default(1)
// @Param        pageSize query     int     false "Размер страницы" default(50)
// @Param        sort     query     string  false "Тип сортировки" default(date)
// @Param        order    query     string  false "Порядок сортировки (asc или desc)" default(desc)
// @Success      200   {object} []model.Message        " "
// @Router       /v1/messages/{chatId} [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
			// userId = r.Context().Value("userId").(uuid.UUID)
		)

		chatId, err := uuid.Parse(chi.URLParam(r, "chatId"))
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			page = 1
		}

		pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
		if err != nil || pageSize < 1 {
			pageSize = 50
		}

		sort := r.URL.Query().Get("sort")
		if sort == "" {
			sort = "date"
		}

		order := r.URL.Query().Get("order")
		if order == "" {
			order = "desc"
		}

		messageObjects, err := getter.GetMessagesFrom(ctx, chatId, page, pageSize, sort, order)
		messages := model.NewMessages(messageObjects)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, messages)
	}
}
