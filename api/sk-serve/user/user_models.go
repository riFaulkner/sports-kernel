package user

import "github.com/rifaulkner/sports-kernel/api/sk-serve/league"

type UserPreferences struct {
	ID                string           `json:"id"`
	OwnerName         string           `json:"ownerName"`
	PreferredLeagueID *string          `json:"preferredLeagueId"`
	IsAdmin           *bool            `json:"isAdmin"`
	Leagues           []*league.League `json:"leagues"`
}

type DecodedAccessCode struct {
	LeagueID   string
	LeagueName string
	TeamID     string
	Role       string
	AccessCode string
}

type UserPreferencesLeagueSnippet struct {
	Id           string `json:"id"`
	LeagueName   string `json:"leagueName"`
	RoleInLeague string `json:"roleInLeague"`
}
