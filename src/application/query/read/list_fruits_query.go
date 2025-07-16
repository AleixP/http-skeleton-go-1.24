package read

import (
	"encoding/json"
	"http-skeleton-go-1.24/src/domain/services"
	"net/http"
)

type ListFruitsQuery struct {
	FruitRepository services.FruitRepository
}

func (listFruitsQuery *ListFruitsQuery) ListFruits(w http.ResponseWriter, r *http.Request) {
	items, err := listFruitsQuery.FruitRepository.List()
	if err != nil {
		http.Error(w, "Error listing records", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}
