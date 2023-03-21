package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type courseRoutes struct {
	uc usecase.CourseUseCase
}

func newCourseRoutes(h chi.Router, useCase usecase.CourseUseCase) {
	routes := courseRoutes{uc: useCase}
	h.Route("/course", func(r chi.Router) {
		r.Get("/", handler(routes.getCourses).HTTP)
		r.Post("/", handler(routes.postCourse).HTTP)
	})
}

// @Tags course
// @Security ApiKeyAuth
// @Accept       multipart/form-data
// @Param        logo formData  file           true  "logo"
// @Param        chat formData  string           true  "chat"
// @Param        name formData  string         true  "name"
// @Param        description formData  string  true  "description"
// @Param        type formData  string        true  "kotlin, math и тд"
// @Success      201   {string}  string        "created"
// @Router       /course [post]
func (r courseRoutes) postCourse(w http.ResponseWriter, request *http.Request) *appError {
	err := request.ParseMultipartForm(8 * 1024 * 1024 * 8)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	form := request.MultipartForm
	name, ok := form.Value["name"]
	if !ok {
		return newAppError(errors.New("the course must have \"name\" "), http.StatusBadRequest)
	}
	chat, ok := form.Value["chat"]
	if !ok {
		return newAppError(errors.New("the course must have \"chat\" "), http.StatusBadRequest)
	}
	description, ok := form.Value["description"]
	if !ok {
		return newAppError(errors.New("the course must have \"description\" "), http.StatusBadRequest)
	}
	courseType, ok := form.Value["type"]
	if !ok {
		return newAppError(errors.New("the course must have \"type\" "), http.StatusBadRequest)
	}
	logo, header, err := request.FormFile("logo")
	if !ok {
		return newAppErrorf("%v :the course must have \"logo\" ", err, http.StatusBadRequest)
	}

	courseRequest := entity.NewCourseRequest(name[0], description[0], courseType[0], chat[0], logo, header)
	course, err := r.uc.CreateCourse(request.Context(), courseRequest)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, course)
	return nil
}

// @Tags course
// @Security ApiKeyAuth
// @Success      200   {string}  string        "ok"
// @Router       /course [get]
func (r courseRoutes) getCourses(w http.ResponseWriter, request *http.Request) *appError {
	courses, err := r.uc.GetCourses(request.Context())
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, courses)
	return nil
}
