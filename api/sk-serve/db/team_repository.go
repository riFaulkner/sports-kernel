package db

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gFirestore "cloud.google.com/go/firestore"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type TeamRepositoryImpl struct {
	Client   firestore.Client
	UserImpl UserImpl
}

func (u *TeamRepositoryImpl) AddDeadCapToTeam(ctx context.Context, leagueID string, teamID string, deadCap []*team.DeadCap) bool {
	// Validate the dead cap passed in
	if deadCap == nil {
		log.Printf("Cannot add dead cap to team, invalid deadcap passed")
		return false
	}
	for _, dc := range deadCap {
		if dc == nil {
			log.Printf("Cannot add dead cap to team, invalid deadcap passed")
			return false
		}
	}

	// Get the team
	teamRef, ok := u.GetTeamByIdOk(ctx, leagueID, teamID)
	if !ok || teamRef == nil {
		if teamRef == nil {
			gqlerror.Errorf("WARN: Team does not exist, failed update contract")
		}
		return false
	}
	if teamRef.TeamLiabilities == nil || teamRef.TeamLiabilities.DeadCap == nil {
		teamRef.TeamLiabilities = &team.TeamLiabilities{
			DeadCap: make([]*team.DeadCapYear, 0, 0),
		}
	}

	teamDeadCap := teamRef.TeamLiabilities.DeadCap
	if len(teamDeadCap) != 0 {
		sort.Slice(teamDeadCap, func(i, j int) bool {
			return teamDeadCap[i].Year < teamDeadCap[j].Year
		})
	}

	for i, value := range deadCap {
		if value.Amount != 0 {
			if len(teamDeadCap) > i {
				teamDeadCap[i].DeadCapAccrued = append(teamDeadCap[i].DeadCapAccrued, value)
			} else {
				deadCapYear := &team.DeadCapYear{
					Year:           time.Now().Year() + i,
					DeadCapAccrued: make([]*team.DeadCap, 0, 1),
				}
				deadCapYear.DeadCapAccrued = append(deadCapYear.DeadCapAccrued, value)
				teamDeadCap = append(teamDeadCap, deadCapYear)
			}
		}
	}

	// Save new deadcap to object
	u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueID).
		Collection(firestore.TeamsCollection).
		Doc(teamID).
		Update(ctx, []gFirestore.Update{
			{
				Path:  "TeamLiabilities.DeadCap",
				Value: teamDeadCap,
			},
		})

	return true
}

func (u *TeamRepositoryImpl) GetAllLeagueTeams(ctx context.Context, leagueId string) ([]*team.Team, error) {
	teams := make([]*team.Team, 0)

	//Create Document Ref - There is no traffic associated with this...
	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId)

	results, err := league.Collection(firestore.TeamsCollection).Documents(ctx).GetAll()

	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf("Error fetching teams from league")) //TODO (@kbthree13): This doesn't seem to be sending the error to the client
		return nil, err
	}

	for _, result := range results {
		team := new(team.Team)
		err = result.DataTo(&team)
		team.ID = result.Ref.ID
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (u *TeamRepositoryImpl) GetTeamById(ctx context.Context, leagueId string, teamId string) (*team.Team, error) {
	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId)

	result, err := league.Collection(firestore.TeamsCollection).Doc(teamId).Get(ctx)

	if err != nil {
		return nil, err
	}
	team := new(team.Team)
	err = result.DataTo(&team)
	team.ID = result.Ref.ID

	if err != nil {
		return nil, err
	}
	return team, nil
}

// Pull out to interface
func (u *TeamRepositoryImpl) GetTeamByIdOk(ctx context.Context, leagueId string, teamId string) (*team.Team, bool) {
	teamReference, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueId).Collection(firestore.TeamsCollection).
		Doc(teamId).
		Get(ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, true
		}
		// print out log warning
		log.Printf("WARN: error fetching contract: %v", err)
		return nil, false
	}

	team := new(team.Team)
	err = teamReference.DataTo(&team)
	if err != nil {
		// print out log warning
		log.Printf("WARN: error marshalling team to object: %v", err)
		return nil, false
	}
	team.ID = teamReference.Ref.ID

	return team, true
}

