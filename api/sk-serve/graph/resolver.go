package graph

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/db"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	User     user.User
	League   league.League
	Team     team.Team
	Contract contract.Contract
}

func Initialize(ctx context.Context) generated.Config {

	client := firestore.NewClient(ctx)

	r := Resolver{}
	r.Contract = &db.ContractImpl{Client: client}
	r.League = &db.LeagueImpl{Client: client}
	r.Team = &db.TeamImpl{Client: client}
	r.User = &db.UserImpl{Client: client}

	return generated.Config{
		Resolvers: &r,
	}

}
