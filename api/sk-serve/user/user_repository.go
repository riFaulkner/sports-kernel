package user

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type UserRepository interface {
	AddLeagueToUserPreferences(ctx context.Context, userID string, leagueSnippet UserPreferencesLeagueSnippet) bool
	GetAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, user UserPreferences) error
	GetUserPreferences(ctx context.Context, userId string) (*UserPreferences, error)
	CreateUserRole(cxt context.Context, newRole *model.NewUserRole) (*model.UserRoles, error)
	GetUserRoles(ctx context.Context, userID *string) ([]*model.UserRoles, error)
}
