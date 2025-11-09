package types

import "time"

type UserStore interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
}
type ProductStore interface {
	GetProducts() ([]*Product, error)
	CreateProduct(product *Product) error
	CheckProduct(name string) (bool, error)
}
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
}
type UserRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=255"`
	LastName  string `json:"lastName" validate:"required,min=2,max=255"`
	Email     string `json:"email" validate:"required,email,max=255"`
	Password  string `json:"password" validate:"required,min=8,max=255"`
}

type ProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=255"`
	Description string  `json:"description" validate:"required,min=2,max=255"`
	Price       float64 `json:"price" validate:"required,min=0.01"`
	Quantity    int     `json:"quantity" validate:"required,min=1"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
