package user

import (
	"context"
	"errors"
	"testing"
)

func TestUserServiceGetByID(t *testing.T) {
	repo := NewMemoryRepository()
	service := NewService(repo)

	user, err := service.GetByID(context.Background(), 1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("expected user ID 1, got %d", user.ID)
	}

	if user.Email != "alice@example.com" {
		t.Errorf("unexpected email: %s", user.Email)
	}
}

func TestUserServiceGetByIDNotFound(t *testing.T) {
	repo := NewMemoryRepository()
	service := NewService(repo)

	_, err := service.GetByID(context.Background(), 999)

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}
