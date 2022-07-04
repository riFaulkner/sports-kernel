package db

import (
	"cloud.google.com/go/firestore"
	"context"
	appFirestore "github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"google.golang.org/appengine/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PlayerRepositoryImpl struct {
	Client appFirestore.Client
}

func (p *PlayerRepositoryImpl) Create(ctx context.Context, player model.PlayerNfl) (*model.PlayerNfl, error) {
	players := p.Client.Collection(appFirestore.PlayerCollection)

	_, err := players.Doc(player.ID).Set(ctx, player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (p *PlayerRepositoryImpl) GetAll(ctx context.Context, numberOfResults *int) ([]*model.PlayerNfl, bool) {
	players := make([]*model.PlayerNfl, 0)

	var results []*firestore.DocumentSnapshot
	var err error

	if *numberOfResults < -1 {
		log.Errorf(ctx, "Invalid number of results passed, less than -1")
		return nil, false
	}

	//-1 to return all players, no query limit
	if *numberOfResults == -1 {
		results, err = p.Client.
			Collection(appFirestore.PlayerCollection).
			OrderBy("overallRank", firestore.Asc).
			Documents(ctx).
			GetAll()
	} else {
		results, err = p.Client.
			Collection(appFirestore.PlayerCollection).
			OrderBy("overallRank", firestore.Asc).
			Limit(*numberOfResults).
			Documents(ctx).
			GetAll()
	}

	if err != nil {
		log.Errorf(ctx, "Error getting documents from firestore")
		return nil, false
	}

	for _, result := range results {
		player := new(model.PlayerNfl)
		err = result.DataTo(&player)
		player.ID = result.Ref.ID
		if err != nil {
			log.Errorf(ctx, "Error marshaling data object from firestore result")
			return nil, false
		}
		players = append(players, player)
	}
	return players, true
}

func (p *PlayerRepositoryImpl) GetPlayerById(ctx context.Context, playerId *string) (*model.PlayerNfl, bool) {
	result, err := p.Client.
		Collection(appFirestore.PlayerCollection).
		Doc(*playerId).
		Get(ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, true
		}

		log.Errorf(ctx, "Failed to get player from firestore")
		return nil, false
	}

	player := new(model.PlayerNfl)
	err = result.DataTo(&player)
	player.ID = result.Ref.ID
	if err != nil {
		log.Errorf(ctx, "Error marshaling data object from firestore")
		return nil, false
	}

	return player, true
}
