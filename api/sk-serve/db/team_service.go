package db

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const collectionTeam = "team"

type TeamImpl struct {
	Client firestore.Client
}

func (u *TeamImpl) GetAll(ctx context.Context, leagueId string) ([]*model.Team, error) {
	teams := make([]*model.Team, 0)

	//Create Document Ref - There is no traffic associated with this...
	league := u.Client.Collection("leagues").Doc(leagueId)

	results, err := league.Collection("teams").Documents(ctx).GetAll()

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
