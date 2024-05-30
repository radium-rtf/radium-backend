package role

import (
	"net/http"
	"strings"

	"github.com/radium-rtf/radium-backend/internal/radium/lib/auth"
)

func Admin(manager auth.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tokenHeader := strings.Split(request.Header.Get("Authorization"), " ")

			isAdmin, err := manager.ExtractIsAdmin(tokenHeader)
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte(err.Error()))
				return
			}

			if !isAdmin {
				writer.WriteHeader(http.StatusForbidden)
				writer.Write([]byte("Доступно только администраторам"))
				return
			}

			next.ServeHTTP(writer, request)
		})
	}
}
