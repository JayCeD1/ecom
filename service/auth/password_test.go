package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}
	if hash == "" {
		t.Error("expected has to be non-empty")
	}
	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}
	if !ComparePasswords(hash, []byte("password")) {
		t.Error("expected password to match")
	}
	if ComparePasswords(hash, []byte("wrong")) {
		t.Error("expected password to not match")
	}
}
