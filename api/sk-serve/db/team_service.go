package db

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"log"
	"sort"
	"time"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gFirestore "cloud.google.com/go/firestore"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type TeamImpl struct {
	Client firestore.Client
}

func (u *TeamImpl) AddDeadCapToTeam(ctx context.Context, leagueID string, teamID string, deadCap []*model.DeadCap) bool {
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
	team, ok := u.GetTeamByIdOk(ctx, leagueID, teamID)
	if !ok || team == nil {
		if team == nil {
			gqlerror.Errorf("WARN: Team does not exist, failed update contract")
		}
		return false
	}
	if team.TeamLiabilities == nil || team.TeamLiabilities.DeadCap == nil {
		team.TeamLiabilities = &model.TeamLiabilities{
			DeadCap: make([]*model.DeadCapYear, 0, 0),
		}
	}

	teamDeadCap := team.TeamLiabilities.DeadCap
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
				deadCapYear := &model.DeadCapYear{
					Year:           time.Now().Year() + i,
					DeadCapAccrued: make([]*model.DeadCap, 0, 1),
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

func (u *TeamImpl) GetAllLeagueTeams(ctx context.Context, leagueId string) ([]*model.Team, error) {
	teams := make([]*model.Team, 0)

	//Create Document Ref - There is no traffic associated with this...
	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId)

	results, err := league.Collection(firestore.TeamsCollection).Documents(ctx).GetAll()

	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf("Error fetching teams from league")) //TODO (@kbthree13): This doesn't seem to be sending the error to the client
		return nil, err
	}

	for _, result := range results {
		team := new(model.Team)
		err = result.DataTo(&team)
		team.ID = result.Ref.ID
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (u *TeamImpl) GetTeamById(ctx context.Context, leagueId string, teamId string) (*model.Team, error) {
	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId)

	result, err := league.Collection(firestore.TeamsCollection).Doc(teamId).Get(ctx)

	if err != nil {
		return nil, err
	}
	team := new(model.Team)
	err = result.DataTo(&team)
	team.ID = result.Ref.ID

	if err != nil {
		return nil, err
	}
	return team, nil
}

// Pull out to interface
func (u *TeamImpl) GetTeamByIdOk(ctx context.Context, leagueId string, teamId string) (*model.Team, bool) {
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

	team := new(model.Team)
	err = teamReference.DataTo(&team)
	if err != nil {
		// print out log warning
		log.Printf("WARN: error marshalling team to object: %v", err)
		return nil, false
	}
	team.ID = teamReference.Ref.ID

	return team, true
}

func (u *TeamImpl) Create(ctx context.Context, leagueId string, teamInput model.NewTeam) (*model.Team, error) {
	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId)

	defaultTeamContractsMetadata := generateDefaultTeamContractsMetadata()
	defaultTeamAssets := generateTeamAssets(teamInput.ID)
	defaultTeamLiabilities := &model.TeamLiabilities{}

	team := model.Team{
		ID:                       teamInput.ID,
		TeamName:                 teamInput.TeamName,
		Division:                 teamInput.Division,
		FoundedDate:              time.Now(),
		CurrentContractsMetadata: defaultTeamContractsMetadata,
		TeamAssets:               defaultTeamAssets,
		TeamLiabilities:          defaultTeamLiabilities,
	}

	_, err := league.Collection("teams").Doc(team.ID).Set(ctx, team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (u *TeamImpl) UpdateTeamContractMetaData(ctx context.Context, leagueID string, teamContracts []*contract.Contract) error {
	if teamContracts == nil || len(teamContracts) == 0 {
		return gqlerror.Errorf("Unable to update contract metadata, no team contracts")
	}
	teamID := teamContracts[0].TeamID
	team, err := u.GetTeamById(ctx, leagueID, teamID)
	if err != nil {
		return nil
	}

	// Create default data
	contractsMetadata := make([]*model.ContractsMetadata, league.MaxContractLength, league.MaxContractLength)
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

			var capUtilization *model.CapUtilizationSummary = nil
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
	if team.TeamLiabilities != nil {
		if team.TeamLiabilities.DeadCap != nil {
			// Process dead cap
			for i, deadCapYear := range team.TeamLiabilities.DeadCap {
				deadCapTotal := 0
				totalContracts := 0
				for _, deadCap := range deadCapYear.DeadCapAccrued {
					totalContracts++
					deadCapTotal += deadCap.Amount
				}
				contractsMetadata[i].DeadCap = &model.CapUtilizationSummary{
					CapUtilization: deadCapTotal,
					NumContracts:   totalContracts,
				}
			}
		}
	}

	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueID)

	_, err = league.
		Collection(firestore.TeamsCollection).
		Doc(team.ID).
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

func (u *TeamImpl) GenerateAccessCode(ctx context.Context, leagueId string, teamId string, role string) (string, error) {
	//Get the designated team
	team, err := u.GetTeamById(ctx, leagueId, teamId)

	if err != nil {
		return "Issue creating access string", err
	}

	accessCode := accessCodeFromString(leagueId + teamId + role)

	codes := team.AccessCodes
	codes = append(codes, &accessCode)

	u.Client.
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

func accessCodeFromString(input string) string {
	hashString := []byte(input)
	md5string := md5.Sum(hashString)
	b64String := base64.RawURLEncoding.EncodeToString(md5string[:])
	return b64String
}

func generateDefaultTeamContractsMetadata() *model.ContractsMetadata {
	return &model.ContractsMetadata{
		Year:              time.Now().Year(),
		TotalUtilizedCap:  0,
		TotalAvailableCap: league.SalaryCap,
		QbUtilizedCap: &model.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		RbUtilizedCap: &model.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		WrUtilizedCap: &model.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		TeUtilizedCap: &model.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
		DeadCap: &model.CapUtilizationSummary{
			CapUtilization: 0,
			NumContracts:   0,
		},
	}
}

func generateTeamAssets(teamID string) *model.TeamAssets {
	year := time.Now().Year()
	draftYears := make([]*model.DraftYear, 0, 5)

	for i := 0; i < 5; i++ {
		draftYear := model.DraftYear{
			Year: year + i,
			Picks: []*model.DraftPick{
				{Round: 1, Value: nil, OriginalOwnerID: &teamID},
				{Round: 2, Value: nil, OriginalOwnerID: &teamID},
				{Round: 3, Value: nil, OriginalOwnerID: &teamID},
				{Round: 4, Value: nil, OriginalOwnerID: &teamID},
				{Round: 5, Value: nil, OriginalOwnerID: &teamID},
			},
		}
		draftYears = append(draftYears, &draftYear)
	}

	teamAssets := model.TeamAssets{
		DraftPicks: draftYears,
	}

	return &teamAssets
}
