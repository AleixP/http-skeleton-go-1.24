package create

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"http-skeleton-go-1.24/src/application/transformers"
	fruitEntity "http-skeleton-go-1.24/src/domain/model/fruit"
	"http-skeleton-go-1.24/src/domain/service"
	"http-skeleton-go-1.24/src/infrastructure/dto"
	"net/http"
)

type CreateFruitHandler struct {
	FruitService *service.FruitService
}

func NewCreateFruitHandler(fruitService *service.FruitService) *CreateFruitHandler {
	return &CreateFruitHandler{FruitService: fruitService}
}

func (handler *CreateFruitHandler) Create(responseWriter http.ResponseWriter, request *http.Request) {
	var fruitDto dto.CrateFruitDtoRequest
	
	if err := json.NewDecoder(request.Body).Decode(&fruitDto); err != nil {
		http.Error(responseWriter, "Invalid request", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(fruitDto); err != nil {
		http.Error(responseWriter, "Missing required fields", http.StatusBadRequest)
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
