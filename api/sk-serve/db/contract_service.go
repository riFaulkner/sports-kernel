package db

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type ContractImpl struct {
	Client firestore.Client
}

func (u *ContractImpl) GetAll(ctx context.Context, leagueID string, teamID string) ([]*model.Contract, error) {
	contracts := make([]*model.Contract, 0)

	//Create Document Ref - There is no traffic associated with this...
	league := u.Client.Collection("leagues").Doc(leagueID)

	results, err := league.Collection("playerContracts").Where("TeamID", "==", teamID).Documents(ctx).GetAll()

	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf("Error fetching teams from league")) //TODO (@kbthree13): This doesn't seem to be sending the error to the client
		return nil, err
	}

	for _, result := range results {
		contract := new(model.Contract)
		err = result.DataTo(&contract)
		if err != nil {
			return nil, err
		}
		contract.ID = result.Ref.ID
		contracts = append(contracts, contract)
	}
	return contracts, nil
}

func (u *ContractImpl) GetContractByLeagueAndPlayerId(ctx context.Context, leagueId string, playerId string) (*model.Contract, error) {
	result, err := u.Client.Collection("leagues").Doc(leagueId).Collection("playerContracts").Doc(playerId).Get(ctx)
	if err != nil {
		return nil, err
	}
	contract := new(model.Contract)
	err = result.DataTo(contract)
	if err != nil {
		return nil, err
	}
	contract.ID = result.Ref.ID
	return contract, nil
}

func (u *ContractImpl) CreateContract(ctx context.Context, leagueId string, contractInput *model.ContractInput) (*model.Contract, error) {
	u.validateContract(ctx, &leagueId, contractInput)

	if len(graphql.GetErrors(ctx)) > 0 {
		return nil, graphql.GetErrors(ctx)
	}

	playerContractsCollection := u.Client.Collection("leagues").Doc(leagueId).Collection("playerContracts")
	add, _, err := playerContractsCollection.Add(ctx, contractInput)

	if err != nil {
		return nil, err
	}

	// No error creating new team, update the team contract metadata
	doc, err := add.Get(ctx)

	if err != nil {
		return nil, err
	}

	contract := new(model.Contract)
	err = doc.DataTo(&contract)
	if err != nil {
		return nil, err
	}
	contract.ID = doc.Ref.ID

	return contract, nil
}

func (u *ContractImpl) validateContract(ctx context.Context, leagueId *string, contractInput *model.ContractInput) {
	contractInput.TotalContractValue = getAndValidateContractTotalValue(ctx, contractInput.ContractDetails)

	u.validatePlayer(ctx, leagueId, &contractInput.PlayerID)

	if len(contractInput.ContractDetails) > 4 {
		graphql.AddError(ctx, gqlerror.Errorf("ContractResolver length is too long"))
	}
}

func (u *ContractImpl) validatePlayer(ctx context.Context, leagueId *string, playerId *string) {
	// TODO:  Validate that the player record exists

	// TODO: Validate player does NOT have a current valid contract
	//returnVal, err := u.GetContractByLeagueAndPlayerId(ctx, *leagueId, *playerId)
}

func (u *ContractImpl) validateTeam(ctx context.Context, leagueId *string, teamId *string) {
	// valid team ID in that league
	// TODO: Validate that this wont push team over cap value
	//result, _ := u.GetAll(ctx, *leagueId, *teamId)
}

func getAndValidateContractTotalValue(ctx context.Context, contractYears []*model.ContractYearInput) *float64 {
	totalContractValue := 0.0
	for _, value := range contractYears {
		totalContractValue += value.TotalAmount
	}
	if totalContractValue == 0.0 {
		graphql.AddError(ctx, gqlerror.Errorf("Invalid contract, contract total value is 0"))
	}
	return &totalContractValue
}
