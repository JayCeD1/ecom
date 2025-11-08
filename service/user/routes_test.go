package user

import (
	"bytes"
	"ecom/types"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.UserRequest{
			FirstName: "user",
			LastName:  "25",
			Email:     "",
			Password:  "passed",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		app := fiber.New()
		app.Post("/register", handler.Register)
		resp, err := app.Test(req, -1)
		if err != nil {
			t.Fatalf("app.Test error: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil

}

func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}
