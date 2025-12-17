package types

import (
	"context"
	"time"
)

type UserStore interface {
	CreateUser(user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(id int) (*User, error)
}
type ProductStore interface {
	GetProducts() ([]*Product, error)
	GetProductsByID(ids []int) ([]Product, error)
	CreateProduct(product *Product) error
	CheckProduct(name string) (bool, error)
	UpdateProductQuantity(cartItem CartItem) error
}

type OrderStore interface {
	CreateOrder(order *Order) (int, error)
	CreateOrderItem(orderItem *OrderItem) error
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
	// Optional: reverse relationship
	OrderItem []OrderItem `json:"-" gorm:"foreignKey:ProductID"`
}

type Order struct {
	ID        int         `json:"id"`
	UserID    int         `json:"userID"`
	User      User        `json:"user" gorm:"foreignKey:UserID"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	Address   string      `json:"address"`
	CreatedAt time.Time   `json:"createdAt"`
	OrderItem []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderID"`
	Order     Order     `json:"order" gorm:"foreignKey:OrderID"`
	ProductID int       `json:"productID"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
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

type CartItem struct {
	ProductID int `json:"productID" validate:"required"`
	Quantity  int `json:"quantity" validate:"required,min=1"`
}
type CartCheckoutRequest struct {
	Items []CartItem `json:"items" validate:"required,min=1,dive"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
