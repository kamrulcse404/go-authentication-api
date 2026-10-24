package user

import (
	"context"
	"sync"
)

type MemoryRepository struct {
	mu    sync.RWMutex
	users map[int64]User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: map[int64]User{
			1: {
				ID:    1,
				Name:  "Alice",
				Email: "alice@example.com",
			},
			2: {
				ID:    2,
				Name:  "Bob",
				Email: "bob@example.com",
			},
		},
	}
}

func (r *MemoryRepository) FindByID(ctx context.Context, id int64) (User, error) {
	err := ctx.Err()
	if err != nil {
		return User{}, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]

	if !exists {
		return User{}, ErrNotFound
	}

	return user, nil
}
