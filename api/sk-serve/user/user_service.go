package user

import (
	"context"

	"github.com/riFaulkner/sports-kernel/sports-kernel/api/sk-serve/graph/model"
)

type User interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, user model.User) error
}
