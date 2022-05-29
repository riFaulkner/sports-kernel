package transactions

import "github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"

type Transaction struct {
	TransactionType model.TransactionType `json:"transactionType"`
	OccurrenceDate  int64                 `json:"occurrenceDate"`
	TransactionData string                `json:"transactionData"`
}
