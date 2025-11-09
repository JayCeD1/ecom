package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")
	jwt, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("error creating jwt: %v", err)
	}
	if jwt == "" {
		t.Error("expected jwt to be non-empty")
	}
}