func (u *TeamRepositoryImpl) GetTeamByOwnerID(ctx context.Context, leagueID string, ownerID string) (*team.Team, bool) {
	documents, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueID).
		Collection(firestore.TeamsCollection).
		Where("TeamOwners", "array-contains", ownerID).
		Documents(ctx).
		GetAll()

	teams, ok := processResults(documents, err)

	if !ok || len(teams) == 0 {
		return nil, ok
	}

	if len(teams) > 1 {
		log.Printf("Owner: %v has multiple teams in the same league: %v", ownerID, leagueID)
	}

	return teams[0], true
}

func (u *TeamRepositoryImpl) Create(ctx context.Context, leagueId string, teamInput team.NewTeam) (*team.Team, error) {
	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId)

	defaultTeamContractsMetadata := generateDefaultTeamContractsMetadata()
	defaultTeamAssets := generateTeamAssets(teamInput.ID)
	defaultTeamLiabilities := &team.TeamLiabilities{}

	team := team.Team{
		ID:                       teamInput.ID,
		TeamName:                 teamInput.TeamName,
		Division:                 teamInput.Division,
		FoundedDate:              time.Now(),
		CurrentContractsMetadata: defaultTeamContractsMetadata,
		TeamAssets:               defaultTeamAssets,
		TeamLiabilities:          defaultTeamLiabilities,
		TeamOwners:               make([]string, 0),
	}

	_, err := league.Collection("teams").Doc(team.ID).Set(ctx, team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (u *TeamRepositoryImpl) UpdateTeamContractMetaData(ctx context.Context, leagueID string, teamContracts []*contract.Contract) error {
	if teamContracts == nil || len(teamContracts) == 0 {
		return gqlerror.Errorf("Unable to update contract metadata, no team contracts")
	}
	teamID := teamContracts[0].TeamID
	teamRef, err := u.GetTeamById(ctx, leagueID, teamID)
	if err != nil {
		return nil
	}

	// Create default data
	contractsMetadata := make([]*team.ContractsMetadata, league.MaxContractLength, league.MaxContractLength)
	for i := 0; i < cap(contractsMetadata); i++ {
		yearMetadata := generateDefaultTeamContractsMetadata()
		yearMetadata.Year = yearMetadata.Year + i
		contractsMetadata[i] = yearMetadata
	}

	if len(teamContracts) == 0 {
		return nil
	}

	for _, contract := range teamContracts {
		for _, contractYear := range contract.ContractDetails {
			if contractYear.Year < contract.CurrentYear {
				continue // noop for already completed years
			}

			// returns how far into the future this current year is, starting at 0
			yearsOut := contractYear.Year - contract.CurrentYear
			contractMetadataYear := contractsMetadata[yearsOut]

			contractMetadataYear.TotalUtilizedCap += contractYear.TotalAmount
			contractMetadataYear.TotalAvailableCap -= contractYear.TotalAmount
			playerType := contract.PlayerPosition

			var capUtilization *team.CapUtilizationSummary = nil
			switch playerType {
			case "QB":
				capUtilization = contractMetadataYear.QbUtilizedCap
			case "RB":
				capUtilization = contractMetadataYear.RbUtilizedCap
			case "WR":
				capUtilization = contractMetadataYear.WrUtilizedCap
			case "TE":
				capUtilization = contractMetadataYear.TeUtilizedCap
			}

			capUtilization.CapUtilization += contractYear.TotalAmount
			capUtilization.NumContracts++
		}
	}
	if teamRef.TeamLiabilities != nil {
		if teamRef.TeamLiabilities.DeadCap != nil {
			// Process dead cap
			for i, deadCapYear := range teamRef.TeamLiabilities.DeadCap {
				deadCapTotal := 0
				totalContracts := 0
				for _, deadCap := range deadCapYear.DeadCapAccrued {
					totalContracts++
					deadCapTotal += deadCap.Amount
				}
				contractsMetadata[i].DeadCap = &team.CapUtilizationSummary{
					CapUtilization: deadCapTotal,
					NumContracts:   totalContracts,
				}
			}
		}
	}

	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueID)

	_, err = league.
		Collection(firestore.TeamsCollection).
		Doc(teamRef.ID).
		Update(ctx, []gFirestore.Update{
			{
				Path:  "CurrentContractsMetadata",
				Value: contractsMetadata[0],
			}, {
				Path:  "ContractsMetadata",
				Value: contractsMetadata,
			},
		})

	return err
}

