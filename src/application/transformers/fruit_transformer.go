package transformers

import (
	"http-skeleton-go-1.24/src/domain/model/fruit"
	"http-skeleton-go-1.24/src/user-interface/dto"
)

func transformFruit(fruit *fruit.Fruit, id int) dto.FruitResponse {
	return dto.FruitResponse{
		Id:    id,
		Name:  fruit.Name,
		Color: fruit.Color,
	}
}
