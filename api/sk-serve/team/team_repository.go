package team

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user"
)

type TeamRepository interface {
	AddUserToTeam(ctx context.Context, accessCode string, ownerID string) (*user.UserPreferences, error)
	Create(ctx context.Context, leagueId string, team NewTeam) (*Team, error)
	GenerateAccessCode(ctx context.Context, leagueId string, teamId string, role string) (string, error)
	GetAllLeagueTeams(ctx context.Context, leagueId string) ([]*Team, error)
	GetTeamById(ctx context.Context, leagueId string, teamId string) (*Team, error)
	GetTeamByOwnerID(ctx context.Context, leagueID string, ownerID string) (*Team, bool)
	UpdateTeamContractMetaData(ctx context.Context, leagueId string, teamContracts []*contract.Contract) error
}
