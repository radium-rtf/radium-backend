package role

import (
	"github.com/radium-rtf/radium-backend/internal/radium/lib/auth"
	"net/http"
	"strings"
)

func CanEditCourse(manager auth.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")

			canEditCourse, err := manager.ExtractCanEditCourse(tokenHeader)
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte(err.Error()))
				return
			}

			if !canEditCourse {
				writer.WriteHeader(http.StatusForbidden)
				writer.Write([]byte("нет разрешения на редактирование курсов, попробуй перезайти (роль вшита в токен)"))
				return
			}

			next.ServeHTTP(writer, request)
		})
	}
}
