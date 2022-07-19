package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user"
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

func (r *mutationResolver) CreateTeam(ctx context.Context, leagueID *string, input team.NewTeam) (*team.Team, error) {
	team, err := r.TeamService.Create(ctx, *leagueID, input)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (r *mutationResolver) UpdateTeamMetaData(ctx context.Context, leagueID string, teamID string) (*team.Team, error) {
	contracts, err := r.ContractResolver.GetAllActiveTeamContracts(ctx, leagueID, teamID)
	if err != nil {
		return nil, err
	}

	err = r.TeamService.UpdateTeamContractMetaData(ctx, leagueID, contracts)
	if err != nil {
		return nil, err
	}

	team, err := r.TeamService.GetTeamById(ctx, leagueID, teamID)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (r *mutationResolver) CreateContract(ctx context.Context, leagueID string, input contract.ContractInput) (*contract.Contract, error) {
	document, err := r.ContractResolver.CreateContract(ctx, leagueID, input)

	if err != nil {
		return nil, err
	}

	teamContracts, err := r.ContractResolver.GetAllActiveTeamContracts(ctx, leagueID, document.TeamID)
	if err != nil {
		log.Println("Failed to update contract metadata")
		return nil, err
	}
	err = r.TeamService.UpdateTeamContractMetaData(ctx, leagueID, teamContracts)

	return document, nil
}

func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.NewPlayerNfl) (*model.PlayerNfl, error) {
	player, err := r.PlayerService.CreatePlayer(ctx, input)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, leagueID string, input *model.NewLeaguePost) (*model.LeaguePost, error) {
	post, err := r.PostResolver.Create(ctx, leagueID, *input)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *mutationResolver) AddComment(ctx context.Context, leagueID string, postID string, input *model.NewPostComment) (*model.PostComment, error) {
	comment, err := r.PostResolver.AddComment(ctx, leagueID, postID, *input)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *mutationResolver) CreateUserRole(ctx context.Context, leagueID *string, newUserRole *model.NewUserRole) (*model.UserRoles, error) {
	return r.UserResolver.CreateUserRole(ctx, newUserRole)
}

func (r *mutationResolver) ContractActionDrop(ctx context.Context, leagueID string, teamID string, contractID string) (bool, error) {
	return r.ContractResolver.DropContract(ctx, leagueID, teamID, contractID)
}

func (r *mutationResolver) ContractActionRestructure(ctx context.Context, leagueID string, restructureDetails contract.ContractRestructureInput) (*contract.Contract, error) {
	return r.ContractResolver.RestructureContract(ctx, &leagueID, &restructureDetails)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserResolver.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) Leagues(ctx context.Context) ([]*league.League, error) {
	leagues, err := r.Resolver.LeagueResolver.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return leagues, nil
}

func (r *queryResolver) League(ctx context.Context, leagueID *string) (*league.League, error) {
	league, err := r.Resolver.LeagueResolver.GetByLeagueId(ctx, *leagueID)
	if err != nil {
		return nil, err
	}

	return league, nil
}

func (r *queryResolver) LeagueContracts(ctx context.Context, leagueID string) ([]*contract.Contract, error) {
	return r.ContractResolver.GetAllLeagueContracts(ctx, leagueID)
}

func (r *queryResolver) LeagueContractsByOwnerID(ctx context.Context, leagueID string, ownerID string) ([]*contract.Contract, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Teams(ctx context.Context, leagueID *string) ([]*team.Team, error) {
	teams, err := r.TeamService.GetAllLeagueTeams(ctx, *leagueID)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (r *queryResolver) TeamByID(ctx context.Context, leagueID string, teamID string) (*team.Team, error) {
	team, err := r.TeamService.GetTeamById(ctx, leagueID, teamID)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (r *queryResolver) TeamByOwnerID(ctx context.Context, leagueID string, ownerID string) (*team.Team, error) {
	return r.TeamService.GetTeamByOwnerID(ctx, leagueID, ownerID)
}

func (r *queryResolver) TeamContracts(ctx context.Context, leagueID *string, teamID *string) ([]*contract.Contract, error) {
	contracts, err := r.ContractResolver.GetAllActiveTeamContracts(ctx, *leagueID, *teamID)
	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (r *queryResolver) Player(ctx context.Context, playerID *string) (*model.PlayerNfl, error) {
	return r.PlayerService.GetPlayerById(ctx, playerID)
}

func (r *queryResolver) Players(ctx context.Context, numOfResults *int) ([]*model.PlayerNfl, error) {
	return r.PlayerService.GetAllPlayers(ctx, numOfResults)
}

func (r *queryResolver) PlayersByPosition(ctx context.Context, position model.PlayerPosition) ([]*model.PlayerNfl, error) {
	return r.PlayerService.GetPlayersByPosition(ctx, position)
}

func (r *queryResolver) Posts(ctx context.Context, leagueID string, numOfResults *int) ([]*model.LeaguePost, error) {
	posts, err := r.PostResolver.GetAll(ctx, leagueID, numOfResults)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *queryResolver) Comments(ctx context.Context, leagueID string, postID string) ([]*model.PostComment, error) {
	comments, err := r.PostResolver.GetComments(ctx, leagueID, postID)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *queryResolver) UserPreferences(ctx context.Context, userID *string) (*user.UserPreferences, error) {
	userPreferences, err := r.UserResolver.GetUserPreferences(ctx, *userID)
	if err != nil {
		log.Printf("Error attempting to resolve user preferences, %s", err)
		return nil, err
	}

	return userPreferences, nil
}

func (r *queryResolver) GetUserRoles(ctx context.Context, userID *string) ([]*model.UserRoles, error) {
	return r.UserResolver.GetUserRoles(ctx, userID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
