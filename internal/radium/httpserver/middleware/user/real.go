package user

import (
	"github.com/google/uuid"
	"net/http"
)

func IsReal() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			id, ok := request.Context().Value("userId").(uuid.UUID)
			if !ok || id.String() != "11af02da-bf9e-4769-aa07-36903517733c" {
				next.ServeHTTP(writer, request)
				return
			}

			writer.WriteHeader(http.StatusForbidden)
			writer.Write([]byte("недостаточно прав"))
		})
	}
}
