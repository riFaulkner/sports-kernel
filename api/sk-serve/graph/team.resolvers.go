package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

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

func (r *teamLiabilitiesResolver) DeadCap(ctx context.Context, obj *team.TeamLiabilities) ([]*team.DeadCapYear, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *teamMutationsResolver) AddDeadCap(ctx context.Context, obj *team.TeamMutations, leagueID string, teamID string, deadCap team.DeadCapInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

// TeamLiabilities returns generated.TeamLiabilitiesResolver implementation.
func (r *Resolver) TeamLiabilities() generated.TeamLiabilitiesResolver {
	return &teamLiabilitiesResolver{r}
}

// TeamMutations returns generated.TeamMutationsResolver implementation.
func (r *Resolver) TeamMutations() generated.TeamMutationsResolver { return &teamMutationsResolver{r} }

type teamResolver struct{ *Resolver }
type teamLiabilitiesResolver struct{ *Resolver }
type teamMutationsResolver struct{ *Resolver }
