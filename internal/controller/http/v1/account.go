package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type accountRoutes struct {
	uc usecase.AccountUseCase
}

func newAccountRoutes(h chi.Router, useCase usecase.AccountUseCase, signingKey string) {
	routes := accountRoutes{uc: useCase}
	h.Route("/account", func(r chi.Router) {
		r.Use(authRequired(signingKey))
		r.Get("/", handler(routes.account).HTTP)
		// r.Get("/course", handler(routes.course).HTTP)
		r.Patch("/", handler(routes.update).HTTP)
		r.Patch("/password", handler(routes.password).HTTP)
	})
}

// @Tags  	    account
// @Accept      json
// @Produce     json
// @Security ApiKeyAuth
// @Success     200 {object} entity.UserDto
// @Router      /account [get]
func (r accountRoutes) account(w http.ResponseWriter, request *http.Request) *appError {
	userId := request.Context().Value("userId")
	print(userId.(string))
	user, err := r.uc.Account(request.Context(), userId.(string))
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, user)
	return nil
}

// @Tags  	    account
// @Accept      json
// @Produce     json
// // @Security ApiKeyAuth
// @Param       request body entity.UpdateUserRequest true "Данные для обновления"
// @Success     200
// @Router      /account [patch]
func (r accountRoutes) update(w http.ResponseWriter, request *http.Request) *appError {
	var update entity.UpdateUserRequest
	err := render.DecodeJSON(request.Body, &update)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	userId := request.Context().Value("userId").(string)
	result, err := r.uc.UpdateUser(request.Context(), userId, update)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, result)

	return nil
}

// @Tags  	    account
// @Accept      json
// @Produce     json
// @Security ApiKeyAuth
// @Param       request body entity.PasswordUpdate true "PasswordUpdate"
// @Success     200
// @Router      /account/password [patch]
func (r accountRoutes) password(w http.ResponseWriter, request *http.Request) *appError {
	var passwordUpdate entity.PasswordUpdate
	err := render.DecodeJSON(request.Body, &passwordUpdate)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	userId := request.Context().Value("userId").(string)
	err = r.uc.UpdatePassword(request.Context(), userId, passwordUpdate)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

// // @Tags account
// // @Security ApiKeyAuth
// // @Success      200   {object} entity.Course "ok"
// // @Router       /account/course [get]
// func (r accountRoutes) course(w http.ResponseWriter, request *http.Request) *appError {
// 	userId := request.Context().Value("userId").(string)
// 	courses, err := r.uc.GetStudentCourses(request.Context(), userId)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}
// 	render.Status(request, http.StatusOK)
// 	render.JSON(w, request, courses)
// 	return nil
// }
