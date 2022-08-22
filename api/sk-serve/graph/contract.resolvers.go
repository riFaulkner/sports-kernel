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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *contractMutationsResolver) Test(ctx context.Context, obj *contract.ContractMutations, input *string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
