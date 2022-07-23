package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
)

func (r *leagueResolver) Teams(ctx context.Context, obj *league.League, search *model.LeagueTeamFiltering) ([]*team.Team, error) {
	if search != nil {
		var singleTeam *team.Team
		var err error

		if search.TeamID != nil {
			singleTeam, err = r.TeamService.GetTeamById(ctx, obj.ID, *search.TeamID)
		} else if search.OwnerID != nil {
			singleTeam, err = r.TeamService.GetTeamByOwnerID(ctx, obj.ID, *search.OwnerID)
		}

		if err != nil || singleTeam == nil {
			return make([]*team.Team, 0, 0), err
		}

		wrapped := make([]*team.Team, 1, 1)
		wrapped[0] = singleTeam
		return wrapped, nil
	}
	return r.TeamService.GetAllLeagueTeams(ctx, obj.ID)
}

// League returns generated.LeagueResolver implementation.
func (r *Resolver) League() generated.LeagueResolver { return &leagueResolver{r} }

type leagueResolver struct{ *Resolver }
