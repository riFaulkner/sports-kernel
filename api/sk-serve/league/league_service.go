package league

import (
	"context"
)

const (
	SalaryCap         = 200_000_000
	MaxContractLength = 4
)

type LeagueRepository interface {
	GetAll(ctx context.Context) ([]*League, error)
	GetByLeagueId(ctx context.Context, leagueID string) (*League, error)
}
