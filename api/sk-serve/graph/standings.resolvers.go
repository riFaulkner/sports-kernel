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
