package league

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type League interface {
	GetAll(ctx context.Context) ([]*model.League, error)
}
