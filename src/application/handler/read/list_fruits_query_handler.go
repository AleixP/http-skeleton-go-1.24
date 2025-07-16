package read

import (
	"encoding/json"
	"http-skeleton-go-1.24/src/application/transformers"
	"http-skeleton-go-1.24/src/domain/service"
	"http-skeleton-go-1.24/src/user-interface/dto"
	"net/http"
)

type ListFruitsQueryHandler struct {
	FruitService *service.FruitService
}

func NewListFruitsQueryHandler(fruitService *service.FruitService) *ListFruitsQueryHandler {
	return &ListFruitsQueryHandler{FruitService: fruitService}
}

func (listFruitsQueryHandler *ListFruitsQueryHandler) ListFruits(w http.ResponseWriter, r *http.Request) {
	items, err := listFruitsQueryHandler.FruitService.List()
	if err != nil {
		http.Error(w, "Error listing records", http.StatusInternalServerError)
		return
	}

	var transformedFruits []dto.FruitResponse
	for _, fruit := range items {
		transformedFruits = append(transformedFruits, transformers.TransformFruit(fruit, *fruit.ID))
	}

	json.NewEncoder(w).Encode(transformedFruits)
}
