package services

import (
	"log"

	"github.com/marcos-wz/capstone/models"
	"github.com/marcos-wz/capstone/store"
)

// CONTRACT **************************************

type FruitService interface {
	ReadFruit(id int) (*models.Fruit, error)
	LoadFruits()
	ListFruits()
}

type fruitService struct {
	store store.CvsStore
}

func NewFruitService(store store.CvsStore) FruitService {
	svc := &fruitService{store}
	return svc
}

// IMPLEMENTATION ********************************

func (fs *fruitService) ReadFruit(id int) (*models.Fruit, error) {
	fruit, err := fs.store.FindItem(id)
	if err != nil {
		return nil, err
	}
	log.Printf("Fruit Get: %+v", fruit)
	return fruit, nil
}

func (*fruitService) LoadFruits() {}

func (*fruitService) ListFruits() {}
