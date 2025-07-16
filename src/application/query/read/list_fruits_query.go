package read

import (
	"encoding/json"
	"http-skeleton-go-1.24/src/application/transformers"
	"http-skeleton-go-1.24/src/domain/service"
	"http-skeleton-go-1.24/src/user-interface/dto"
	"net/http"
)

type ListFruitsQuery struct {
	FruitService *service.FruitService
}

func NewListFruitsQuery(fruitService *service.FruitService) *ListFruitsQuery {
	return &ListFruitsQuery{FruitService: fruitService}
}

func (listFruitsQuery *ListFruitsQuery) ListFruits(w http.ResponseWriter, r *http.Request) {
	items, err := listFruitsQuery.FruitService.List()
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
