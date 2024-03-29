package playernfl

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type PlayerRepository interface {
	Create(ctx context.Context, player model.PlayerNfl) (*model.PlayerNfl, error)
	GetAll(ctx context.Context) ([]*model.PlayerNfl, bool)
	GetPlayersByPosition(ctx context.Context, position model.PlayerPosition) ([]*model.PlayerNfl, bool)
	GetPlayersWithLimit(ctx context.Context, numberOfResults int) ([]*model.PlayerNfl, bool)
	GetPlayerById(ctx context.Context, playerId *string) (*model.PlayerNfl, bool)
}
