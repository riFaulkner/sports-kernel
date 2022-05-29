package transactions

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type Resolver interface {
	CreateTransaction(ctx context.Context, leagueID *string, input *model.TransactionInput) error
}
