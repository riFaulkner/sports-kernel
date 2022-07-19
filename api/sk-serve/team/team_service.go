package team

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type TeamService struct {
	TeamRepository TeamRepository
}

func (t TeamService) GetTeamByOwnerID(ctx context.Context, leagueID string, ownerID string) (*Team, error) {
	team, ok := t.TeamRepository.GetTeamByOwnerID(ctx, leagueID, ownerID)
	if !ok {
		return nil, gqlerror.Errorf("Error occurred getting ownerID: %v teams in league: %v", ownerID, leagueID)
	}
	return team, nil
}

func (t TeamService) GetAllLeagueTeams(ctx context.Context, leagueId string) ([]*Team, error) {
	return t.TeamRepository.GetAllLeagueTeams(ctx, leagueId)
}
func (t TeamService) GetTeamById(ctx context.Context, leagueId string, teamId string) (*Team, error) {
	return t.TeamRepository.GetTeamById(ctx, leagueId, teamId)
}
func (t TeamService) Create(ctx context.Context, leagueId string, team NewTeam) (*Team, error) {
	return t.TeamRepository.Create(ctx, leagueId, team)
}
func (t TeamService) UpdateTeamContractMetaData(ctx context.Context, leagueId string, teamContracts []*contract.Contract) error {
	return t.TeamRepository.UpdateTeamContractMetaData(ctx, leagueId, teamContracts)
}
