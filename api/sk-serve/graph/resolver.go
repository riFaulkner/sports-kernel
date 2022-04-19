package graph

import (
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/db"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/playerNFL"
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
	ContractResolver contract.Contract
	PlayerResolver   playernfl.PlayerNfl
}

func Initialize(client firestore.Client) generated.Config {

	r := Resolver{}
	r.ContractResolver = &db.ContractImpl{Client: client}
	r.LeagueResolver = &db.LeagueImpl{Client: client}
	r.TeamResolver = &db.TeamImpl{Client: client}
	r.UserResolver = &db.UserImpl{Client: client}
	r.PlayerResolver = &db.PlayerImpl{Client: client}

	return generated.Config{
		Resolvers: &r,
	}

}
