package graph

import (
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	User   user.User
	League league.League
	Team   team.Team
}
