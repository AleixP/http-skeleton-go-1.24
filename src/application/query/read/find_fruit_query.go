package read

import (
	"encoding/json"
	"http-skeleton-go-1.24/src/application/transformers"
	"http-skeleton-go-1.24/src/domain/services"
	"net/http"
	"strconv"
)

type FindFruitQuery struct {
	FruitRepository services.FruitRepository
}

func (findFruitQuery *FindFruitQuery) FindById(w http.ResponseWriter, r *http.Request, id string) {
	fruit, err := findFruitQuery.FruitRepository.FindById(id)
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
