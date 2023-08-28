package playernfl

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/db"
	appFirestore "github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"time"
)

type PlayerRepositoryCacheImpl struct {
	dbPlayerRepository db.PlayerRepositoryImpl
	cache              *cache.Cache
}

const playerCacheKeyPrefix = "playerID"

func NewCachedImpl(client appFirestore.Client, cache *cache.Cache) PlayerRepository {
	var playerRepository PlayerRepository = &PlayerRepositoryCacheImpl{
		dbPlayerRepository: db.PlayerRepositoryImpl{client},
		cache:              cache,
	}
	return playerRepository
}

func (r *PlayerRepositoryCacheImpl) Create(ctx context.Context, player model.PlayerNfl) (*model.PlayerNfl, error) {
	return r.dbPlayerRepository.Create(ctx, player)
}

func (r *PlayerRepositoryCacheImpl) GetAll(ctx context.Context) ([]*model.PlayerNfl, bool) {
	return r.dbPlayerRepository.GetAll(ctx)
}

func (r *PlayerRepositoryCacheImpl) GetPlayersWithLimit(ctx context.Context, numberOfResults int) ([]*model.PlayerNfl, bool) {
	return r.dbPlayerRepository.GetPlayersWithLimit(ctx, numberOfResults)
}

func (r *PlayerRepositoryCacheImpl) GetPlayersByPosition(ctx context.Context, position model.PlayerPosition) ([]*model.PlayerNfl, bool) {
	return r.dbPlayerRepository.GetPlayersByPosition(ctx, position)
}

func (r *PlayerRepositoryCacheImpl) GetPlayerById(ctx context.Context, playerId *string) (*model.PlayerNfl, bool) {
	if playerValue, found := r.cache.Get(getPlayerCacheKey(*playerId)); found {
		player := playerValue.(model.PlayerNfl)
		return &player, true
	}
	if player, ok := r.dbPlayerRepository.GetPlayerById(ctx, playerId); ok {
		r.cache.Set(getPlayerCacheKey(player.ID), *player, time.Hour)
		return player, true
	}

	return nil, false
}

func getPlayerCacheKey(playerId string) string {
	return fmt.Sprintf("%s:%s==", playerCacheKeyPrefix, playerId)
}
