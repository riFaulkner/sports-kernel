package playernfl

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type PlayerNfl interface {
	GetAll(ctx context.Context, numberOfResults *int) ([]*model.PlayerNfl, error)
	GetPlayerById(ctx context.Context, playerId *string) (*model.PlayerNfl, error)
}
