package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	cache "github.com/patrickmn/go-cache"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/db"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	playernfl "github.com/rifaulkner/sports-kernel/api/sk-serve/nfl"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/post"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/scoring"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user/onboarding"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserResolver          user.UserService
	LeagueResolver        league.LeagueRepository
	TeamService           team.TeamService
	ContractResolver      contract.Resolver
	PlayerService         playernfl.PlayerService
	PostResolver          post.LeaguePost
	ScoringService        scoring.Service
	UserOnBoardingService onboarding.UserOnboardingService
}

func Initialize(client firestore.Client) generated.Config {
	cache := cache.New(5*time.Minute, 30*time.Minute)

	transactionImpl := db.TransactionImpl{Client: client}
	teamImpl := db.TeamRepositoryImpl{Client: client}

	userService := initializeUserService(client)
	teamService := initializeTeamService(client)
	leagueService := &db.LeagueImpl{Client: client}

	r := Resolver{}
	r.ContractResolver = &db.ContractImpl{
		Client:          client,
		TeamImpl:        teamImpl,
		TransactionImpl: transactionImpl,
	}
	r.LeagueResolver = leagueService
	r.TeamService = teamService
	r.UserResolver = userService
	r.PlayerService = initializePlayerService(client)
	r.PostResolver = &db.PostImpl{Client: client}
	r.ScoringService = *scoring.NewScoringService(cache)
	r.UserOnBoardingService = initializeUserOnBoardingService(userService, teamService, leagueService)

	return generated.Config{
		Resolvers: &r,
	}
}

func initializePlayerService(client firestore.Client) playernfl.PlayerService {
	return *playernfl.NewPlayerService(&db.PlayerRepositoryImpl{
		Client: client,
	})
}

func initializeTeamService(client firestore.Client) team.TeamService {
	return team.TeamService{
		TeamRepository: &db.TeamRepositoryImpl{
			Client: client,
		},
	}
}

func initializeUserService(client firestore.Client) user.UserService {
	return user.UserService{
		UserRepository: &db.UserImpl{
			Client: client,
		},
	}
}

func initializeUserOnBoardingService(userService user.UserService, teamService team.TeamService, leagueService *db.LeagueImpl) onboarding.UserOnboardingService {
	return onboarding.UserOnboardingService{
		UserService:   userService,
		TeamService:   teamService,
		LeagueService: leagueService,
	}
}
