package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/scoring"
)

func (r *scoringQueriesResolver) WeekMatchUps(ctx context.Context, obj *scoring.ScoringQueries, season int, week *int) ([]*scoring.MatchUp, error) {
	return r.ScoringService.GetWeekMatchUps(season, week)
}

func (r *scoringQueriesResolver) MatchUpScoring(ctx context.Context, obj *scoring.ScoringQueries, season int, week *int, matchUpNumber int) ([]*scoring.MatchUpTeamScoring, error) {
	return r.ScoringService.GetMatchUpScoring(season, week, matchUpNumber)
}

// ScoringQueries returns generated.ScoringQueriesResolver implementation.
func (r *Resolver) ScoringQueries() generated.ScoringQueriesResolver {
	return &scoringQueriesResolver{r}
}

type scoringQueriesResolver struct{ *Resolver }
