package db

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	appFirestore "github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

const collectionPlayers = "playersNFL"

type PlayerImpl struct {
	Client appFirestore.Client
}

func (p *PlayerImpl) GetAll(ctx context.Context, numberOfResults *int) ([]*model.PlayerNfl, error) {
	players := make([]*model.PlayerNfl, 0)

	results, err := p.Client.
		Collection(collectionPlayers).OrderBy("overallRank", firestore.Asc).Limit(*numberOfResults).Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		player := new(model.PlayerNfl)
		err = result.DataTo(&player)
		player.ID = result.Ref.ID
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

func (p *PlayerImpl) GetPlayerById(ctx context.Context, playerId *string) (*model.PlayerNfl, error) {
	result, err := p.Client.Collection(collectionPlayers).Doc(*playerId).Get(ctx)
	if err != nil {
		return nil, err
	}
	player := new(model.PlayerNfl)
	err = result.DataTo(&player)
	id := result.Ref.ID
	player.ID = id
	if err != nil {
		return nil, err
	}

	return player, nil
}
