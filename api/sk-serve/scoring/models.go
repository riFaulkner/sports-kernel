package scoring

import "github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"

type ScoringQueries struct {
}

type MatchUp struct {
	AwayTeam      string `json:"away_team"`
	HomeTeam      string `json:"home_team"`
	MatchUpNumber int    `json:"match_up_number"`
}

type PlayerScoring struct {
	Team            string                `json:"team_name"`
	PlayerName      string                `json:"player_name"`
	NflTeam         string                `json:"nfl_team"`
	ProjectedPoints float64               `json:"projected_points"`
	Points          *float64              `json:"points"`
	EligibleSlots   []string              `json:"eligible_pos"`
	Position        *model.PlayerPosition `json:"position"`
	IsInLineUp      bool                  `json:"is_in_line_up"`
	InjuryStatus    string                `json:"injury_status"`
	NflOpponent     string                `json:"nfl_opponent"`
	GamePlayed      int                   `json:"game_played"`
}

type MatchUpTeamScoring struct {
	IsHomeTeam  bool            `json:"is_home_team"`
	Roster      []PlayerScoring `json:"roster"`
	TeamName    string          `json:"team_name"`
	TotalPoints *float64        `json:"total_points"`
	LineUp      *LineUp         `json:"line_up"`
}

type LineUp struct {
	Qb        []PlayerScoring `json:"qb"`
	Rb        []PlayerScoring `json:"rb"`
	Flex      []PlayerScoring `json:"flex"`
	Wr        []PlayerScoring `json:"wr"`
	Te        []PlayerScoring `json:"te"`
	SuperFlex []PlayerScoring `json:"superFlex"`
}
