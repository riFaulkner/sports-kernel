package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		OwnerName: input.OwnerName,
		Email:     input.Email,
		Avatar:    input.Avatar,
	}

	err := r.User.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.User.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) Leagues(ctx context.Context) ([]*model.League, error) {
	leagues, err := r.League.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return leagues, nil
}

func (r *queryResolver) Teams(ctx context.Context, leagueID *string) ([]*model.Team, error) {
	teams, err := r.Team.GetAll(ctx, *leagueID)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (r *queryResolver) UserPreferences(ctx context.Context, userID *string) (*model.UserPreferences, error) {
	userPreferences, err := r.User.GetUserPreferences(ctx, *userID)
	if err != nil {
		log.Printf("Error attempting to resolve user preferences, %s", err)
		return nil, err
	}

	return userPreferences, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
