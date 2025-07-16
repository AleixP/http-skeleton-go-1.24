package read

import (
	"encoding/json"
	"http-skeleton-go-1.24/src/application/transformers"
	"http-skeleton-go-1.24/src/domain/service"
	"net/http"
	"strconv"
)

type FindFruitQueryHandler struct {
	FruitService *service.FruitService
}

func NewFindFruitQueryHandler(fruitService *service.FruitService) *FindFruitQueryHandler {
	return &FindFruitQueryHandler{FruitService: fruitService}
}
func (findFruitQueryHandler *FindFruitQueryHandler) FindById(w http.ResponseWriter, r *http.Request, id string) {
	fruit, err := findFruitQueryHandler.FruitService.FindById(id)
	if err != nil {
		http.Error(w, "Fruit not found", http.StatusNotFound)
		return
	}
	intId, _ := strconv.Atoi(id)
	response := transformers.TransformFruit(fruit, intId)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
