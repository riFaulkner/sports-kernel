package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/standings"
)

func (r *standingsQueriesResolver) WeekStandings(ctx context.Context, obj *standings.StandingsQueries, season int, week *int) ([]*standings.Standings, error) {
	return r.StandingsService.GetStandings(season, week)
}

// StandingsQueries returns generated.StandingsQueriesResolver implementation.
func (r *Resolver) StandingsQueries() generated.StandingsQueriesResolver {
	return &standingsQueriesResolver{r}
}

type standingsQueriesResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
