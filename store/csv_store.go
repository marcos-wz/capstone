package store

import "github.com/marcos-wz/capstone/models"

// CONTRACT **************************************

type CvsStore interface {
	FindItem(id int) (*models.Fruit, error)
	InsertItems() error
	ListItems() (*models.Fruits, error)
}

type cvsStore struct{}

func NewCvsStore() CvsStore {
	return &cvsStore{}
}

// IMPLEMENTATION *******************************

func (*cvsStore) FindItem(id int) (*models.Fruit, error) {
	return &models.Fruit{}, nil
}

func (*cvsStore) InsertItems() error {
	return nil
}

func (*cvsStore) ListItems() (*models.Fruits, error) {
	return nil, nil
}
