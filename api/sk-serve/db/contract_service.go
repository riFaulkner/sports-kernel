package db

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sort"
)

type ContractImpl struct {
	Client         firestore.Client
	PlayerResolver PlayerImpl
}

func (u *ContractImpl) GetAllLeagueContracts(ctx context.Context, leagueID string) ([]*contract.Contract, error) {
	contracts := make([]*contract.Contract, 0, 20)
	contractsReturned, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueID).
		Collection(firestore.PlayerContractsCollection).
		Documents(ctx).
		GetAll()
	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf("Error fetching contracts for league"))
		return nil, err
	}

	for _, result := range contractsReturned {
		contract := new(contract.Contract)
		err = result.DataTo(&contract)
		if err != nil {
			return nil, err
		}
		contract.ID = result.Ref.ID
		contracts = append(contracts, contract)
	}
	return contracts, nil
}

func (u *ContractImpl) GetAllTeamContracts(ctx context.Context, leagueID string, teamID string) ([]*contract.Contract, error) {
	contracts := make([]*contract.Contract, 0)

	//Create Document Ref - There is no traffic associated with this...
	league := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueID)

	results, err := league.Collection(firestore.PlayerContractsCollection).Where("TeamID", "==", teamID).Documents(ctx).GetAll()

	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf("Error fetching teams from league")) //TODO (@kbthree13): This doesn't seem to be sending the error to the client
		return nil, err
	}

	for _, result := range results {
		contract := new(contract.Contract)
		err = result.DataTo(&contract)
		if err != nil {
			return nil, err
		}
		contract.ID = result.Ref.ID
		contracts = append(contracts, contract)
	}

	sort.SliceStable(contracts, func(i, j int) bool {
		return contracts[i].TotalContractValue > contracts[j].TotalContractValue
	})

	return contracts, nil
}

func (u *ContractImpl) GetContractByLeagueAndPlayerId(ctx context.Context, leagueId string, playerId string) (*contract.Contract, error) {
	// Todo add a filter for active filters
	result, err := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueId).Collection(firestore.PlayerContractsCollection).Doc(playerId).Get(ctx)
	if err != nil {
		return nil, err
	}
	contract := new(contract.Contract)
	err = result.DataTo(contract)
	if err != nil {
		return nil, err
	}
	contract.ID = result.Ref.ID
	return contract, nil
}

func (u *ContractImpl) CreateContract(ctx context.Context, leagueId string, contractInput *model.ContractInput) (*contract.Contract, error) {
	u.validateContract(ctx, &leagueId, contractInput)

	if len(graphql.GetErrors(ctx)) > 0 {
		return nil, graphql.GetErrors(ctx)
	}

	u.Client.Collection(firestore.LeaguesCollection)

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

	contract := new(contract.Contract)
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
		graphql.AddError(ctx, gqlerror.Errorf("Contract length is too long"))
	}
}

func (u *ContractImpl) validatePlayer(ctx context.Context, leagueId *string, playerId *string) (*string, error) {
	_, err := u.GetContractByLeagueAndPlayerId(ctx, *leagueId, *playerId)
	// Firesstore returns an error when the record does not exist. So we want to make sure there
	// was an error to validate that player does not have a contract
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
	}

	return nil, errors.New("Failed to create contract, player already has an active contract")
	// if there is a value, return an error
}

func (u *ContractImpl) validateTeam(ctx context.Context, leagueId *string, teamId *string) {
	// valid team ID in that league
	// TODO: Validate that this wont push team over cap value
	//result, _ := u.GetAllTeamContracts(ctx, *leagueId, *teamId)
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
