//go:build ignore
// +build ignore

package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type groupRoutes struct {
	uc usecase.GroupUseCase
}

func newGroupRoutes(h chi.Router, useCase usecase.GroupUseCase, signingKey string) {
	routes := groupRoutes{uc: useCase}

	h.Group(func(r chi.Router) {
		r.Use(authRequired(signingKey))
		r.Post("/teacher/{userId}/group/{groupId}", handler(routes.teacher).HTTP)
		r.Patch("/join/group/{groupId}", handler(routes.join).HTTP)
	})

	h.Route("/group", func(r chi.Router) {
		r.Post("/", handler(routes.create).HTTP)
	})
}

// @Tags group
// // @Security ApiKeyAuth
// // @Success      201   {string}  string        "created"
// // @Param       request body entity.GroupName true "GroupName"
// // @Router       /group [post]
func (r groupRoutes) create(w http.ResponseWriter, request *http.Request) *appError {
	var groupName entity.GroupName
	err := render.DecodeJSON(request.Body, &groupName)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	group, err := r.uc.Create(request.Context(), groupName)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, group)
	return nil
}

// @Tags group
// @Security ApiKeyAuth
// @Param        groupId   path      string  true  "group id"
// @Success      200   {string}  string        "created"
// @Router       /join/group/{groupId} [patch]
func (r groupRoutes) join(w http.ResponseWriter, request *http.Request) *appError {
	groupId := chi.URLParam(request, "groupId")
	userId := request.Context().Value("userId").(uuid.UUID)
	var joinGroup = entity.GroupJoin{GroupId: groupId, UserId: userId}
	err := r.uc.Join(request.Context(), joinGroup)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

// @Tags group
// @Security ApiKeyAuth
// @Param        userId   path      string  true  "user id"
// @Param        groupId   path      string  true  "group id"
// @Success      201   {string}  string        "created"
// @Router       /teacher/{userId}/group/{groupId} [post]
func (r groupRoutes) teacher(w http.ResponseWriter, request *http.Request) *appError {
	groupId := chi.URLParam(request, "groupId")
	teacherId := request.Context().Value("userId").(uuid.UUID)
	var groupTeacher = entity.GroupTeacher{GroupId: groupId, UserId: teacherId, Id: uuid.NewString()}
	err := r.uc.CreateTeacher(request.Context(), groupTeacher)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	return nil
}
