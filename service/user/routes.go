package user

import (
	"ecom/types"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	return nil
}
func (h *Handler) Register(c *fiber.Ctx) error {

	var payload types.UserRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return nil
}
