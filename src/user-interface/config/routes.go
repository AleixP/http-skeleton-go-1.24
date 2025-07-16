package config

import (
	"http-skeleton-go-1.24/src/application/handler/create"
	"http-skeleton-go-1.24/src/application/handler/read"
	"http-skeleton-go-1.24/src/domain/service"
	"http-skeleton-go-1.24/src/infrastructure/middleware"
	"net/http"
	"strings"
)

func NewRouter(fruitService *service.FruitService) http.Handler {
	mux := http.NewServeMux()

	registerFruitRoutes(fruitService, mux)
	return mux
}

func registerFruitRoutes(fruitService *service.FruitService, mux *http.ServeMux) {
	mux.HandleFunc("/fruits", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler := create.NewCreateFruitHandler(fruitService)
			handler.Create(w, r)
		case http.MethodGet:
			handler := read.NewListFruitsQueryHandler(fruitService)
			handler.ListFruits(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.Handle("/fruits/", middleware.MiddlewareUp(findFruitByIdQueryHandler(fruitService)))
}

func findFruitByIdQueryHandler(fruitService *service.FruitService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/fruits/")
		if r.Method == http.MethodGet {
			handler := read.NewFindFruitQueryHandler(fruitService)
			handler.FindById(w, r, id)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
