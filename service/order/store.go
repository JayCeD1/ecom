package order

import (
	"context"
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

func (s *Store) CreateOrder(ctx context.Context, order *types.Order) (int, error) {
	result := s.db.WithContext(ctx).Create(&order)

	if result.Error != nil {
		return 0, result.Error
	}
	return order.ID, nil
}

func (s *Store) CreateOrderItem(ctx context.Context, orderItem *types.OrderItem) error {
	result := s.db.WithContext(ctx).Create(&orderItem)
	return result.Error
}
