package league

import (
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"time"
)

type LeagueQueries struct {
	League    *League              `json:"league"`
	Contracts []*contract.Contract `json:"contracts"`
}

type LeagueMutations struct {
	MakeTradeProposal *bool `json:"makeTradeProposal"`
}

type League struct {
	ID            string      `json:"id"`
	CurrentSeason int         `json:"currentSeason"`
	LeagueName    string      `json:"leagueName"`
	LogoURL       string      `json:"logoUrl"`
	StartDate     time.Time   `json:"startDate"`
	Divisions     []*Division `json:"divisions"`
}

type Division struct {
	DivisionName string `json:"divisionName"`
	LeadingWins  *int   `json:"leadingWins"`
}

type NewLeagueInput struct {
	LeagueName string     `json:"leagueName"`
	LogoUrl    *string    `json:"logoUrl"`
	StartDate  *time.Time `json:"startDate"`
	Divisions  []string   `json:"divisions"`
}
