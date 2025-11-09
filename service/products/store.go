package products

import (
	"ecom/types"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	var products []*types.Product
	s.db.Find(&products)
	return products, nil
}
