package scoring

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type Service struct {
	cache *cache.Cache
}

func NewScoringService(cache *cache.Cache) *Service {
	return &Service{cache: cache}
}

func (s *Service) GetWeekMatchUps(season int, week *int) ([]*MatchUp, error) {
	weekValue := -1
	if week != nil {
		weekValue = *week
	}
	cacheKey := generateWeeklyMatchUpCacheKey(season, weekValue)
	if matchUps, found := s.cache.Get(cacheKey); found {
		return matchUps.([]*MatchUp), nil
	}

	matchUps, err := getMatchUpsViaHttp(season, weekValue)
	if err == nil && len(matchUps) > 0 {
		s.cache.Set(cacheKey, matchUps, time.Hour)
	}
	return matchUps, err
}

func (s *Service) GetMatchUpScoring(season int, week *int, matchUpNumber int) ([]*MatchUpTeamScoring, error) {
	weekValue := -1
	if week != nil {
		weekValue = *week
	}

	requestURL := "https://us-central1-sports-kernel.cloudfunctions.net/getScores"
	audience := "https://us-central1-sports-kernel.cloudfunctions.net/getScores/"

	reader := bytes.NewReader([]byte(fmt.Sprintf(`{"matchup": %d, "season": %d, "week": %d}`, matchUpNumber, season, weekValue)))

	var b bytes.Buffer

	if err := makePostRequest(reader, &b, requestURL, audience); err != nil {
		log.Printf("scoring.Service - makePostRequest: %v", err)
		return nil, fmt.Errorf("Failied to get scores for matchup %d", matchUpNumber)
	}

	var returnValue []*MatchUpTeamScoring

	err := json.Unmarshal(b.Bytes(), &returnValue)

	if err != nil {
		return nil, fmt.Errorf("scoring.service %v", err)
	}

	for idx := range returnValue {
		for i := range returnValue[idx].Roster {
			setPlayerPosition(&returnValue[idx].Roster[i])
		}
		generateLineUp(returnValue[idx])
	}

	return returnValue, nil
}

func getMatchUpsViaHttp(season int, week int) ([]*MatchUp, error) {
	requestURL := "https://us-central1-sports-kernel.cloudfunctions.net/getMatchups"
	audience := "https://us-central1-sports-kernel.cloudfunctions.net/getMatchups/"

	reader := bytes.NewReader([]byte(fmt.Sprintf("{\"season\": %d, \"week\":%d}", season, week)))

	var b bytes.Buffer

	if err := makePostRequest(reader, &b, requestURL, audience); err != nil {
		log.Printf("makeGetRequest: %v", err)
		return nil, fmt.Errorf("failied to get matchups")
	}

	var returnValue []*MatchUp

	if err := json.Unmarshal(b.Bytes(), &returnValue); err != nil {
		return nil, fmt.Errorf("scoring.service %v", err)
	}

	return returnValue, nil
}

func makePostRequest(r io.Reader, w io.Writer, targetURL string, audience string) error {
	var resp *http.Response
	var err error
	if os.Getenv("ENV") != "PROD" {
		request, err := http.NewRequest(http.MethodPost, targetURL, r)
		tokenSource, err := IDTokenTokenSource(context.Background(), audience)
		if err != nil {
			return fmt.Errorf("Error getting token %v", err)
		}
		token, err := tokenSource.Token()

		token.SetAuthHeader(request)
		request.Header.Set("Content-Type", "application/json")

		resp, err = http.DefaultClient.Do(request)
	} else {
		log.Printf("Using production client version")
		ctx := context.Background()
		var client *http.Client
		client, err = idtoken.NewClient(ctx, audience)
		if err != nil {
			return fmt.Errorf("idtoken.NewClient: %v", err)
		}

		resp, err = client.Post(targetURL, "application/json", r)
		log.Printf("Response: %v", resp)
		log.Printf("Response error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("scoring.service - Return status %d", resp.StatusCode)
		return fmt.Errorf("scoring.service - Return status %v", resp.StatusCode)
	}

	if _, err = io.Copy(w, resp.Body); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	return nil
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

func IDTokenTokenSource(ctx context.Context, audience string) (oauth2.TokenSource, error) {
	// First we try the idtoken package, which only works for service accounts
	ts, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		if err.Error() != `idtoken: credential must be service_account, found "authorized_user"` {
			return nil, err
		}
		// If that fails, we use our Application Default Credentials to fetch an id_token on the fly
		gts, err := google.DefaultTokenSource(ctx)
		if err != nil {
			return nil, err
		}
		ts = oauth2.ReuseTokenSource(nil, &idTokenSource{TokenSource: gts})
	}
	return ts, nil
}

// idTokenSource is an oauth2.TokenSource that wraps another
// It takes the id_token from TokenSource and passes that on as a bearer token
type idTokenSource struct {
	TokenSource oauth2.TokenSource
}

func (s *idTokenSource) Token() (*oauth2.Token, error) {
	token, err := s.TokenSource.Token()
	if err != nil {
		return nil, err
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("token did not contain an id_token")
	}

	return &oauth2.Token{
		AccessToken: idToken,
		TokenType:   "Bearer",
		Expiry:      token.Expiry,
	}, nil
}

func generateWeeklyMatchUpCacheKey(season int, week int) string {
	return fmt.Sprintf("weeklyMatchUp-%v-%v", season, week)
}
