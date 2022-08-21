package db

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gFirestore "cloud.google.com/go/firestore"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	accessCodesPath = "AccessCodes"
)

type TeamRepositoryImpl struct {
	Client firestore.Client
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
	defaultTeamScoring := generateDefaultTeamScoring()

	team := team.Team{
		ID:                       teamInput.ID,
		TeamName:                 teamInput.TeamName,
		Division:                 teamInput.Division,
		FoundedDate:              time.Now(),
		CurrentContractsMetadata: defaultTeamContractsMetadata,
		TeamAssets:               defaultTeamAssets,
		TeamLiabilities:          defaultTeamLiabilities,
		TeamOwners:               make([]string, 0),
		TeamScoring:              defaultTeamScoring,
	}

	_, err := league.Collection(firestore.TeamsCollection).Doc(team.ID).Set(ctx, team)
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

func (u *TeamRepositoryImpl) AddAccessCode(ctx context.Context, leagueId string, teamId string, accessCode string) error {
	_, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueId).
		Collection(firestore.TeamsCollection).
		Doc(teamId).
		Update(ctx, []gFirestore.Update{
			{
				Path:  accessCodesPath,
				Value: gFirestore.ArrayUnion(accessCode),
			},
		})

	return err
}

func (u *TeamRepositoryImpl) AddUserToTeam(ctx context.Context, leagueID string, teamID string, ownerID string) bool {
	_, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueID).
		Collection(firestore.TeamsCollection).
		Doc(teamID).Update(ctx, []gFirestore.Update{
		{
			Path:  "TeamOwners",
			Value: gFirestore.ArrayUnion(ownerID)},
	})
	if err != nil {
		log.Printf("error adding user to team")
		return false
	}
	return true
}

func (u *TeamRepositoryImpl) RemoveAccessCode(ctx context.Context, leagueID string, teamID string, accessCode string) bool {
	_, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueID).
		Collection(firestore.TeamsCollection).
		Doc(teamID).Update(ctx, []gFirestore.Update{
		{
			Path:  accessCodesPath,
			Value: gFirestore.ArrayRemove(accessCode),
		},
	})
	if err != nil {
		log.Printf("Error removing access code for user")
		return false
	}
	return true
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

func generateDefaultTeamScoring() []team.TeamScoring {
	return []team.TeamScoring{
		{
			Year: time.Now().Year(),
			Summary: team.TeamScoringSeasonSummary{
				Wins:          0,
				Losses:        0,
				Ties:          0,
				CurrentStreak: 0,
			},
			Weeks: make([]team.TeamScoringWeek, 0),
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
