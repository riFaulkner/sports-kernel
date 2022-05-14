package db

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

const collectionLeague = "leagues"

type LeagueImpl struct {
	Client firestore.Client
}

func (u *LeagueImpl) GetAll(ctx context.Context) ([]*model.League, error) {
	leagues := make([]*model.League, 0)

	results, err := u.Client.
		Collection(collectionLeague).
		Documents(ctx).
		GetAll()

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		league := new(model.League)
		err = result.DataTo(&league)
		league.ID = result.Ref.ID
		if err != nil {
			return nil, err
		}
		leagues = append(leagues, league)
	}
	return leagues, nil
}

func (u *LeagueImpl) GetByLeagueId(ctx context.Context, leagueId string) (*model.League, error) {
	result, err := u.Client.
		Collection(collectionLeague).
		Doc(leagueId).
		Get(ctx)
	if err != nil {
		return nil, err
	}

	league := new(model.League)
	err = result.DataTo(&league)
	id := result.Ref.ID
	league.ID = id
	if err != nil {
		return nil, err
	}
	return league, nil
}
