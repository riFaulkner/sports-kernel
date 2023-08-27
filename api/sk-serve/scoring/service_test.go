package scoring

import (
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateLineUp(t *testing.T) {
	var defaultTotalPoints = 0.0
	var qbPosition = model.PlayerPositionQb
	var rbPosition = model.PlayerPositionRb
	var wrPosition = model.PlayerPositionWr
	var tePosition = model.PlayerPositionTe

	var pointVal1 = 1.0
	var pointVal2 = 2.0
	var pointVal3 = 3.0
	var pointVal4 = 4.0
	var pointVal5 = 5.0
	var pointVal6 = 6.0
	var pointVal9 = 9.0
	var pointVal14 = 14.0
	var pointVal20 = 20.0

	lineUp := generateDefaultLineUp()

	type args struct {
		team *MatchUpTeamScoring
	}
	var tests = []struct {
		name string
		args args
		want *MatchUpTeamScoring
	}{
		{name: "Empty Scoring", args: args{team: &MatchUpTeamScoring{}},
			want: &MatchUpTeamScoring{
				TotalPoints: &defaultTotalPoints,
				LineUp:      &lineUp,
			},
		},
		{name: "Team Name persists", args: args{team: &MatchUpTeamScoring{TeamName: "blah"}}, want: &MatchUpTeamScoring{TeamName: "blah",
			TotalPoints: &defaultTotalPoints,
			LineUp:      &lineUp,
		}},
		{name: "IsHomeTeam persists", args: args{team: &MatchUpTeamScoring{IsHomeTeam: true}}, want: &MatchUpTeamScoring{IsHomeTeam: true, TotalPoints: &defaultTotalPoints,
			LineUp: &lineUp,
		}},
		{name: "QBs get sorted", args: args{
			team: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "QB1", NflTeam: "CHI", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal1},
					{Team: "team 1", PlayerName: "QB2", NflTeam: "DAL", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "QB3", NflTeam: "DEN", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal3},
				},
			}},
			want: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "QB3", NflTeam: "DEN", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal3},
					{Team: "team 1", PlayerName: "QB2", NflTeam: "DAL", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "QB1", NflTeam: "CHI", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal1},
				},
				TotalPoints: &pointVal5,
				LineUp: &LineUp{
					Qb: []PlayerScoring{
						{Team: "team 1", PlayerName: "QB3", NflTeam: "DEN", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal3},
					},
					Rb:   []PlayerScoring{},
					Wr:   []PlayerScoring{},
					Te:   []PlayerScoring{},
					Flex: []PlayerScoring{},
					SuperFlex: []PlayerScoring{
						{Team: "team 1", PlayerName: "QB2", NflTeam: "DAL", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal2},
					},
				},
			},
		},
		{name: "RBs get sorted", args: args{
			team: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "RB1", NflTeam: "CHI", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal1},
					{Team: "team 1", PlayerName: "RB2", NflTeam: "DAL", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "RB3", NflTeam: "DEN", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal3},
					{Team: "team 1", PlayerName: "RB4", NflTeam: "ATL", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal4},
					{Team: "team 1", PlayerName: "RB5", NflTeam: "SF", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal5},
				},
			}},
			want: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "RB5", NflTeam: "SF", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal5},
					{Team: "team 1", PlayerName: "RB4", NflTeam: "ATL", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal4},
					{Team: "team 1", PlayerName: "RB3", NflTeam: "DEN", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal3},
					{Team: "team 1", PlayerName: "RB2", NflTeam: "DAL", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "RB1", NflTeam: "CHI", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal1},
				},
				TotalPoints: &pointVal14,
				LineUp: &LineUp{
					Qb: []PlayerScoring{},
					Rb: []PlayerScoring{
						{Team: "team 1", PlayerName: "RB5", NflTeam: "SF", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal5},
					},
					Wr: []PlayerScoring{},
					Te: []PlayerScoring{},
					Flex: []PlayerScoring{
						{Team: "team 1", PlayerName: "RB4", NflTeam: "ATL", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal4},
						{Team: "team 1", PlayerName: "RB3", NflTeam: "DEN", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal3},
					},
					SuperFlex: []PlayerScoring{
						{Team: "team 1", PlayerName: "RB2", NflTeam: "DAL", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal2},
					},
				},
			},
		},
		{name: "WRs get sorted", args: args{
			team: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "WR1", NflTeam: "CHI", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 6.0, Points: &pointVal1},
					{Team: "team 1", PlayerName: "WR2", NflTeam: "DAL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 5.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "WR3", NflTeam: "DEN", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 4.0, Points: &pointVal3},
					{Team: "team 1", PlayerName: "WR4", NflTeam: "ATL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal4},
					{Team: "team 1", PlayerName: "WR5", NflTeam: "SF", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal5},
					{Team: "team 1", PlayerName: "WR6", NflTeam: "NYG", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal6},
				},
			}},
			want: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "WR6", NflTeam: "NYG", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal6},
					{Team: "team 1", PlayerName: "WR5", NflTeam: "SF", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal5},
					{Team: "team 1", PlayerName: "WR4", NflTeam: "ATL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal4},
					{Team: "team 1", PlayerName: "WR3", NflTeam: "DEN", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 4.0, Points: &pointVal3},
					{Team: "team 1", PlayerName: "WR2", NflTeam: "DAL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 5.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "WR1", NflTeam: "CHI", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 6.0, Points: &pointVal1},
				},
				TotalPoints: &pointVal20,
				LineUp: &LineUp{
					Qb: []PlayerScoring{},
					Rb: []PlayerScoring{},
					Wr: []PlayerScoring{
						{Team: "team 1", PlayerName: "WR6", NflTeam: "NYG", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal6},
						{Team: "team 1", PlayerName: "WR5", NflTeam: "SF", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal5},
					},
					Te: []PlayerScoring{},
					Flex: []PlayerScoring{
						{Team: "team 1", PlayerName: "WR4", NflTeam: "ATL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal4},
						{Team: "team 1", PlayerName: "WR3", NflTeam: "DEN", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 4.0, Points: &pointVal3},
					},
					SuperFlex: []PlayerScoring{
						{Team: "team 1", PlayerName: "WR2", NflTeam: "DAL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 5.0, Points: &pointVal2},
					},
				},
			},
		},
		{name: "TEs get sorted", args: args{
			team: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "TE1", NflTeam: "CHI", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 4.0, Points: &pointVal1},
					{Team: "team 1", PlayerName: "TE2", NflTeam: "DAL", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "TE3", NflTeam: "DEN", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal3},
					{Team: "team 1", PlayerName: "TE4", NflTeam: "SEA", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal4},
				},
			}},
			want: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "TE4", NflTeam: "SEA", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal4},
					{Team: "team 1", PlayerName: "TE3", NflTeam: "DEN", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal3},
					{Team: "team 1", PlayerName: "TE2", NflTeam: "DAL", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "TE1", NflTeam: "CHI", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 4.0, Points: &pointVal1},
				},
				TotalPoints: &pointVal9,
				LineUp: &LineUp{
					Qb: []PlayerScoring{},
					Rb: []PlayerScoring{},
					Wr: []PlayerScoring{},
					Te: []PlayerScoring{
						{Team: "team 1", PlayerName: "TE4", NflTeam: "SEA", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal4},
						{Team: "team 1", PlayerName: "TE3", NflTeam: "DEN", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal3},
					},
					Flex: []PlayerScoring{},
					SuperFlex: []PlayerScoring{
						{Team: "team 1", PlayerName: "TE2", NflTeam: "DAL", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal2},
					},
				},
			},
		},
		{name: "Full Roster", args: args{
			team: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "QB1", NflTeam: "CHI", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal1},
					{Team: "team 1", PlayerName: "QB2", NflTeam: "DAL", EligibleSlots: []string{"QB"}, Position: &qbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "RB1", NflTeam: "CHI", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal1},
					{Team: "team 1", PlayerName: "RB2", NflTeam: "DAL", EligibleSlots: []string{"RB"}, Position: &rbPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "WR1", NflTeam: "CHI", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 6.0, Points: &pointVal1},
					{Team: "team 1", PlayerName: "WR2", NflTeam: "DAL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 5.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "WR3", NflTeam: "DEN", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 4.0, Points: &pointVal3},
					{Team: "team 1", PlayerName: "WR4", NflTeam: "ATL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal4},
					{Team: "team 1", PlayerName: "WR5", NflTeam: "SF", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal5},
					{Team: "team 1", PlayerName: "WR6", NflTeam: "NYG", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal6},
					{Team: "team 1", PlayerName: "TE1", NflTeam: "CHI", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 4.0, Points: &pointVal1},
					{Team: "team 1", PlayerName: "TE2", NflTeam: "DAL", EligibleSlots: []string{"TE"}, Position: &tePosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal2},
				},
			}},
			want: &MatchUpTeamScoring{
				Roster: []PlayerScoring{
					{Team: "team 1", PlayerName: "WR6", NflTeam: "NYG", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal6},
					{Team: "team 1", PlayerName: "WR5", NflTeam: "SF", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal5},
					{Team: "team 1", PlayerName: "WR4", NflTeam: "ATL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal4},
					{Team: "team 1", PlayerName: "WR3", NflTeam: "DEN", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 4.0, Points: &pointVal3},
					{Team: "team 1", PlayerName: "WR2", NflTeam: "DAL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: true, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 5.0, Points: &pointVal2},
					{Team: "team 1", PlayerName: "WR1", NflTeam: "CHI", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 6.0, Points: &pointVal1},
				},
				TotalPoints: &pointVal20,
				LineUp: &LineUp{
					Qb: []PlayerScoring{},
					Rb: []PlayerScoring{},
					Wr: []PlayerScoring{
						{Team: "team 1", PlayerName: "WR6", NflTeam: "NYG", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 1.0, Points: &pointVal6},
						{Team: "team 1", PlayerName: "WR5", NflTeam: "SF", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 2.0, Points: &pointVal5},
					},
					Te: []PlayerScoring{},
					Flex: []PlayerScoring{
						{Team: "team 1", PlayerName: "WR4", NflTeam: "ATL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 3.0, Points: &pointVal4},
						{Team: "team 1", PlayerName: "WR3", NflTeam: "DEN", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 4.0, Points: &pointVal3},
					},
					SuperFlex: []PlayerScoring{
						{Team: "team 1", PlayerName: "WR2", NflTeam: "DAL", EligibleSlots: []string{"WR"}, Position: &wrPosition, IsInLineUp: false, InjuryStatus: "ACTIVE", NflOpponent: "GB", GamePlayed: 0, ProjectedPoints: 5.0, Points: &pointVal2},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateLineUp(tt.args.team)

			// Assertions
			assert.Equal(t, tt.want.IsHomeTeam, tt.args.team.IsHomeTeam)
			assert.Equal(t, tt.want.TeamName, tt.args.team.TeamName)
			assert.Equal(t, *tt.want.TotalPoints, *tt.args.team.TotalPoints)
			assert.Equal(t, tt.want.Roster, tt.args.team.Roster)
			assert.Equal(t, *tt.want.LineUp, *tt.args.team.LineUp)
		})
	}
}

func generateDefaultLineUp() LineUp {
	return LineUp{
		Qb:        make([]PlayerScoring, 0, 1),
		Rb:        make([]PlayerScoring, 0, 1),
		Wr:        make([]PlayerScoring, 0, 2),
		Te:        make([]PlayerScoring, 0, 2),
		Flex:      make([]PlayerScoring, 0, 2),
		SuperFlex: make([]PlayerScoring, 0, 1),
	}
}
