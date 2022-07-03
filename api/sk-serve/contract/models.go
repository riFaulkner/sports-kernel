package contract

import "github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"

type Contract struct {
	ID                  string                          `json:"id"`
	PlayerID            string                          `json:"playerId"`
	TeamID              string                          `json:"teamId"`
	CurrentYear         int                             `json:"currentYear"`
	RestructureStatus   model.ContractRestructureStatus `json:"restructureStatus"`
	ContractStatus      model.ContractStatus            `json:"contractStatus"`
	TotalContractValue  int                             `json:"totalContractValue"`
	TotalRemainingValue int                             `json:"totalRemainingValue"`
	ContractLength      int                             `json:"contractLength"`
	PlayerPosition      *string                         `json:"playerPosition"`
	ContractDetails     []*ContractYear                 `json:"contractDetails"`
	ContractHistory     []*HistoryRecord                `json:"contractHistory"`
}

type HistoryRecord struct {
	DateUpdated     int64           `json:"dateUpdated"`
	ContractDetails []*ContractYear `json:"contractDetails"`
}

type ContractInput struct {
	PlayerID            string                           `json:"playerId"`
	TeamID              string                           `json:"teamId"`
	CurrentYear         int                              `json:"currentYear"`
	TotalContractValue  *int                             `json:"totalContractValue"`
	TotalRemainingValue *int                             `json:"totalRemainingValue"`
	ContractLength      *int                             `json:"contractLength"`
	PlayerPosition      string                           `json:"playerPosition"`
	ContractDetails     []*ContractYearInput             `json:"contractDetails"`
	ContractStatus      *model.ContractStatus            `json:"contractStatus"`
	RestructureStatus   *model.ContractRestructureStatus `json:"contractRestructureStatus"`
}

type ContractRestructureInput struct {
	ContractID                 string               `json:"contractId"`
	ContractRestructureDetails []*ContractYearInput `json:"contractRestructureDetails"`
}

type ContractYear struct {
	Year             int `json:"year"`
	TotalAmount      int `json:"totalAmount"`
	PaidAmount       int `json:"paidAmount"`
	GuaranteedAmount int `json:"guaranteedAmount"`
}

type ContractYearInput struct {
	Year             int `json:"year"`
	TotalAmount      int `json:"totalAmount"`
	PaidAmount       int `json:"paidAmount"`
	GuaranteedAmount int `json:"guaranteedAmount"`
}
