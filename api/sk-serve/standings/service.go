package standings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/patrickmn/go-cache"
)

const (
	standingsURL      = "https://us-central1-sports-kernel.cloudfunctions.net/getStandings"
	standingsAudience = "https://us-central1-sports-kernel.cloudfunctions.net/getStandings/"
)

type Service struct {
	cache *cache.Cache
}

func NewStandingsService(cache *cache.Cache) *Service {
	return &Service{cache: cache}
}

func (s *Service) GetStandings(season int, week *int) ([]*Standings, error) {
	weekValue := -1
	if week != nil {
		weekValue = *week
	}

	var standingsArray []*Standings
	if err := fetchStandings(season, weekValue, &standingsArray); err != nil {
		return nil, fmt.Errorf("standings.service %v", err)
	}

	return standingsArray, nil
}

func fetchStandings(season int, week int, standingsArray *[]*Standings) error {
	reader := bytes.NewReader([]byte(fmt.Sprintf(`{"season": %d, "week": %d}`, season, week)))
	var b bytes.Buffer

	if err := makePostRequest(reader, &b, standingsURL, standingsAudience); err != nil {
		log.Printf("standings.Service - makePostRequest: %v", err)
		return fmt.Errorf("Failied to get standings")
	}

	return json.Unmarshal(b.Bytes(), &standingsArray)
}
