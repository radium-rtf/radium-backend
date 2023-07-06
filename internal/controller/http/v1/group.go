package v1

import (
	"github.com/radium-rtf/radium-backend/pkg/mapper"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type groupRoutes struct {
	uc usecase.GroupUseCase
	mapper.Group
}

func newGroupRoutes(useCase usecase.GroupUseCase, signingKey string) chi.Router {
	routes := groupRoutes{uc: useCase}

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/{id}", handler(routes.getById).HTTP)
		r.Get("/", handler(routes.get).HTTP)
	})

	r.Group(func(r chi.Router) {
		r.Use(authRequired(signingKey))
		r.Post("/", handler(routes.post).HTTP)
		r.Patch("/invite/{inviteCode}", handler(routes.invite).HTTP)
	})

	return r
}

// @Tags group
// @Security ApiKeyAuth
// @Success      201   {string}  entity.GroupDto        "created"
// @Param       request body entity.GroupPost true " "
// @Router       /group [post]
func (r groupRoutes) post(w http.ResponseWriter, request *http.Request) *appError {
	var post entity.GroupPost
	err := render.DecodeJSON(request.Body, &post)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	group := r.PostToGroup(post)
	group, err = r.uc.Create(request.Context(), group)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	dto := r.ToDto(group)
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, dto)
	return nil
}

// @Tags group
// @Security ApiKeyAuth
// @Param        id   path      string  true  "groupId"
// @Success      200   {string}  entity.GroupDto         " "
// @Router       /group/{id} [get]
func (r groupRoutes) getById(w http.ResponseWriter, request *http.Request) *appError {
	id, err := uuid.Parse(chi.URLParam(request, "id"))
	group, err := r.uc.GetById(request.Context(), id)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	dto := r.ToDto(group)
	render.Status(request, http.StatusOK)
	render.JSON(w, request, dto)
	return nil
}

// @Tags group
// @Security ApiKeyAuth
// @Success      200   {string}  entity.GroupDto         " "
// @Router       /group [get]
func (r groupRoutes) get(w http.ResponseWriter, request *http.Request) *appError {
	groups, err := r.uc.Get(request.Context())
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	dto := r.ToGroupsDto(groups)
	render.Status(request, http.StatusOK)
	render.JSON(w, request, dto)
	return nil
}

// @Tags group
// @Security ApiKeyAuth
// @Param        inviteCode   path      string  true  "inviteCode"
// @Success      200   {string}  string        " "
// @Router       /group/invite/{inviteCode} [patch]
func (r groupRoutes) invite(w http.ResponseWriter, request *http.Request) *appError {
	inviteCode := chi.URLParam(request, "inviteCode")
	userId := request.Context().Value("userId").(uuid.UUID)
	var joinGroup = entity.GroupJoin{InviteCode: inviteCode, UserId: userId}
	err := r.uc.Join(request.Context(), joinGroup)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}
