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

func (r *contractInputResolver) PlayerPosition(ctx context.Context, obj *contract.ContractInput, data *model.PlayerPosition) error {
	panic(fmt.Errorf("not implemented"))
}

// Contract returns generated.ContractResolver implementation.
func (r *Resolver) Contract() generated.ContractResolver { return &contractResolver{r} }

// ContractInput returns generated.ContractInputResolver implementation.
func (r *Resolver) ContractInput() generated.ContractInputResolver { return &contractInputResolver{r} }

type contractResolver struct{ *Resolver }
type contractInputResolver struct{ *Resolver }
