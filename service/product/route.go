package product

import (
	"ecom/types"
	"ecom/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	r.Get("/", h.GetProducts)
	r.Post("/", h.CreateProduct)
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

func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	var payload types.ProductRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//Check if the product exists
	isAvailable, err := h.store.CheckProduct(payload.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if isAvailable {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "product already exists",
		})
	}
	err2 := h.store.CreateProduct(&types.Product{
		Name:        payload.Name,
		Price:       payload.Price,
		Description: payload.Description,
		Image:       "image.png",
		Quantity:    payload.Quantity,
	})
	if err2 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err2.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
}