func (u *TeamRepositoryImpl) GenerateAccessCode(ctx context.Context, leagueId string, teamId string, role string) (string, error) {
	//Get the designated team
	teamReference, err := u.GetTeamById(ctx, leagueId, teamId)

	if err != nil {
		return "Issue creating access string", err
	}

	roleString := ""

	if role == "TEAM_OWNER" {
		roleString = "teamOwner"
	} else {
		roleString = "leagueManager"
	}

	//Generate a random string, length 5, to append to the pre-encoded string
	randString := randomString(5)
	//Concat data string, and encode in base64
	accessCode := accessCodeFromString(leagueId + "," + teamId + "," + roleString + "," + randString)

	codes := teamReference.AccessCodes
	codes = append(codes, &accessCode)

	_, err = u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueId).
		Collection(firestore.TeamsCollection).
		Doc(teamId).
		Update(ctx, []gFirestore.Update{
			{
				Path:  "AccessCodes",
				Value: codes,
			},
		})

	return accessCode, nil
}

func (u *TeamRepositoryImpl) AddUserToTeam(ctx context.Context, accessCode string, ownerId string) (string, error) {
	rawText, err := decodeAccessCodeString(accessCode)

	if err != nil {
		log.Printf("WARN: could not validate access code: %v", err)
		return "Error Parsing Code", err
	}
	//0: League ID; 1:Team ID; 2: USER_ROLE; 3: Salt
	rawTextArray := strings.Split(rawText, ",")

	team := new(team.Team)

	//Check for access code
	doc, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(rawTextArray[0]).
		Collection(firestore.TeamsCollection).
		Doc(rawTextArray[1]).Get(ctx)

	cast_err := doc.DataTo(&team)
	if cast_err != nil {
		log.Printf("WARN: Error casting team to object")
		return "Error marshalling object", cast_err
	}

	isInArray, stringIndex := containsString(team.AccessCodes, accessCode)

	if isInArray == false && stringIndex == -1 {
		log.Printf("INFO: Access code not found in document")
		return "Access Code Not Found", nil
	}

	//Add User to Team
	team.TeamOwners = append(team.TeamOwners, ownerId)

	//Remove Access Code
	newCodes := removeElement(team.AccessCodes, stringIndex)
	team.AccessCodes = newCodes

	_, err = u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(rawTextArray[0]).
		Collection(firestore.TeamsCollection).
		Doc(rawTextArray[1]).
		Update(ctx, []gFirestore.Update{
			{
				Path:  "AccessCodes",
				Value: team.AccessCodes,
			},
			{
				Path:  "TeamOwners",
				Value: team.TeamOwners,
			},
		})

	//Update or Create User Preferences
	leagueRef, err := u.Client.Collection(firestore.LeaguesCollection).Doc(rawTextArray[0]).Get(ctx)

	if err != nil {
		return "Error getting league ref", err
	}

	league_obj := new(league.League)

	err = leagueRef.DataTo(&league_obj)

	if err != nil {
		return "Error casting league to object", err
	}

	preferences, err := u.UserImpl.GetUserPreferences(ctx, ownerId)

	if err != nil {
		log.Printf("INFO: No User Preferences found")

		leagues := make([]*league.League, 0)

		newLeague := league.League{
			ID:         leagueRef.Ref.ID,
			LeagueName: league_obj.LeagueName,
		}

		leagues = append(leagues, &newLeague)
		//leagues["ID"] = leagueRef.Ref.ID
		//leagues["LeagueName"] = league_obj.LeagueName

		newUser := model.User{
			ID:        ownerId,
			Avatar:    "",
			OwnerName: "",
			Email:     "",
		}

		err := u.UserImpl.Create(ctx, newUser)

		if err != nil {
			return "Error creating preferences", err
		}

		//Add League to new User
		_, update_err := u.Client.
			Collection(firestore.UsersCollection).
			Doc(ownerId).
			Update(ctx, []gFirestore.Update{
				{
					Path:  "Leagues",
					Value: leagues,
				},
			})

		newRole := model.NewUserRole{
			UserID: ownerId,
			Role:   rawTextArray[2] + ":" + rawTextArray[0],
		}

		u.UserImpl.CreateUserRole(ctx, &newRole)

		if update_err != nil {
			return "Error Updating User Doc", err
		}

		return "NewUser:Success", nil
	}

	newRole := model.NewUserRole{
		UserID: ownerId,
		Role:   rawTextArray[2] + ":" + rawTextArray[0],
	}

	u.UserImpl.CreateUserRole(ctx, &newRole)

	//Add new league to preferences if applicable
	userRef, err := u.Client.Collection(firestore.UsersCollection).Doc(ownerId).Get(ctx)

	if err != nil {
		return "Error retrieving user data", err
	}

	user := new(user.UserPreferences)

	usercast_err := userRef.DataTo(&user)

	if usercast_err != nil {
		return "Error casting user to object", err
	}

	currentLeagues := user.Leagues

	for _, curLeague := range currentLeagues {
		if curLeague.ID == rawTextArray[0] {
			return preferences.ID + ":Success", nil
		}
	}

	newLeague := league.League{
		ID:         leagueRef.Ref.ID,
		LeagueName: league_obj.LeagueName,
	}

	currentLeagues = append(currentLeagues, &newLeague)

	_, update_err := u.Client.
		Collection(firestore.UsersCollection).
		Doc(ownerId).
		Update(ctx, []gFirestore.Update{
			{
				Path:  "Leagues",
				Value: currentLeagues,
			},
		})

	if update_err != nil {
		return "Error updating league list", err
	}

	return preferences.ID + ":Success:Final:" + leagueRef.Ref.ID, nil
}

