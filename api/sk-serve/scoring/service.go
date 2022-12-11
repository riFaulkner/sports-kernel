package scoring

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

func GetWeekMatchUps() ([]*MatchUp, error) {
	requestURL := "https://us-central1-sports-kernel.cloudfunctions.net/getMatchups"
	audience := "https://us-central1-sports-kernel.cloudfunctions.net/getMatchups/"

	reader := bytes.NewReader([]byte(`{}`))
	request, err := http.NewRequest(http.MethodPost, requestURL, reader)
	tokenSource, err := IDTokenTokenSource(context.Background(), audience)
	if err != nil {
		return nil, fmt.Errorf("Error getting token %v", err)
	}
	token, err := tokenSource.Token()

	token.SetAuthHeader(request)

	response, err := http.DefaultClient.Do(request)

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("client: could not read response body: %s\n", err)
	}

	var returnValue []*MatchUp
	//_ = json.Unmarshal([]byte(dataJson), &arr)

	err = json.Unmarshal(resBody, &returnValue)

	if err != nil {
		return nil, fmt.Errorf("scoring.service %v", err)
	}

	fmt.Printf("client: response body: %s\n", resBody)
	return returnValue, nil
}

func GetMatchUpScoring(matchUpNumber int) ([]*MatchUpTeamScoring, error) {
	requestURL := "https://us-central1-sports-kernel.cloudfunctions.net/getScores"
	audience := "https://us-central1-sports-kernel.cloudfunctions.net/getScores/"

	reader := bytes.NewReader([]byte(`{}`))
	request, err := http.NewRequest(http.MethodPost, requestURL, reader)
	query := request.URL.Query()
	query.Add("matchup", strconv.Itoa(matchUpNumber))

	request.URL.RawQuery = query.Encode()

	tokenSource, err := IDTokenTokenSource(context.Background(), audience)
	if err != nil {
		return nil, fmt.Errorf("Error getting token %v", err)
	}
	token, err := tokenSource.Token()

	token.SetAuthHeader(request)

	response, err := http.DefaultClient.Do(request)

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("client: could not read response body: %s\n", err)
	}

	var returnValue []*MatchUpTeamScoring

	err = json.Unmarshal(resBody, &returnValue)

	if err != nil {
		return nil, fmt.Errorf("scoring.service %v", err)
	}

	// Now that we have to objects, add the stuff that actually matters
	// Set the Player positions so it's easy-to-read
	// Set the line-up
	// Remove any player from the line-up to create the bench
	for idx := range returnValue {
		for i := range returnValue[idx].Roster {
			setPlayerPosition(&returnValue[idx].Roster[i])
		}
		generateLineUp(returnValue[idx])
	}

	return returnValue, nil
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
