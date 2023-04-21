package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
	"strconv"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type courseRoutes struct {
	uc usecase.CourseUseCase
}

func newCourseRoutes(course chi.Router, h chi.Router, useCase usecase.CourseUseCase, signingKey string) {
	routes := courseRoutes{uc: useCase}

	course.Get("/", handler(routes.getCourses).HTTP)
	course.Get("/title/{courseId}", handler(routes.getCourseTitle).HTTP)

	course.Group(func(r chi.Router) {
		r.Use(authRequired(signingKey))
		r.Post("/", handler(routes.postCourse).HTTP)
	})

	h.Group(func(r chi.Router) {
		r.Use(authRequired(signingKey))
		r.Post("/link/course/{courseId}}", handler(routes.postLink).HTTP)
		r.Post("/collaborator/course/{courseId}", handler(routes.postCollaborator).HTTP)
		r.Post("/join/course/{courseId}", handler(routes.join).HTTP)
		r.Get("/my/course", handler(routes.myCourses).HTTP)
	})
}

// @Tags course
// @Security ApiKeyAuth
// @Accept       multipart/form-data
// @Param        logo formData  file           true  "logo"
// @Param        name formData  string         true  "name"
// @Param        description formData  string  true  "description"
// @Param        type formData  string        true  "kotlin, math и тд"
// @Success      201   {object} entity.Course       "created"
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
	userId := request.Context().Value("userId").(string)
	courseRequest := entity.NewCourseRequest(name[0], description[0], courseType[0], userId, logo, header)
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
// @Success      200   {object} entity.Course        "ok"
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

// @Tags course
// @Param        courseId   path     integer  true  "course id"
// @Success      200   {object} entity.CourseTitle  "ok"
// @Router       /course/title/{courseId} [get]
func (r courseRoutes) getCourseTitle(w http.ResponseWriter, request *http.Request) *appError {
	courseId := chi.URLParam(request, "courseId")
	id, err := strconv.Atoi(courseId)
	courses, err := r.uc.GetCourseTitle(request.Context(), id)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, courses)
	return nil
}

// @Tags course
// @Security ApiKeyAuth
// @Param       request body entity.Link true "link"
// @Param        courseId   path     integer  true  "course id"
// @Success      201   {object} entity.Link "ok"
// @Router       /link/course/{courseId}} [post]
func (r courseRoutes) postLink(w http.ResponseWriter, request *http.Request) *appError {
	var link entity.Link
	err := render.DecodeJSON(request.Body, &link)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	courseId := chi.URLParam(request, "courseId")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	link, err = r.uc.CreateLink(request.Context(), id, link)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, link)
	return nil
}

// @Tags course
// @Security ApiKeyAuth
// @Param       request body entity.Collaborator true "collaborator"
// @Param        courseId   path      integer  true  "course id"
// @Success      201   {object} entity.Link "ok"
// @Router       /collaborator/course/{courseId} [post]
func (r courseRoutes) postCollaborator(w http.ResponseWriter, request *http.Request) *appError {
	var collaborator entity.Collaborator
	err := render.DecodeJSON(request.Body, &collaborator)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	courseId := chi.URLParam(request, "courseId")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	collaborator, err = r.uc.CreateCollaborator(request.Context(), id, collaborator)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, collaborator)
	return nil
}

// @Tags course
// @Security ApiKeyAuth
// @Success      200   {object} entity.Course "ok"
// @Router       /my/course [get]
func (r courseRoutes) myCourses(w http.ResponseWriter, request *http.Request) *appError {
	userId := request.Context().Value("userId").(string)
	courses, err := r.uc.GetStudentCourses(request.Context(), userId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusOK)
	render.JSON(w, request, courses)
	return nil
}

// @Tags course
// @Security ApiKeyAuth
// @Param        courseId   path      integer  true  "course id"
// @Success      201   {object} entity.Course "created"
// @Router       /join/course/{courseId} [post]
func (r courseRoutes) join(w http.ResponseWriter, request *http.Request) *appError {
	userId := request.Context().Value("userId").(string)
	courseId := chi.URLParam(request, "courseId")
	courses, err := r.uc.Join(request.Context(), userId, courseId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, courses)
	return nil
}
