package playernfl

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type PlayerService struct {
	PlayerRepository PlayerRepository
}

func (p *PlayerService) GetAllPlayers(ctx context.Context, numberOfResults *int) ([]*model.PlayerNfl, error) {
	players, ok := p.PlayerRepository.GetAll(ctx, numberOfResults)
	if !ok {
		return nil, gqlerror.Errorf("Failed to fetch players")
	}
	return players, nil
}

func (p *PlayerService) GetPlayerById(ctx context.Context, playerId *string) (*model.PlayerNfl, error) {
	player, ok := p.PlayerRepository.GetPlayerById(ctx, playerId)
	if !ok {
		return nil, gqlerror.Errorf("Failed to fetch player with id %v", playerId)
	}
	return player, nil
}

func (p *PlayerService) GetPlayersByPosition(ctx context.Context, position model.PlayerPosition) ([]*model.PlayerNfl, error) {
	players, ok := p.PlayerRepository.GetPlayersByPosition(ctx, position)
	if !ok {
		return nil, gqlerror.Errorf("Unable to fetch players by position: %v", position)
	}
	return players, nil
}

func (p *PlayerService) CreatePlayer(ctx context.Context, playerInput model.NewPlayerNfl) (*model.PlayerNfl, error) {
	newPlayer := convertNewPlayerInputToPlayer(playerInput)

	return p.PlayerRepository.Create(ctx, newPlayer)
}

func generatePlayerId(name string) string {
	hashString := []byte(name)
	md5string := md5.Sum(hashString)
	b64String := base64.RawURLEncoding.EncodeToString(md5string[:])
	return b64String
}

func convertNewPlayerInputToPlayer(newPlayerInput model.NewPlayerNfl) model.PlayerNfl {
	playerId := generatePlayerId(newPlayerInput.PlayerName)

	positionRank := 0
	playerBirthday := ""
	overallRank := 0
	avatarUrl := ""

	if newPlayerInput.PositionRank != nil {
		positionRank = *newPlayerInput.PositionRank
	}
	if newPlayerInput.Birthday != nil {
		playerBirthday = *newPlayerInput.Birthday
	}
	if newPlayerInput.OverallRank != nil {
		overallRank = *newPlayerInput.OverallRank
	}
	if newPlayerInput.Avatar != nil {
		avatarUrl = *newPlayerInput.Avatar
	}

	newPlayer := model.PlayerNfl{
		ID:           playerId,
		PlayerName:   newPlayerInput.PlayerName,
		Team:         newPlayerInput.Team,
		Position:     newPlayerInput.Position,
		PositionRank: positionRank,
		Birthday:     playerBirthday,
		OverallRank:  overallRank,
		Avatar:       avatarUrl,
	}
	return newPlayer
}
