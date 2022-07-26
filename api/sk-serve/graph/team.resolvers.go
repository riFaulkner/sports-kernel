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

func (r *deadCapResolver) Contract(ctx context.Context, obj *team.DeadCap) (*contract.Contract, error) {
	if leagueID, ok := getLeagueIDFromContext(ctx); ok {
		contract, ok := r.ContractResolver.GetById(ctx, leagueID, obj.AssociatedContractID)
		if ok {
			return contract, nil
		}
	}
	return nil, gqlerror.Errorf("Error getting contract for dead cap reference")
}

func (r *teamResolver) ActiveContracts(ctx context.Context, obj *team.Team) ([]*contract.Contract, error) {
	leagueID := graphql.GetOperationContext(ctx).Variables["leagueId"]
	if str, ok := leagueID.(string); ok {
		return r.ContractResolver.GetAllActiveTeamContracts(ctx, str, obj.ID)
	}
	return nil, gqlerror.Errorf("Error getting leagueId to retrieve active contract")
}

func getLeagueIDFromContext(ctx context.Context) (string, bool) {
	leagueID := graphql.GetOperationContext(ctx).Variables["leagueId"]
	str, ok := leagueID.(string)
	return str, ok
}

// DeadCap returns generated.DeadCapResolver implementation.
func (r *Resolver) DeadCap() generated.DeadCapResolver { return &deadCapResolver{r} }

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

type deadCapResolver struct{ *Resolver }
type teamResolver struct{ *Resolver }
