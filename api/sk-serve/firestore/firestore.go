package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Client interface {
	Collection(path string) *firestore.CollectionRef
}

// Use a service account
func NewClient(ctx context.Context) (Client, error) {
	//TODO: Replace the path below with the new location of the secret? Or change
	sa := option.WithCredentialsFile("path/to/creds")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return app.Firestore(ctx)

}
