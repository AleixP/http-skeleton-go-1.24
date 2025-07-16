package create

import (
	"encoding/json"
	"http-skeleton-go-1.24/src/application/transformers"
	fruitEntity "http-skeleton-go-1.24/src/domain/model/fruit"
	"http-skeleton-go-1.24/src/domain/services"
	"http-skeleton-go-1.24/src/infrastructure/dto"
	"net/http"
)

type CreateFruitHandler struct {
	FruitService *services.FruitService
}

func NewCreateFruitHandler(fruitService *services.FruitService) *CreateFruitHandler {
	return &CreateFruitHandler{FruitService: fruitService}
}

func (handler *CreateFruitHandler) Create(responseWriter http.ResponseWriter, request *http.Request) {
	var fruitDto dto.CrateFruitDtoRequest
	if err := json.NewDecoder(request.Body).Decode(&fruitDto); err != nil {
		http.Error(responseWriter, "Invalid request", http.StatusBadRequest)
		return
	}
	fruit := fruitEntity.CreateFromPrimitive(fruitDto.Name, fruitDto.Color)
	id, err := handler.FruitService.Create(fruit)
	if err != nil {
		http.Error(responseWriter, "Error saving fruit: ", http.StatusInternalServerError)
		return
	}

	response := transformers.TransformFruit(fruit, int(id))
	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(response)
}
