package db

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"sort"
	"time"

	firestoreMain "cloud.google.com/go/firestore"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const teamCollection = "teams"

type TeamImpl struct {
	Client firestore.Client
}

func (u *TeamImpl) AddDeadCapToTeam(ctx context.Context, leagueID string, teamID string, deadCap []*model.DeadCap) (bool, error) {
	team, err := u.GetTeamById(ctx, leagueID, teamID)
	if err != nil {
		return false, err
	}
	teamDeadCap := team.TeamLiabilities.DeadCap
	if len(teamDeadCap) != 0 {
		sort.Slice(teamDeadCap, func(i, j int) bool {
			return teamDeadCap[i].Year < teamDeadCap[j].Year
		})
	}

	for i, value := range deadCap {
		if value.Amount != 0 {
			deadCapYear := &model.DeadCapYear{
				Year:           time.Now().Year() + i,
				DeadCapAccrued: make([]*model.DeadCap, 0, 1),
			}
			if len(teamDeadCap) >= i {
				deadCapYear = teamDeadCap[i]
			}
			deadCapYear.DeadCapAccrued = append(deadCapYear.DeadCapAccrued, value)
		}
	}

	return false, nil
}

func (u *TeamImpl) GetAllLeagueTeams(ctx context.Context, leagueId string) ([]*model.Team, error) {
	teams := make([]*model.Team, 0)

	//Create Document Ref - There is no traffic associated with this...
	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId)

	results, err := league.Collection(teamCollection).Documents(ctx).GetAll()

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

	result, err := league.Collection(teamCollection).Doc(teamId).Get(ctx)

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

func (u *TeamImpl) Create(ctx context.Context, leagueId string, teamInput model.NewTeam) (*model.Team, error) {
	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId)

	defaultTeamContractsMetadata := generateDefaultTeamContractsMetadata()
	defaultTeamAssets := generateTeamAssets(teamInput.ID)

	team := model.Team{
		ID:                       teamInput.ID,
		TeamName:                 teamInput.TeamName,
		Division:                 teamInput.Division,
		FoundedDate:              time.Now(),
		CurrentContractsMetadata: defaultTeamContractsMetadata,
		TeamAssets:               defaultTeamAssets,
	}

	_, err := league.Collection("teams").Doc(team.ID).Set(ctx, team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (u *TeamImpl) UpdateTeamContractMetaData(ctx context.Context, leagueID string, teamContracts []*contract.Contract) error {
	contractsMetadata := make([]*model.ContractsMetadata, 0, 4)
	for i, _ := range contractsMetadata {
		yearMetadata := generateDefaultTeamContractsMetadata()
		yearMetadata.Year = yearMetadata.Year + i
		contractsMetadata[i] = yearMetadata
	}

	currentContractsMetadataDefault := generateDefaultTeamContractsMetadata()

	if len(teamContracts) == 0 {
		return nil
	}

	for _, contract := range teamContracts {
		contractValue := contract.ContractDetails[contract.CurrentYear-1].TotalAmount
		currentContractsMetadataDefault.TotalUtilizedCap += contractValue
		currentContractsMetadataDefault.TotalAvailableCap -= contractValue
		playerType := *contract.PlayerPosition

		switch playerType {
		case "QB":
			currentContractsMetadataDefault.QbUtilizedCap.CapUtilization += contractValue
			currentContractsMetadataDefault.QbUtilizedCap.NumContracts++
		case "RB":
			currentContractsMetadataDefault.RbUtilizedCap.CapUtilization += contractValue
			currentContractsMetadataDefault.RbUtilizedCap.NumContracts++
		case "WR":
			currentContractsMetadataDefault.WrUtilizedCap.CapUtilization += contractValue
			currentContractsMetadataDefault.WrUtilizedCap.NumContracts++
		case "TE":
			currentContractsMetadataDefault.TeUtilizedCap.CapUtilization += contractValue
			currentContractsMetadataDefault.TeUtilizedCap.NumContracts++
		}
	}

	teamID := teamContracts[0].TeamID

	team, err := u.GetTeamById(ctx, leagueID, teamID)
	if err != nil {
		return err
	}

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

	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueID)

	_, err = league.
		Collection(teamCollection).
		Doc(teamID).
		Update(ctx, []firestoreMain.Update{
			{
				Path:  "CurrentContractsMetadata",
				Value: currentContractsMetadataDefault,
			},
		})

	return err
}

func generateDefaultTeamContractsMetadata() *model.ContractsMetadata {
	return &model.ContractsMetadata{
		Year:              time.Now().Year(),
		TotalUtilizedCap:  0,
		TotalAvailableCap: 200000000,
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
	var draftYears []*model.DraftYear

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
