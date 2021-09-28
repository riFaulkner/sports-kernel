package contract

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type Contract interface {
	GetAll(ctx context.Context, leagueID string, teamID string) ([]*model.Contract, error)
}
