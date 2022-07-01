package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

func (r *contractResolver) Player(ctx context.Context, obj *contract.Contract) (*model.PlayerNfl, error) {
	return r.PlayerResolver.GetPlayerById(ctx, &obj.PlayerID) //return dataloader.GetPlayers(ctx, &obj.PlayerID)
}

// Contract returns generated.ContractResolver implementation.
func (r *Resolver) Contract() generated.ContractResolver { return &contractResolver{r} }

type contractResolver struct{ *Resolver }
