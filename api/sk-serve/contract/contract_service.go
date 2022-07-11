package contract

import (
	"context"
)

type Resolver interface {
	GetAllLeagueContracts(ctx context.Context, leagueID string) ([]*Contract, error)
	GetAllTeamContracts(ctx context.Context, leagueID string, teamID string) ([]*Contract, error)
	GetAllActiveTeamContracts(ctx context.Context, leagueID string, teamID string) ([]*Contract, error)
	CreateContract(ctx context.Context, leagueId string, contractInput ContractInput) (*Contract, error)
	DropContract(ctx context.Context, leagueID string, teamID string, contractID string) (bool, error)
	RestructureContract(ctx context.Context, leagueID *string, restructureDetails *ContractRestructureInput) (*Contract, error)
}
