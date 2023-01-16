package team

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
)

type TeamRepository interface {
	AddDeadCapToTeam(ctx context.Context, leagueID string, teamID string, deadCap DeadCap) bool
	AddUserToTeam(ctx context.Context, leagueID string, teamID string, ownerID string) bool
	Create(ctx context.Context, leagueId string, team NewTeam) (*Team, error)
	AddAccessCode(ctx context.Context, leagueId string, teamId string, accessCode string) error
	GetAllLeagueTeams(ctx context.Context, leagueId string) ([]*Team, error)
	GetTeamById(ctx context.Context, leagueId string, teamId string) (*Team, error)
	GetTeamByOwnerID(ctx context.Context, leagueID string, ownerID string) (*Team, bool)
	GetTeamsByIds(ctx context.Context, leagueID string, teamIds []string) ([]*Team, error)
	RemoveAccessCode(ctx context.Context, leagueID string, teamID string, accessCode string) bool
	UpdateTeamContractMetaData(ctx context.Context, leagueId string, teamContracts []*contract.Contract) error
}
