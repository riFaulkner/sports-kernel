package league

import (
	"time"
)

type League struct {
	ID         string      `json:"id"`
	LeagueName string      `json:"leagueName"`
	LogoURL    string      `json:"logoUrl"`
	StartDate  time.Time   `json:"startDate"`
	Divisions  []*Division `json:"divisions"`
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
