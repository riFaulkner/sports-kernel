package db

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
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
