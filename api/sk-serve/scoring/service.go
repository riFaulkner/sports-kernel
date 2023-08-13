package scoring

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"log"
	"sort"
	"time"
)

const (
	weekMatchUpURL      = "https://us-central1-sports-kernel.cloudfunctions.net/getMatchups"
	weekMatchUpAudience = "https://us-central1-sports-kernel.cloudfunctions.net/getMatchups/"
	scoresURL           = "https://us-central1-sports-kernel.cloudfunctions.net/getScores"
	scoresAudience      = "https://us-central1-sports-kernel.cloudfunctions.net/getScores/"
)

type Service struct {
	cache *cache.Cache
}

func NewScoringService(cache *cache.Cache) *Service {
	return &Service{cache: cache}
}

func (s *Service) GetMatchUpScoring(season int, week *int, matchUpNumber int) ([]*MatchUpTeamScoring, error) {
	weekValue := -1
	if week != nil {
		weekValue = *week
	}

	var scoringArray []*MatchUpTeamScoring
	if err := fetchMatchUpScore(season, weekValue, matchUpNumber, &scoringArray); err != nil {
		return nil, fmt.Errorf("scoring.service %v", err)
	}

	for idx := range scoringArray {
		for i := range scoringArray[idx].Roster {
			setPlayerPosition(&scoringArray[idx].Roster[i])
		}
		generateLineUp(scoringArray[idx])
	}

	return scoringArray, nil
}

func (s *Service) GetWeekMatchUps(season int, week *int) ([]*MatchUp, error) {
	weekValue := -1
	if week != nil {
		weekValue = *week
	}

	return s.fetchMatchUpsForSeasonWeek(season, weekValue)
}

func (s *Service) fetchMatchUpsForSeasonWeek(season int, week int) ([]*MatchUp, error) {
	cacheKey := generateWeeklyMatchUpCacheKey(season, week)
	if matchUps, found := s.cache.Get(cacheKey); found {
		return matchUps.([]*MatchUp), nil
	}

	requestBody := bytes.NewReader([]byte(fmt.Sprintf("{\"season\": %d, \"week\":%d}", season, week)))

	var responseBodyAsBuffer bytes.Buffer

	if err := makePostRequest(requestBody, &responseBodyAsBuffer, weekMatchUpURL, weekMatchUpAudience); err != nil {
		log.Printf("makeGetRequest: %v", err)
		return nil, fmt.Errorf("failied to get matchups")
	}

	var matchUps []*MatchUp

	if err := json.Unmarshal(responseBodyAsBuffer.Bytes(), &matchUps); err != nil {
		return nil, fmt.Errorf("scoring.service %v", err)
	}

	if len(matchUps) > 0 {
		s.cache.Set(cacheKey, matchUps, time.Hour)
	}

	return matchUps, nil
}

func fetchMatchUpScore(season int, week int, matchUpNumber int, scoringArray *[]*MatchUpTeamScoring) error {
	reader := bytes.NewReader([]byte(fmt.Sprintf(`{"matchup": %d, "season": %d, "week": %d}`, matchUpNumber, season, week)))
	var b bytes.Buffer

	if err := makePostRequest(reader, &b, scoresURL, scoresAudience); err != nil {
		log.Printf("scoring.Service - makePostRequest: %v", err)
		return fmt.Errorf("Failied to get scores for matchup %d", matchUpNumber)
	}

	return json.Unmarshal(b.Bytes(), &scoringArray)
}

func generateLineUp(team *MatchUpTeamScoring) {
	lineUp := LineUp{
		Qb:        make([]PlayerScoring, 0, 1),
		Rb:        make([]PlayerScoring, 0, 1),
		Wr:        make([]PlayerScoring, 0, 2),
		Te:        make([]PlayerScoring, 0, 2),
		Flex:      make([]PlayerScoring, 0, 2),
		SuperFlex: make([]PlayerScoring, 0, 1),
	}
	sort.Slice(team.Roster, func(i, j int) bool {
		// Check the injury status of the player
		// Push injured or out players to the bottom of the stack
		//var injuryStatus = team.Roster[i].InjuryStatus
		//if getInjuryStatusValue(injuryStatus) >= 4 {
		//
		//}
		if *team.Roster[i].Points == *team.Roster[j].Points {
			return team.Roster[i].ProjectedPoints > team.Roster[j].ProjectedPoints
		}
		return *team.Roster[i].Points > *team.Roster[j].Points
	})

	for i, v := range team.Roster {
		position := *v.Position
		switch position {
		case model.PlayerPositionQb:
			if !checkLineUpSpot(&lineUp.Qb, v) {
				if !checkLineUpSpot(&lineUp.SuperFlex, v) {
					continue
				}
			}
			team.Roster[i].IsInLineUp = true
		case model.PlayerPositionRb:
			if !checkLineUpSpot(&lineUp.Rb, v) {
				if !checkLineUpSpot(&lineUp.Flex, v) {
					if !checkLineUpSpot(&lineUp.SuperFlex, v) {
						continue
					}
				}
			}
			team.Roster[i].IsInLineUp = true
		case model.PlayerPositionTe:
			if !checkLineUpSpot(&lineUp.Te, v) {
				if !checkLineUpSpot(&lineUp.SuperFlex, v) {
					continue
				}
			}
			team.Roster[i].IsInLineUp = true
		case model.PlayerPositionWr:
			if !checkLineUpSpot(&lineUp.Wr, v) {
				if !checkLineUpSpot(&lineUp.Flex, v) {
					if !checkLineUpSpot(&lineUp.SuperFlex, v) {
						continue
					}
				}
			}
			team.Roster[i].IsInLineUp = true
		}
	}

	team.TotalPoints = sumLineUp(team.Roster)

	team.LineUp = &lineUp

}

func checkLineUpSpot(slot *[]PlayerScoring, player PlayerScoring) bool {
	if *player.Points < 0 {
		return false
	}
	if len(*slot) < cap(*slot) {
		*slot = append(*slot, player)
		return true
	}

	return false
}

func getInjuryStatusValue(injuryStatus string) int {
	switch injuryStatus {
	case "ACTIVE":
		return 1
	case "QUESTIONABLE":
		return 2
	case "OUT":
		return 3
	case "INJURY_RESERVE":
		return 4
	default:
		log.Printf("Unknown injury status: %v", injuryStatus)
		return 100
	}
}

func sumLineUp(roster []PlayerScoring) *float64 {
	sum := 0.0
	for _, v := range roster {
		if v.IsInLineUp {
			sum += *v.Points
		}
	}
	return &sum
}

func setPlayerPosition(player *PlayerScoring) {
	for _, v := range player.EligibleSlots {
		if position := model.PlayerPosition(v); position.IsValid() {
			player.Position = &position
			return
		}
	}
}

func generateWeeklyMatchUpCacheKey(season int, week int) string {
	return fmt.Sprintf("weeklyMatchUp-%v-%v", season, week)
}
