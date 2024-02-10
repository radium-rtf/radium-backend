package role

import (
	"github.com/radium-rtf/radium-backend/internal/radium/lib/auth"
	"net/http"
	"strings"
)

func Author(manager auth.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")

			isAuthor, err := manager.ExtractIsAuthor(tokenHeader)
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte(err.Error()))
				return
			}

			if !isAuthor {
				writer.WriteHeader(http.StatusForbidden)
				writer.Write([]byte("нет роли автора, попробуй перезайти (роль вшита в токен)"))
				return
			}

			next.ServeHTTP(writer, request)
		})
	}
}
