package config

import (
	"http-skeleton-go-1.24/src/application/handler/create"
	"http-skeleton-go-1.24/src/application/query/read"
	"http-skeleton-go-1.24/src/domain/service"
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
			handler := read.NewListFruitsQuery(fruitService)
			handler.ListFruits(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/fruits/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/fruits/")
		switch r.Method {
		case http.MethodGet:
			handler := read.NewFindFruitQuery(fruitService)
			handler.FindById(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
