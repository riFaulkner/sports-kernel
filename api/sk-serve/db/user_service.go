package db

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user"
	"google.golang.org/api/iterator"
)

const collectionName = "users"

type UserImpl struct {
	Client firestore.Client
}

func (u *UserImpl) GetAll(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 0)

	results, err := u.Client.
		Collection(collectionName).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		user := new(model.User)
		err = result.DataTo(&user)
		id := result.Ref.ID
		if id == "" {
			//err = Error("Document does not have an ID")
			return nil, err
		}
		user.ID = id
		if err != nil {
			return nil, err
		}
		user.ID = result.Ref.ID
		users = append(users, user)
	}
	return users, nil
}

func (u *UserImpl) Create(ctx context.Context, user model.User) error {
	_, _, err := u.Client.Collection(collectionName).Add(ctx, user)
	return err
}

func (u *UserImpl) GetUserPreferences(ctx context.Context, userId string) (*user.UserPreferences, error) {
	result, err := u.Client.Collection(collectionName).Doc(userId).Get(ctx)
	if err != nil {
		return nil, err
	}

	userPreferences := new(user.UserPreferences)
	err = result.DataTo(&userPreferences)
	id := result.Ref.ID
	userPreferences.ID = id
	if err != nil {
		return nil, err
	}

	return userPreferences, nil
}

func (u *UserImpl) CreateUserRole(ctx context.Context, newRole *model.NewUserRole) (*model.UserRoles, error) {
	user := u.Client.Collection(firestore.UsersCollection).Doc(newRole.UserID)

	response, _, err := user.Collection(firestore.UserRolesCollection).Add(ctx, newRole)
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, err
	}

	role := model.UserRoles{
		ID:     response.ID,
		UserID: newRole.UserID,
		Role:   newRole.Role,
	}

	return &role, nil
}

func (u *UserImpl) GetUserRoles(cxt context.Context, userID *string) ([]*model.UserRoles, error) {
	league := u.Client.Collection(firestore.UsersCollection).Doc(*userID)

	iter := league.Collection(firestore.UserRolesCollection).
		Documents(cxt)

	userLeagueRoles := make([]*model.UserRoles, 0)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		role := new(model.UserRoles)
		err = doc.DataTo(role)
		if err != nil {
			graphql.AddError(cxt, err)
			continue
		}
		role.ID = doc.Ref.ID
		userLeagueRoles = append(userLeagueRoles, role)
	}

	return userLeagueRoles, nil
}
