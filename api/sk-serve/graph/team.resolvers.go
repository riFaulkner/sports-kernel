package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *teamResolver) ActiveContracts(ctx context.Context, obj *team.Team) ([]*contract.Contract, error) {
	leagueID := graphql.GetOperationContext(ctx).Variables["leagueId"]
	if str, ok := leagueID.(string); ok {
		return r.ContractResolver.GetAllActiveTeamContracts(ctx, str, obj.ID)
	}
	return nil, gqlerror.Errorf("Error getting leagueId to retrieve active contract")
}

func (r *teamMutationsResolver) AddDeadCap(ctx context.Context, obj *team.TeamMutations, leagueID string, teamID string, deadCap team.DeadCapInput) (bool, error) {
	result, err := r.TeamService.AddDeadCapToTeam(ctx, leagueID, teamID, deadCap)

	if err != nil {
		return false, err
	}
	teamContracts, err := r.ContractResolver.GetAllActiveTeamContracts(ctx, leagueID, teamID)
	if err != nil {
		return true, gqlerror.Errorf("Unable to recalculate team metadata")
	}
	r.TeamService.UpdateTeamContractMetaData(ctx, leagueID, teamContracts)

	return result, err
}

func (r *teamQueriesResolver) ByTeamID(ctx context.Context, obj *team.TeamQueries, teamID string) (*team.Team, error) {
	team, err := r.TeamService.GetTeamById(ctx, obj.LeagueID, teamID)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (r *teamQueriesResolver) ByOwnerID(ctx context.Context, obj *team.TeamQueries, ownerID string) (*team.Team, error) {
	return r.TeamService.GetTeamByOwnerID(ctx, obj.LeagueID, ownerID)
}

func (r *teamQueriesResolver) TeamIds(ctx context.Context, obj *team.TeamQueries, teamIds []string) ([]*team.Team, error) {
	return r.TeamService.GetTeamsByIds(ctx, obj.LeagueID, teamIds)
}

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

// TeamMutations returns generated.TeamMutationsResolver implementation.
func (r *Resolver) TeamMutations() generated.TeamMutationsResolver { return &teamMutationsResolver{r} }

// TeamQueries returns generated.TeamQueriesResolver implementation.
func (r *Resolver) TeamQueries() generated.TeamQueriesResolver { return &teamQueriesResolver{r} }

type teamResolver struct{ *Resolver }
type teamMutationsResolver struct{ *Resolver }
type teamQueriesResolver struct{ *Resolver }
