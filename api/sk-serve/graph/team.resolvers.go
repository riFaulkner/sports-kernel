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

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }
