package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user"
)

func (r *userPreferencesLeagueSnippetResolver) RoleInLeague(ctx context.Context, obj *user.UserPreferencesLeagueSnippet) (model.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserPreferencesLeagueSnippet returns generated.UserPreferencesLeagueSnippetResolver implementation.
func (r *Resolver) UserPreferencesLeagueSnippet() generated.UserPreferencesLeagueSnippetResolver {
	return &userPreferencesLeagueSnippetResolver{r}
}

type userPreferencesLeagueSnippetResolver struct{ *Resolver }
