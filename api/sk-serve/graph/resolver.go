package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/db"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	playernfl "github.com/rifaulkner/sports-kernel/api/sk-serve/nfl"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/post"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserResolver     user.User
	LeagueResolver   league.League
	TeamResolver     team.Team
	ContractResolver contract.Resolver
	PlayerService    playernfl.PlayerService
	PostResolver     post.LeaguePost
}

func Initialize(client firestore.Client) generated.Config {
	transactionImpl := db.TransactionImpl{Client: client}
	teamImpl := db.TeamImpl{Client: client}

	r := Resolver{}
	r.ContractResolver = &db.ContractImpl{
		Client:          client,
		TeamImpl:        teamImpl,
		TransactionImpl: transactionImpl,
	}

	r.LeagueResolver = &db.LeagueImpl{Client: client}
	r.TeamResolver = &teamImpl
	r.UserResolver = &db.UserImpl{Client: client}
	r.PlayerService = initializePlayerService(client)
	r.PostResolver = &db.PostImpl{Client: client}

	return generated.Config{
		Resolvers: &r,
	}
}

func initializePlayerService(client firestore.Client) playernfl.PlayerService {
	return playernfl.PlayerService{
		PlayerRepository: &db.PlayerRepositoryImpl{
			Client: client,
		},
	}
}
