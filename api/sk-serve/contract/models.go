package contract

import "github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"

type Contract struct {
	ID                  string                          `json:"id"`
	PlayerID            string                          `json:"playerId"`
	TeamID              string                          `json:"teamId"`
	CurrentYear         int                             `json:"currentYear"`
	RestructureStatus   model.ContractRestructureStatus `json:"restructureStatus"`
	TotalContractValue  int                             `json:"totalContractValue"`
	TotalRemainingValue int                             `json:"totalRemainingValue"`
	ContractLength      int                             `json:"contractLength"`
	PlayerPosition      *string                         `json:"playerPosition"`
	ContractDetails     []*model.ContractYear           `json:"contractDetails"`
	ContractHistory     []*HistoryRecord                `json:"contractHistory"`
}

type HistoryRecord struct {
	DateUpdated     int64                 `json:"dateUpdated"`
	ContractDetails []*model.ContractYear `json:"contractDetails"`
}

type ContractInput struct {
	PlayerID            string                          `json:"playerId"`
	TeamID              string                          `json:"teamId"`
	CurrentYear         int                             `json:"currentYear"`
	RestructureStatus   model.ContractRestructureStatus `json:"restructureStatus"`
	TotalContractValue  *int                            `json:"totalContractValue"`
	TotalRemainingValue *int                            `json:"totalRemainingValue"`
	ContractLength      *int                            `json:"contractLength"`
	PlayerPosition      string                          `json:"playerPosition"`
	ContractDetails     []*model.ContractYearInput      `json:"contractDetails"`
}
