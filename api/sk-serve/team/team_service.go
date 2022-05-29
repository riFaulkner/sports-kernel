package team

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type Team interface {
	GetAll(ctx context.Context, leagueId string) ([]*model.Team, error)
	GetTeamById(ctx context.Context, leagueId string, teamId string) (*model.Team, error)
	Create(ctx context.Context, leagueId string, team model.NewTeam) (*model.Team, error)
	UpdateTeamContractMetaData(ctx context.Context, leagueId string, teamContracts []*contract.Contract) error
}
