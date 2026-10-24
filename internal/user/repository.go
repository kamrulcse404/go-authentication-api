package user

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	FindByID(ctx context.Context, id int64) (User, error)
}
