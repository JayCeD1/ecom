package api

import (
	"database/sql"
	"ecom/service/user"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}
func (s *Server) Run() error {
	app := fiber.New()

	// create a versioned API group
	apiV1 := app.Group("/api/v1")

	// mount user routes under /api/v1/users
	userGroup := apiV1.Group("/users")
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(userGroup)

	if err := app.Listen(s.addr); err != nil {
		return err
	}
	return nil
}
