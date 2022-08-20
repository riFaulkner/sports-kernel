package user

type UserPreferences struct {
	ID                string                          `json:"id"`
	OwnerName         string                          `json:"ownerName"`
	PreferredLeagueID *string                         `json:"preferredLeagueId"`
	IsAdmin           *bool                           `json:"isAdmin"`
	Leagues           []*UserPreferencesLeagueSnippet `json:"leagues"`
}

type UserPreferencesLeagueSnippet struct {
	Id           string `json:"id"`
	LeagueName   string `json:"leagueName"`
	RoleInLeague string `json:"roleInLeague"`
}
