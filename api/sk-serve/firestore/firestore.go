package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

const (
	LeaguesCollection         = "leagues"
	PlayerCollection          = "playersNFL"
	PlayerContractsCollection = "playerContracts"
	TransactionCollection     = "transactions"
	UserRolesCollection       = "user-roles"
	UsersCollection           = "users"
)

type Client interface {
	Collection(path string) *firestore.CollectionRef
}

func NewClient(ctx context.Context) *firestore.Client {
	projectID := "sports-kernel"
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client %v", err)
	}
	return client
}
