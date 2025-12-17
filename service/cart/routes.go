package cart

import (
	"ecom/service/auth"
	"ecom/types"
	"ecom/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:        store,
		productStore: productStore,
		userStore:    userStore,
	}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	r.Post("/checkout", auth.WithJWT(h.Checkout, h.userStore))
}

func (h *Handler) Checkout(c *fiber.Ctx) error {
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var cart types.CartCheckoutRequest
	if err := c.BodyParser(&cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := utils.Validate.Struct(cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// get products
	productIDs, err := getCartItemIDs(cart.Items)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	products, err := h.productStore.GetProductsByID(c.UserContext(), productIDs)
	// create order
	orderID, amount, err := h.createOrder(c.UserContext(), products, cart.Items, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orderID": orderID,
		"amount":  amount,
	})
}
