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

	results, err := league.Collection("playerContracts").Where("teamID", "==", teamID).Documents(ctx).GetAll()

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
		contracts = append(contracts, contract)
	}
	return contracts, nil
}
