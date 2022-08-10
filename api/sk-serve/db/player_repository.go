package db

import (
	"cloud.google.com/go/firestore"
	"context"
	appFirestore "github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/appengine/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
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
			OrderBy("OverallRank", firestore.Asc).
			Documents(ctx).
			GetAll()
	} else {
		results, err = p.Client.
			Collection(appFirestore.PlayerCollection).
			OrderBy("OverallRank", firestore.Asc).
			Limit(*numberOfResults).
			Documents(ctx).
			GetAll()
	}

	if err != nil {
		log.Errorf(ctx, "Error getting documents from firestore")
		return nil, false
	}

	players := transformResultsToPlayers(results, ctx)
	return players, true
}
func (p *PlayerRepositoryImpl) GetPlayersByPosition(ctx context.Context, position model.PlayerPosition) ([]*model.PlayerNfl, bool) {
	results, err := p.Client.
		Collection(appFirestore.PlayerCollection).
		Where("Position", "==", position).
		Documents(ctx).
		GetAll()

	if err != nil {
		log.Errorf(ctx, "Failed getting players by position: %v Error: %v", position, err)
		return nil, false
	}

	players := transformResultsToPlayers(results, ctx)
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

func transformResultsToPlayers(results []*firestore.DocumentSnapshot, ctx context.Context) []*model.PlayerNfl {
	players := make([]*model.PlayerNfl, 0, len(results))

	for _, result := range results {
		player := new(model.PlayerNfl)
		err := result.DataTo(&player)
		player.ID = result.Ref.ID
		if err != nil {
			log.Errorf(ctx, "Error marshaling data object from firestore result playerId: %v", result.Ref.ID)
			continue
		}
		birthday, err := time.Parse("2006-01-02", player.Birthday)

		age := 0
		if err != nil {
			gqlerror.Errorf("Error formatting birthday for player %v")
			//log.Errorf(ctx, "Error formatting birthday for player %v", player.ID)
		} else {
			age = int(time.Now().Sub(birthday).Hours() / 24 / 365)
		}
		player.Age = age

		players = append(players, player)
	}
	return players
}