func decodeAccessCodeString(accessCode string) (string, error) {

	data, err := base64.RawStdEncoding.DecodeString(accessCode)

	if err != nil {
		log.Printf("WARN: issue decoding the Access Code: %v", err)
		return "Error Decoding", err
	}

	return string(data), nil
}

func accessCodeFromString(input string) string {
	data := []byte(input)
	b64String := base64.RawURLEncoding.EncodeToString(data[:])
	return b64String
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func containsString(s []*string, str string) (bool, int) {
	for i, v := range s {
		if *v == str {
			return true, i
		}
	}

	return false, -1
}

func removeElement(s []*string, i int) []*string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func generateDefaultTeamContractsMetadata() *team.ContractsMetadata {
	return &team.ContractsMetadata{
		Year:              time.Now().Year(),
		TotalUtilizedCap:  0,
		TotalAvailableCap: league.SalaryCap,
		QbUtilizedCap: &team.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		RbUtilizedCap: &team.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		WrUtilizedCap: &team.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		TeUtilizedCap: &team.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		DeadCap: &team.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
	}
}

func generateTeamAssets(teamID string) *team.TeamAssets {
	year := time.Now().Year()
	draftYears := make([]*team.DraftYear, 0, 5)

	for i := 0; i < 5; i++ {
		draftYear := team.DraftYear{
			Year: year + i,
			Picks: []*team.DraftPick{
				{Round: 1, Value: nil, OriginalOwnerID: &teamID},
				{Round: 2, Value: nil, OriginalOwnerID: &teamID},
				{Round: 3, Value: nil, OriginalOwnerID: &teamID},
				{Round: 4, Value: nil, OriginalOwnerID: &teamID},
				{Round: 5, Value: nil, OriginalOwnerID: &teamID},
			},
		}
		draftYears = append(draftYears, &draftYear)
	}

	teamAssets := team.TeamAssets{
		DraftPicks: draftYears,
	}

	return &teamAssets
}

func processResults(teamsReference []*gFirestore.DocumentSnapshot, err error) ([]*team.Team, bool) {
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, true
		}
		// print out log warning
		log.Printf("WARN: error fetching contract: %v", err)
		return nil, false
	}

	teams := make([]*team.Team, len(teamsReference), len(teamsReference))

	for i, teamReference := range teamsReference {
		team := new(team.Team)

		err = teamReference.DataTo(&team)
		if err != nil {
			// print out log warning
			log.Printf("WARN: error marshalling team to object: %v", err)
			return nil, false
		}
		team.ID = teamReference.Ref.ID
		teams[i] = team
	}

	return teams, true
}
