package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

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

func (r *mutationResolver) CreateTeam(ctx context.Context, leagueID *string, input model.NewTeam) (*model.Team, error) {
	team := model.Team{
		TeamName:    input.TeamName,
		FoundedDate: time.Now(),
	}

	err := r.Team.Create(ctx, *leagueID, team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.User.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) Leagues(ctx context.Context) ([]*model.League, error) {
	leagues, err := r.Resolver.League.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return leagues, nil
}

func (r *queryResolver) LeagueByID(ctx context.Context, leagueID *string) (*model.League, error) {
	league, err := r.Resolver.League.GetByLeagueId(ctx, *leagueID)
	if err != nil {
		return nil, err
	}

	return league, nil
}

func (r *queryResolver) Teams(ctx context.Context, leagueID *string) ([]*model.Team, error) {
	teams, err := r.Team.GetAll(ctx, *leagueID)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (r *queryResolver) Contracts(ctx context.Context, leagueID *string, teamID *string) ([]*model.Contract, error) {
	contracts, err := r.Contract.GetAll(ctx, *leagueID, *teamID)
	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (r *queryResolver) Players(ctx context.Context, numOfResults *int) ([]*model.PlayerNfl, error) {
	players, err := r.Player.GetAll(ctx, numOfResults)
	if err != nil {
		return nil, err
	}

	return players, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) League(ctx context.Context, leagueID *string) (*model.League, error) {
	//league, err := r.League.GetByLeagueId(ctx, leagueID)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return league, nil
	panic("help")
}
