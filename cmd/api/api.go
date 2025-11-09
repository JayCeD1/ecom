package api

import (
	"ecom/service/product"
	"ecom/service/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	addr string
	db   *gorm.DB
}

func NewServer(addr string, db *gorm.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}
func (s *Server) Run() error {
	app := fiber.New()

	// create a versioned API group
	apiV1 := app.Group("/api/v1")

	// ===== USER ROUTES =====
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userGroup := apiV1.Group("/users")
	userHandler.RegisterRoutes(userGroup)

	// ===== PRODUCT ROUTES =====
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productGroup := apiV1.Group("/products")
	productHandler.RegisterRoutes(productGroup)

	if err := app.Listen(s.addr); err != nil {
		return err
	}
	return nil
}
