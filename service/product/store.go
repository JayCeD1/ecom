package product

import (
	"context"
	"ecom/types"
	"errors"
	"fmt"

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

func (s *Store) GetProducts(ctx context.Context) ([]*types.Product, error) {
	var products []*types.Product
	s.db.WithContext(ctx).Find(&products)
	return products, nil
}

func (s *Store) GetProductsByID(ctx context.Context, ids []int) ([]types.Product, error) {
	var products []types.Product
	s.db.WithContext(ctx).Where("id IN (?)", ids).Find(&products)
	return products, nil
}

func (s *Store) CreateProduct(ctx context.Context, product *types.Product) error {
	s.db.WithContext(ctx).Create(&product)
	return nil
}

func (s *Store) CheckProduct(ctx context.Context, name string) (bool, error) {
	result := s.db.WithContext(ctx).Where("name = ?", name).First(&types.Product{})
	//return result.RowsAffected > 0, result.Error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // product does NOT exist
		}
		return false, result.Error // some other DB error
	}
	return true, nil // product exists
}

func (s *Store) UpdateProductQuantity(ctx context.Context, item types.CartItem) error {
	res := s.db.WithContext(ctx).Model(&types.Product{}).Where("id = ? AND quantity >= ?", item.ProductID, item.Quantity).UpdateColumn("quantity", gorm.Expr("quantity - ?", item.Quantity))

	if res.RowsAffected == 0 {
		return fmt.Errorf("product %d not updated: insufficient stock", item.ProductID)
	}
	return res.Error
}
