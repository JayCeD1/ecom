package products

import (
	"ecom/types"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	r.Post("/", h.GetProducts)
}

func (h *Handler) GetProducts(c *fiber.Ctx) error {
	products, err := h.store.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(products)
}
