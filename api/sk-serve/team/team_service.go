package team

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type Team interface {
	//Change this to a path reference to the league or whatever object we're looking for
	GetAll(ctx context.Context, leagueId string) ([]*model.Team, error)
	GetTeamById(ctx context.Context, leagueId string, teamId string) (*model.Team, error)
	Create(ctx context.Context, leagueId string, team model.Team) error
	UpdateTeamContractMetaData(ctx context.Context, leagueId string, teamContracts []*model.Contract) error
}
