package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/scoring"
)

func (r *scoringQueriesResolver) WeekMatchUps(ctx context.Context, obj *scoring.ScoringQueries) ([]*scoring.MatchUp, error) {
	return scoring.GetWeekMatchUps()
}

func (r *scoringQueriesResolver) MatchUpScoring(ctx context.Context, obj *scoring.ScoringQueries, matchUpNumber int) ([]*scoring.MatchUpTeamScoring, error) {
	return scoring.GetMatchUpScoring(matchUpNumber)
}

// ScoringQueries returns generated.ScoringQueriesResolver implementation.
func (r *Resolver) ScoringQueries() generated.ScoringQueriesResolver {
	return &scoringQueriesResolver{r}
}

type scoringQueriesResolver struct{ *Resolver }
