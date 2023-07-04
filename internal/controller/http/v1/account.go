package v1

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type accountRoutes struct {
	uc usecase.AccountUseCase
}

func newAccountRoutes(useCase usecase.AccountUseCase, signingKey string) *chi.Mux {
	r := chi.NewRouter()
	routes := accountRoutes{uc: useCase}

	r.Use(authRequired(signingKey))
	r.Get("/", handler(routes.account).HTTP)
	r.Get("/courses", handler(routes.getStudentCourses).HTTP)
	r.Patch("/", handler(routes.update).HTTP)
	r.Patch("/password", handler(routes.password).HTTP)

	return r
}

// @Tags  	    account
// @Accept      json
// @Produce     json
// @Security ApiKeyAuth
// @Success     200 {object} entity.UserDto
// @Router      /account [get]
func (r accountRoutes) account(w http.ResponseWriter, request *http.Request) *appError {
	userId := request.Context().Value("userId").(uuid.UUID)
	user, err := r.uc.Account(request.Context(), userId)
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
// @Security ApiKeyAuth
// @Param       request body entity.UpdateUserRequest true "Данные для обновления"
// @Success     200
// @Router      /account [patch]
func (r accountRoutes) update(w http.ResponseWriter, request *http.Request) *appError {
	var update entity.UpdateUserRequest
	err := render.DecodeJSON(request.Body, &update)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	userId := request.Context().Value("userId").(uuid.UUID)
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
	userId := request.Context().Value("userId").(uuid.UUID)
	err = r.uc.UpdatePassword(request.Context(), userId, passwordUpdate)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

// @Tags account
// @Security ApiKeyAuth
// @Success      200   {object} entity.CourseDto "ok"
// @Router       /account/courses [get]
func (r accountRoutes) getStudentCourses(w http.ResponseWriter, request *http.Request) *appError {
	userId := request.Context().Value("userId").(uuid.UUID)
	courses, err := r.uc.GetStudentCourses(request.Context(), userId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	c := make([]entity.CourseDto, 0, len(courses))
	dto.Map(&c, courses)
	render.Status(request, http.StatusOK)
	render.JSON(w, request, c)
	return nil
}
