package standings

type StandingsQueries struct {
}

type Standings struct {
	TeamName      string  `json:"team_name"`
	DivisionId    int     `json:"division_id"`
	DivisionName  string  `json:"division_name"`
	TeamWins      int     `json:"team_wins"`
	TeamLosses    int     `json:"team_losses"`
	TeamTies      int     `json:"team_ties"`
	PointsFor     float32 `json:"points_for"`
	PointsAgaints float32 `json:"points_against"`
}
