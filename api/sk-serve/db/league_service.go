package db

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"time"
)

const collectionLeague = "leagues"

type LeagueImpl struct {
	Client firestore.Client
}

func (u *LeagueImpl) CreateLeague(ctx context.Context, input league.NewLeagueInput) (*league.League, error) {
	newLeague := league.League{
		LeagueName:    input.LeagueName,
		CurrentSeason: time.Now().Year(),
		LogoURL:       getLogo(input.LogoUrl),
		StartDate:     getTime(input.StartDate),
		Divisions:     getDivisions(input.Divisions),
	}

	_, _, err := u.Client.Collection(firestore.LeaguesCollection).Add(ctx, newLeague)
	if err != nil {
		return nil, err
	}
	return &newLeague, err
}

func (u *LeagueImpl) GetAll(ctx context.Context) ([]*league.League, error) {
	leagues := make([]*league.League, 0)

	results, err := u.Client.
		Collection(collectionLeague).
		Documents(ctx).
		GetAll()

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		league := new(league.League)
		err = result.DataTo(&league)
		league.ID = result.Ref.ID
		if err != nil {
			return nil, err
		}
		leagues = append(leagues, league)
	}
	return leagues, nil
}

func (u *LeagueImpl) GetByLeagueId(ctx context.Context, leagueId string) (*league.League, error) {
	result, err := u.Client.
		Collection(collectionLeague).
		Doc(leagueId).
		Get(ctx)
	if err != nil {
		return nil, err
	}

	league := new(league.League)
	err = result.DataTo(&league)
	id := result.Ref.ID
	league.ID = id
	if err != nil {
		return nil, err
	}
	return league, nil
}

func getDivisions(inputDivisions []string) []*league.Division {
	zeroWins := 0

	divisions := []*league.Division{
		{DivisionName: "East", LeadingWins: &zeroWins},
		{DivisionName: "West", LeadingWins: &zeroWins},
	}

	if len(inputDivisions) > 0 {
		divisions[0].DivisionName = inputDivisions[0]
		if len(inputDivisions) > 1 {
			divisions[1].DivisionName = inputDivisions[1]
		}
	}
	return divisions
}

func getLogo(inputLogo *string) string {
	if inputLogo != nil {
		return *inputLogo
	}
	return ""
}

func getTime(inputTime *time.Time) time.Time {
	if inputTime != nil {
		return *inputTime
	}
	return time.Now()
}
