package user

import (
	"ecom/service/auth"
	"ecom/types"
	"ecom/utils"

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

	//Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//check if the user exists
	if _, err := h.store.GetUserByEmail(payload.Email); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user already exists",
		})
	}
	//if not, create the user
	password, err := auth.HashPassword(payload.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err2 := h.store.CreateUser(&types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  password,
	})
	if err2 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err2.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
}
