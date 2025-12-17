package user

import (
	"ecom/config"
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

	var payload types.LoginRequest
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
	user, err2 := h.store.GetUserByEmail(c.UserContext(), payload.Email)

	if err2 != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}

	if !auth.ComparePasswords(user.Password, []byte(payload.Password)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}
	secret := []byte(config.Envs.JWTSecret)
	jwt, err3 := auth.CreateJWT(secret, user.ID)

	if err3 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err3.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": jwt,
	})
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
	if _, err := h.store.GetUserByEmail(c.UserContext(), payload.Email); err == nil {
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

	err2 := h.store.CreateUser(c.UserContext(), &types.User{
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
