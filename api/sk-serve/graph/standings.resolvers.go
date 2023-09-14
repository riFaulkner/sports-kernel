package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/standings"
)

func (r *standingsResolver) PointsFor(ctx context.Context, obj *standings.Standings) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *standingsResolver) PointsAgaints(ctx context.Context, obj *standings.Standings) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *standingsQueriesResolver) WeekStandings(ctx context.Context, obj *standings.StandingsQueries, season int, week *int) ([]*standings.Standings, error) {
	panic(fmt.Errorf("not implemented"))
}

// Standings returns generated.StandingsResolver implementation.
func (r *Resolver) Standings() generated.StandingsResolver { return &standingsResolver{r} }

// StandingsQueries returns generated.StandingsQueriesResolver implementation.
func (r *Resolver) StandingsQueries() generated.StandingsQueriesResolver {
	return &standingsQueriesResolver{r}
}

type standingsResolver struct{ *Resolver }
type standingsQueriesResolver struct{ *Resolver }
