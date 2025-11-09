package cart

import (
	"ecom/types"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.OrderStore
}

func NewHandler(store types.OrderStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	r.Post("/checkout", h.Checkout)
}

func (h *Handler) Checkout(c *fiber.Ctx) error {
	return nil
}
