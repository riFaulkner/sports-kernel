package web_utilities

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

func GetLeagueIDFromContext(ctx context.Context) (string, bool) {
	leagueID := graphql.GetOperationContext(ctx).Variables["leagueId"]
	str, ok := leagueID.(string)
	return str, ok
}
