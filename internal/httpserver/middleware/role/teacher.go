package role

import (
	"github.com/radium-rtf/radium-backend/internal/lib/auth"
	"net/http"
	"strings"
)

func Teacher(manager auth.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")

			isTeacher, err := manager.ExtractIsTeacher(tokenHeader)
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte(err.Error()))
				return
			}

			if !isTeacher {
				writer.WriteHeader(http.StatusForbidden)
				writer.Write([]byte("нет роли преподавателя, попробуй перезайти (роль вшита в токен)"))
				return
			}

			next.ServeHTTP(writer, request)
		})
	}
}
