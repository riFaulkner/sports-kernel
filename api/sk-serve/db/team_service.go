package db

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"log"
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

func (u *TeamImpl) GetAll(ctx context.Context, leagueId string) ([]*model.Team, error) {
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

func (u *TeamImpl) UpdateTeamContractMetaData(ctx context.Context, leagueId string, teamContracts []*contract.Contract) error {
	currentContractsMetadataDefault := model.ContractsMetadata{
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
	}

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

	teamId := teamContracts[0].TeamID

	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId)

	_, err := league.Collection(teamCollection).Doc(teamId).Update(ctx, []firestoreMain.Update{
		{
			Path:  "CurrentContractsMetadata",
			Value: currentContractsMetadataDefault,
		},
	})

	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

func generateDefaultTeamContractsMetadata() *model.ContractsMetadata {
	return &model.ContractsMetadata{
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
