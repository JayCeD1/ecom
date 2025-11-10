package product

import (
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

func (s *Store) GetProducts() ([]*types.Product, error) {
	var products []*types.Product
	s.db.Find(&products)
	return products, nil
}

func (s *Store) GetProductsByID(ids []int) ([]types.Product, error) {
	var products []types.Product
	s.db.Where("id IN (?)", ids).Find(&products)
	return products, nil
}

func (s *Store) CreateProduct(product *types.Product) error {
	s.db.Create(&product)
	return nil
}

func (s *Store) CheckProduct(name string) (bool, error) {
	result := s.db.Where("name = ?", name).First(&types.Product{})
	//return result.RowsAffected > 0, result.Error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // product does NOT exist
		}
		return false, result.Error // some other DB error
	}
	return true, nil // product exists
}

func (s *Store) UpdateProductQuantity(item types.CartItem) error {
	res := s.db.Model(&types.Product{}).Where("id = ? AND quantity >= ?", item.ProductID, item.Quantity).UpdateColumn("quantity", gorm.Expr("quantity - ?", item.Quantity))

	if res.RowsAffected == 0 {
		return fmt.Errorf("product %d not updated: insufficient stock", item.ProductID)
	}
	return res.Error
}
