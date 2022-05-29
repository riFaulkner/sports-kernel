package contract

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type Resolver interface {
	GetAllLeagueContracts(ctx context.Context, leagueID string) ([]*Contract, error)
	GetAllTeamContracts(ctx context.Context, leagueID string, teamID string) ([]*Contract, error)
	CreateContract(ctx context.Context, leagueId string, input *model.ContractInput) (*Contract, error)
	RestructureContract(ctx context.Context, leagueID *string, restructureDetails *model.ContractRestructureInput) (*Contract, error)
}
