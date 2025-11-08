package user

import (
	"ecom/types"
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

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	s.db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	var user types.User
	s.db.Where("id = ?", id).First(&user)

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (s *Store) CreateUser(user *types.User) error {
	s.db.Create(&user)
	return nil
}
