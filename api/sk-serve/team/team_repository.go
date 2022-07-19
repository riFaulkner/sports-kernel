package team

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
)

type TeamRepository interface {
	Create(ctx context.Context, leagueId string, team NewTeam) (*Team, error)
	GetAllLeagueTeams(ctx context.Context, leagueId string) ([]*Team, error)
	GetTeamById(ctx context.Context, leagueId string, teamId string) (*Team, error)
	GetTeamByOwnerID(ctx context.Context, leagueID string, ownerID string) (*Team, bool)
	UpdateTeamContractMetaData(ctx context.Context, leagueId string, teamContracts []*contract.Contract) error
}
