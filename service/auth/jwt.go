package auth

import (
	"ecom/config"
	"ecom/types"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

type JWTClaims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		UserID: strconv.Itoa(userID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		},
	})
	return claims.SignedString(secret)
}

// WithJWT validates the Authorization: Bearer <token>, loads the user, and sets c.Locals("userID", int)
func WithJWT(next fiber.Handler, store types.UserStore) fiber.Handler {
	return func(c fiber.Ctx) error {
		// get token from user request
		tokenStr := getTokenFromRequest(c)
		// validate the jwt
		token, err := validateToken(tokenStr)
		if err != nil || !token.Valid {
			log.Printf("invalid token: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		claims, ok := token.Claims.(*JWTClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
		}

		uid, err := strconv.Atoi(claims.UserID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id in token"})
		}

		// if we need to fetch the user from the database, do it
		// Optional: ensure user still exists
		if _, err2 := store.GetUserByID(uid); err2 != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user not found"})
		}

		// set context "UserID" to the user id (request scoped)
		c.Locals(UserKey, uid)
		return next(c)
	}
}

func getTokenFromRequest(c fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}
	tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	return tokenStr
}

func validateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIDFromContext(c fiber.Ctx) (int, error) {
	userID, ok := c.Locals(UserKey).(int)
	if !ok {
		return 0, fiber.ErrNotFound
	}
	return userID, nil
}
