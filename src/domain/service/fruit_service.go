package service

import "http-skeleton-go-1.24/src/domain/model/fruit"

type FruitRepository interface {
	Create(fruit *fruit.Fruit) (int64, error)
	List() ([]*fruit.Fruit, error)
	FindById(id string) (*fruit.Fruit, error)
}

type FruitService struct {
	fruitRepository FruitRepository
}

func NewFruitService(fruitRepository FruitRepository) *FruitService {
	return &FruitService{fruitRepository: fruitRepository}
}

func (s *FruitService) Create(fruit *fruit.Fruit) (int64, error) {
	return s.fruitRepository.Create(fruit)
}

func (s *FruitService) List() ([]*fruit.Fruit, error) {
	return s.fruitRepository.List()
}
func (s *FruitService) FindById(id string) (*fruit.Fruit, error) {
	return s.fruitRepository.FindById(id)
}
