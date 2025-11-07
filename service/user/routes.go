package user

import "github.com/gofiber/fiber/v2"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	return nil
}
func (h *Handler) Register(c *fiber.Ctx) error {
	return nil
}
