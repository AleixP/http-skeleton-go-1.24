package middleware

import (
	"log"
	"net/http"
)

func MiddlewareUp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {

		if request.Header.Get("Authorization") == "" {
			log.Println("No Authorization header found")
		}
		next.ServeHTTP(responseWriter, request)
	})
}
