package db

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/transactions"
	"time"
)

type TransactionImpl struct {
	Client firestore.Client
}

func (u *TransactionImpl) CreateTransaction(ctx context.Context, leagueID *string, input *model.TransactionInput) error {
	transaction := transactions.Transaction{
		TransactionType: input.TransactionType,
		TransactionData: input.TransactionData,
		OccurrenceDate:  time.Now(),
	}

	_, _, err := u.Client.Collection(firestore.LeaguesCollection).
		Doc(*leagueID).
		Collection(firestore.TransactionCollection).
		Add(ctx, transaction)

	return err
}
