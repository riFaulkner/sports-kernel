package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		OwnerName: input.OwnerName,
		Email:     input.Email,
		Avatar:    input.Avatar,
	}

	err := r.UserResolver.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) CreateTeam(ctx context.Context, leagueID *string, input model.NewTeam) (*model.Team, error) {
	currentContractsMetadataDefault := model.ContractsMetadata{
		TotalUtilizedCap:  0,
		TotalAvailableCap: 200000000,
		QbUtilizedCap: &model.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		RbUtilizedCap: &model.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		WrUtilizedCap: &model.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		TeUtilizedCap: &model.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
	}
	team := model.Team{
		TeamName:                 input.TeamName,
		FoundedDate:              time.Now(),
		CurrentContractsMetadata: &currentContractsMetadataDefault,
	}

	err := r.TeamResolver.Create(ctx, *leagueID, team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (r *mutationResolver) CreateContract(ctx context.Context, leagueID *string, input *model.ContractInput) (*model.Contract, error) {
	document, err := r.ContractResolver.CreateContract(ctx, *leagueID, input)

	if err != nil {
		return nil, err
	}

	teamContracts, err := r.ContractResolver.GetAll(ctx, *leagueID, document.TeamID)
	if err != nil {
		log.Println("Failed to update contract metadata")
		return nil, err
	}
	err = r.TeamResolver.UpdateTeamContractMetaData(ctx, *leagueID, teamContracts)

	return document, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserResolver.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) Leagues(ctx context.Context) ([]*model.League, error) {
	leagues, err := r.Resolver.LeagueResolver.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return leagues, nil
}

func (r *queryResolver) League(ctx context.Context, leagueID *string) (*model.League, error) {
	league, err := r.Resolver.LeagueResolver.GetByLeagueId(ctx, *leagueID)
	if err != nil {
		return nil, err
	}

	return league, nil
}

func (r *queryResolver) Teams(ctx context.Context, leagueID *string) ([]*model.Team, error) {
	teams, err := r.TeamResolver.GetAll(ctx, *leagueID)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (r *queryResolver) TeamContracts(ctx context.Context, leagueID *string, teamID *string) ([]*model.Contract, error) {
	contracts, err := r.ContractResolver.GetAll(ctx, *leagueID, *teamID)
	if err != nil {
		return nil, err
	}

	for _, s := range contracts {
		s.Player, err = r.Player(ctx, &s.PlayerID)
		if err != nil {
			graphql.AddErrorf(ctx, fmt.Sprintf("Error getting player info for playerID: %s", s.PlayerID))
		}
	}

	return contracts, nil
}

func (r *queryResolver) Players(ctx context.Context, numOfResults *int) ([]*model.PlayerNfl, error) {
	players, err := r.PlayerResolver.GetAll(ctx, numOfResults)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (r *queryResolver) Player(ctx context.Context, playerID *string) (*model.PlayerNfl, error) {
	player, err := r.PlayerResolver.GetPlayerById(ctx, playerID)

	if err != nil {
		return nil, err
	}

	return player, nil
}

func (r *queryResolver) UserPreferences(ctx context.Context, userID *string) (*model.UserPreferences, error) {
	userPreferences, err := r.UserResolver.GetUserPreferences(ctx, *userID)
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
