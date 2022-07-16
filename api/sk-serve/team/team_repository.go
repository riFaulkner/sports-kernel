package team

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type TeamRepository interface {
	Create(ctx context.Context, leagueId string, team model.NewTeam) (*model.Team, error)
	GetAllLeagueTeams(ctx context.Context, leagueId string) ([]*model.Team, error)
	GetTeamById(ctx context.Context, leagueId string, teamId string) (*model.Team, error)
	GetTeamByOwnerID(ctx context.Context, leagueID string, ownerID string) (*model.Team, bool)
	UpdateTeamContractMetaData(ctx context.Context, leagueId string, teamContracts []*contract.Contract) error
}
