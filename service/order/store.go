package order

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

func (s *Store) CreateOrder(order *types.Order) (int, error) {
	result := s.db.Create(&order)

	if result.Error != nil {
		return 0, result.Error
	}
	return order.ID, nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	result := s.db.Create(&orderItem)
	return result.Error
}
