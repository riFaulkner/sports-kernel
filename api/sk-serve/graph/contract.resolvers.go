package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

func (r *contractResolver) Player(ctx context.Context, obj *contract.Contract) (*model.PlayerNfl, error) {
	return r.PlayerService.GetPlayerById(ctx, &obj.PlayerID)
}

func (r *contractMutationsResolver) Drop(ctx context.Context, obj *contract.ContractMutations, leagueID string, teamID string, contractID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *contractMutationsResolver) Restructure(ctx context.Context, obj *contract.ContractMutations, leagueID string, teamID string, contractID string) (*contract.Contract, error) {
	panic(fmt.Errorf("not implemented"))
}

// Contract returns generated.ContractResolver implementation.
func (r *Resolver) Contract() generated.ContractResolver { return &contractResolver{r} }

// ContractMutations returns generated.ContractMutationsResolver implementation.
func (r *Resolver) ContractMutations() generated.ContractMutationsResolver {
	return &contractMutationsResolver{r}
}

type contractResolver struct{ *Resolver }
type contractMutationsResolver struct{ *Resolver }
