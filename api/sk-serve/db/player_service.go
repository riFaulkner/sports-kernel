package db

import (
	"context"
	"crypto/md5"
	"encoding/base64"

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

func (p *PlayerImpl) Create(ctx context.Context, playerInput model.NewPlayerNfl) (*model.PlayerNfl, error) {
	players := p.Client.Collection(collectionPlayers)

	playerId := generatePlayerId(playerInput.PlayerName)

	newPlayer := model.PlayerNfl{
		ID:           playerId,
		PlayerName:   playerInput.PlayerName,
		TeamNfl:      *playerInput.TeamNfl,
		Position:     playerInput.Position,
		PositionRank: *playerInput.PositionRank,
		Birthday:     *playerInput.Birthday,
		OverallRank:  *playerInput.OverallRank,
		Avatar:       "",
	}

	_, err := players.Doc(newPlayer.ID).Set(ctx, newPlayer)
	if err != nil {
		return nil, err
	}

	return &newPlayer, nil
}

func generatePlayerId(name string) string {
	hashString := []byte(name)
	md5string := md5.Sum(hashString)
	b64String := base64.RawURLEncoding.EncodeToString(md5string[:])
	return b64String
}
