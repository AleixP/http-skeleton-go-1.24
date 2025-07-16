package middleware

import (
	"net/http"
)

func MiddlewareUp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Header.Get("Authorization") == "" {
			http.Error(responseWriter, "Forbidden action", http.StatusForbidden)
			return
		}
		next.ServeHTTP(responseWriter, request)
	})
}
