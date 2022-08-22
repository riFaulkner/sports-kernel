package db

import (
	gfirestore "cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sort"
	"time"
)

type ContractImpl struct {
	Client          firestore.Client
	TeamImpl        TeamRepositoryImpl
	TransactionImpl TransactionImpl
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

func (u *ContractImpl) GetAllActiveTeamContracts(ctx context.Context, leagueID string, teamID string) ([]*contract.Contract, error) {
	contracts := make([]*contract.Contract, 0)

	results, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueID).
		Collection(firestore.PlayerContractsCollection).
		Where("TeamID", "==", teamID).
		Where("ContractStatus", "==", model.ContractStatusActive).
		Documents(ctx).
		GetAll()

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

func (u *ContractImpl) GetById(ctx context.Context, leagueID string, contractID string) (*contract.Contract, bool) {
	contractRef, err := u.Client.Collection(firestore.LeaguesCollection).Doc(leagueID).Collection(firestore.PlayerContractsCollection).Doc(contractID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, true
		}
		// TODO: log error
		return nil, false
	}
	contract := new(contract.Contract)
	err = contractRef.DataTo(&contract)
	if err != nil {
		// TODO:  print to log
		return nil, false
	}
	contract.ID = contractRef.Ref.ID
	return contract, true
}

func (u *ContractImpl) CreateContract(ctx context.Context, leagueId string, contractInput contract.ContractInput) (*contract.Contract, error) {
	u.validateContract(ctx, &leagueId, &contractInput)

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

func (u *ContractImpl) RestructureContract(ctx context.Context, leagueID *string, restructureDetails *contract.ContractRestructureInput) (*contract.Contract, error) {
	// Get the contract
	contractRef, err := u.Client.Collection(firestore.LeaguesCollection).
		Doc(*leagueID).
		Collection(firestore.PlayerContractsCollection).
		Doc(restructureDetails.ContractID).
		Get(ctx)

	if err != nil {
		return nil, err
	}
	contractValue := new(contract.Contract)
	contractRef.DataTo(&contractValue)
	contractValue.ID = contractRef.Ref.ID

	// Validate the contract.

	// contract should be eligible for restructure
	if contractValue.RestructureStatus != model.ContractRestructureStatusEligible {
		return nil, gqlerror.Errorf("Contract %s is ineligible for restructure", restructureDetails.ContractID)
	}
	// contract restructure total should match the original total
	restructureTotal := 0
	guaranteedTotal := 0
	hasChange := false
	for i, value := range restructureDetails.ContractRestructureDetails {
		restructureTotal += value.TotalAmount
		guaranteedTotal += value.GuaranteedAmount
		if contractValue.ContractDetails[i].TotalAmount != value.TotalAmount {
			hasChange = true
		}
	}

	if !hasChange {
		return nil, gqlerror.Errorf("Contract restructure did not have any changes to the contract")
	}

	if contractValue.TotalContractValue != restructureTotal {
		return nil, gqlerror.Errorf("Contract %s restructured contract did not match the value of the original contract", contractValue.ID)
	}
	// contract totals should be 100% guaranteed
	if restructureTotal != guaranteedTotal {
		return nil, gqlerror.Errorf("Contract %s restructure invalid; guaranteed amount not equal to total amount", contractValue.ID)
	}

	// After validation, update the contract

	// Add the old contract details to a metadata field
	contractHistory := contractValue.ContractHistory
	if contractHistory == nil {
		contractHistory = make([]*contract.HistoryRecord, 0, 1)
	}
	mostRecentHistory := contract.HistoryRecord{
		DateUpdated:     time.Now().UnixMilli(),
		ContractDetails: contractValue.ContractDetails,
	}

	contractHistory = append(contractHistory, &mostRecentHistory)

	newContractDetails := make([]*contract.ContractYear, 0, len(restructureDetails.ContractRestructureDetails))
	for _, yearDetailsInput := range restructureDetails.ContractRestructureDetails {
		yearDetails := contract.ContractYear{
			Year:             yearDetailsInput.Year,
			TotalAmount:      yearDetailsInput.TotalAmount,
			PaidAmount:       yearDetailsInput.PaidAmount,
			GuaranteedAmount: yearDetailsInput.GuaranteedAmount,
		}
		newContractDetails = append(newContractDetails, &yearDetails)
	}

	contractValue.ContractDetails = newContractDetails
	contractValue.RestructureStatus = model.ContractRestructureStatusPreviouslyRestructured
	contractValue.ContractHistory = contractHistory

	// Save the new contract
	u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(*leagueID).
		Collection(firestore.PlayerContractsCollection).
		Doc(restructureDetails.ContractID).
		Set(ctx, contractValue)

	// Save the transaction
	inputData, err := json.Marshal(restructureDetails)
	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf("Unable to marshal JSON of restructure object, err: %v", err))
	} else {
		transactionInput := model.TransactionInput{
			TransactionType: model.TransactionTypeContractRestructure,
			TransactionData: string(inputData),
		}
		err = u.TransactionImpl.CreateTransaction(ctx, leagueID, &transactionInput)
		if err != nil {
			return nil, gqlerror.Errorf("Unable to make transaction %v", err)
		}
	}

	// TODO: Recalculate team metadata

	return contractValue, nil
}

