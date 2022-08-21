package team

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user/crossfunctional"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type TeamService struct {
	TeamRepository TeamRepository
}

func (t TeamService) AddUserToTeamAndConsumeAccessCode(ctx context.Context, decodedAccessCode crossfunctional.DecodedAccessCode, ownerID string) bool {
	//	Add the user's ID to the TeamOwners array
	//	Remove the Access token used from the ActiveAccessTokens array
	ok := t.TeamRepository.AddUserToTeam(ctx, decodedAccessCode.LeagueID, decodedAccessCode.TeamID, ownerID)
	if ok {
		return t.TeamRepository.RemoveAccessCode(ctx, decodedAccessCode.LeagueID, decodedAccessCode.TeamID, decodedAccessCode.AccessCode)
	}
	return false
}
func (t TeamService) Create(ctx context.Context, leagueId string, team NewTeam) (*Team, error) {
	return t.TeamRepository.Create(ctx, leagueId, team)
}
func (t TeamService) GenerateAccessCode(ctx context.Context, leagueID string, teamID string, role model.Role) (string, error) {
	//Generate a random string, length 5, to append to the pre-encoded string
	randString := randomString(5)
	//Concat data string, and encode in base64
	accessCode := accessCodeFromString(leagueID + "," + teamID + "," + role.String() + "," + randString)

	return accessCode, t.TeamRepository.AddAccessCode(ctx, leagueID, teamID, accessCode)
}
func (t TeamService) GetTeamByOwnerID(ctx context.Context, leagueID string, ownerID string) (*Team, error) {
	team, ok := t.TeamRepository.GetTeamByOwnerID(ctx, leagueID, ownerID)
	if !ok {
		return nil, gqlerror.Errorf("Error occurred getting ownerID: %v teams in league: %v", ownerID, leagueID)
	}
	return team, nil
}
func (t TeamService) GetAllLeagueTeams(ctx context.Context, leagueId string) ([]*Team, error) {
	return t.TeamRepository.GetAllLeagueTeams(ctx, leagueId)
}
func (t TeamService) GetTeamById(ctx context.Context, leagueId string, teamId string) (*Team, error) {
	return t.TeamRepository.GetTeamById(ctx, leagueId, teamId)
}
func (t TeamService) UpdateTeamContractMetaData(ctx context.Context, leagueId string, teamContracts []*contract.Contract) error {
	return t.TeamRepository.UpdateTeamContractMetaData(ctx, leagueId, teamContracts)
}
func (t TeamService) ValidateAccessToken(ctx context.Context, accessCode string) (crossfunctional.DecodedAccessCode, bool) {
	decodedAccessCode := crossfunctional.DecodedAccessCode{
		LeagueID:   "",
		LeagueName: "",
		TeamID:     "",
		Role:       "",
		AccessCode: accessCode,
	}

	rawText, err := decodeAccessCodeString(accessCode)
	if err != nil {
		log.Printf("WARN: could not validate access code: %v", err)
		return decodedAccessCode, false
	}

	leagueID, teamID, role := parseAccessCodeElements(rawText)

	team, err := t.GetTeamById(ctx, leagueID, teamID)
	isInArray, stringIndex := containsString(team.AccessCodes, accessCode)
	if isInArray == false && stringIndex == -1 {
		log.Printf("INFO: Access code not found in document")
		return decodedAccessCode, false
	}

	// Validate that the user isn't already in the league, and get the league name

	decodedAccessCode.LeagueID = leagueID
	decodedAccessCode.TeamID = teamID
	decodedAccessCode.Role = role

	return decodedAccessCode, true
}

func accessCodeFromString(input string) string {
	data := []byte(input)
	b64String := base64.RawURLEncoding.EncodeToString(data[:])
	return b64String
}

func containsString(s []*string, str string) (bool, int) {
	for i, v := range s {
		if *v == str {
			return true, i
		}
	}

	return false, -1
}
func decodeAccessCodeString(accessCode string) (string, error) {
	data, err := base64.RawStdEncoding.DecodeString(accessCode)

	if err != nil {
		log.Printf("WARN: issue decoding the Access Code: %v", err)
		return "Error Decoding", err
	}

	return string(data), nil
}
func parseAccessCodeElements(accessCodeString string) (string, string, string) {
	elements := strings.Split(accessCodeString, ",")
	return elements[0], elements[1], elements[2]
}
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
