package contract

import "github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"

type Contract struct {
	ID                  string                          `json:"id"`
	PlayerID            string                          `json:"playerId"`
	TeamID              string                          `json:"teamId"`
	CurrentYear         int                             `json:"currentYear"`
	RestructureStatus   model.ContractRestructureStatus `json:"restructureStatus"`
	TotalContractValue  float64                         `json:"totalContractValue"`
	TotalRemainingValue float64                         `json:"totalRemainingValue"`
	ContractLength      int                             `json:"contractLength"`
	PlayerPosition      *string                         `json:"playerPosition"`
	ContractDetails     []*model.ContractYear           `json:"contractDetails"`
	Player              *model.PlayerNfl                `json:"player"`
}