func (u *ContractImpl) DropContract(ctx context.Context, leagueID string, teamID string, contractID string) (bool, error) {
	playerContract, ok := u.GetContractById(ctx, leagueID, contractID)

	if !ok {
		return false, gqlerror.Errorf("Error fetching contract, try again later")
	}

	if playerContract == nil {
		return false, gqlerror.Errorf("Contract does not exist")
	}

	// teamID was used for security validate, make sure it's the same team ID that is on the contract
	if playerContract.TeamID != teamID {
		return false, gqlerror.Errorf("TeamId provided did not match the contract's teamID")
	}

	deadCapYears := make([]*team.DeadCap, 0, 2)

	sort.Slice(playerContract.ContractDetails, func(i, j int) bool {
		return playerContract.ContractDetails[i].Year < playerContract.ContractDetails[j].Year
	})
	currentContractYear := playerContract.CurrentYear
	currentContractDetails := playerContract.ContractDetails[(currentContractYear - 1)]
	playerName := u.getPlayerName(ctx, playerContract.PlayerID)

	deadCapYears = append(deadCapYears, &team.DeadCap{
		AssociatedContractID: &playerContract.ID,
		Amount:               calculateDeadCap(currentContractDetails),
		DeadCapNote:          &playerName,
	})

	futureAccumulatedDeadCap := 0
	for _, year := range playerContract.ContractDetails {
		if year.Year > playerContract.CurrentYear {
			futureAccumulatedDeadCap += calculateDeadCap(year)
		}
	}

	deadCapYears = append(deadCapYears, &team.DeadCap{
		AssociatedContractID: &playerContract.ID,
		Amount:               futureAccumulatedDeadCap,
		DeadCapNote:          &playerName,
	})
	// Add dead cap to the team
	ok = u.TeamImpl.AddDeadCapToTeam(ctx, leagueID, teamID, deadCapYears)

	if !ok {
		// consider using a transaction here to roll back
		return false, gqlerror.Errorf("Failed to add dead cap to team")
	}

	// Move the contract to new status
	_, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueID).
		Collection(firestore.PlayerContractsCollection).
		Doc(contractID).
		Update(ctx, []gfirestore.Update{
			{
				Path:  "ContractStatus",
				Value: model.ContractStatusInactiveDropped,
			},
		})
	if err != nil {
		return false, gqlerror.Errorf("Unable to update contract status: %v", err)
	}

	// Save the transaction
	inputData, err := json.Marshal(map[string]interface{}{
		"contractID":   contractID,
		"deadCapAdded": deadCapYears,
	})

	if err != nil {
		graphql.AddError(ctx, gqlerror.Errorf("Unable to marshal JSON of drop object, err: %v", err))
	} else {
		transactionInput := model.TransactionInput{
			TransactionType: model.TransactionTypeDropPlayer,
			TransactionData: string(inputData),
		}
		err = u.TransactionImpl.CreateTransaction(ctx, &leagueID, &transactionInput)
		if err != nil {
			return false, gqlerror.Errorf("Unable to make transaction: %v", err)
		}
	}

	// update team contracts metadata
	teamContracts, err := u.GetAllActiveTeamContracts(ctx, leagueID, teamID)
	if err != nil {
		return true, gqlerror.Errorf("Unable to recalculate team metadata")
	}
	u.TeamImpl.UpdateTeamContractMetaData(ctx, leagueID, teamContracts)

	return true, nil
}

// Pull out to interface
func (u *ContractImpl) GetContractById(ctx context.Context, leagueID string, contractID string) (*contract.Contract, bool) {
	playerContractRef, err := u.Client.
		Collection(firestore.LeaguesCollection).
		Doc(leagueID).
		Collection(firestore.PlayerContractsCollection).
		Doc(contractID).
		Get(ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, true
		}
		// print out log warning
		log.Printf("WARN: error fetching contract: %v", err)
		return nil, false
	}

	playerContract := new(contract.Contract)

	err = playerContractRef.DataTo(&playerContract)
	if err != nil {
		// print out log warning
		log.Printf("WARN: error marshalling contract to object: %v", err)
		return nil, false
	}

	playerContract.ID = playerContractRef.Ref.ID

	return playerContract, true
}

func (u *ContractImpl) validateContract(ctx context.Context, leagueId *string, contractInput *contract.ContractInput) {
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

func getAndValidateContractTotalValue(ctx context.Context, contractYears []*contract.ContractYearInput) *int {
	totalContractValue := 0
	for _, value := range contractYears {
		totalContractValue += value.TotalAmount
	}
	if totalContractValue == 0.0 {
		graphql.AddError(ctx, gqlerror.Errorf("Invalid contract, contract total value is 0"))
	}
	return &totalContractValue
}

// TODO: This is terrible, we should be using the player repository but I didn't want to deal with what I knew would give a circular
// dependency issue and cause a bunch or refactoring. Bad rick.
func (u *ContractImpl) getPlayerName(ctx context.Context, playerID string) string {
	playerRef, err := u.Client.
		Collection(firestore.PlayerCollection).
		Doc(playerID).
		Get(ctx)

	if err != nil {
		log.Printf("Error getting player name from DB")
		return ""
	}
	player := new(model.PlayerNfl)
	playerRef.DataTo(&player)
	return player.PlayerName
}

func calculateDeadCap(contractDetails *contract.ContractYear) int {
	paid := contractDetails.PaidAmount
	guaranteed := contractDetails.GuaranteedAmount

	amountPaidOff := guaranteed - paid

	if amountPaidOff < 0 {
		return 0
	}
	return amountPaidOff
}
