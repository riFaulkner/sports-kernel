package team

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type Team interface {
	GetAll(ctx context.Context, leagueId string) ([]*model.Team, error)
}
