package transactions

import (
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"time"
)

type Transaction struct {
	TransactionType model.TransactionType `json:"transactionType"`
	OccurrenceDate  time.Time             `json:"occurrenceDate"`
	TransactionData string                `json:"transactionData"`
}
