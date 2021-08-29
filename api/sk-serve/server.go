package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/riFaulkner/sports-kernel/sports-kernel/api/sk-serve/db"
	"github.com/riFaulkner/sports-kernel/sports-kernel/api/sk-serve/firestore"
	"github.com/riFaulkner/sports-kernel/sports-kernel/api/sk-serve/graph"
	"github.com/riFaulkner/sports-kernel/sports-kernel/api/sk-serve/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx)

	if err != nil {
		log.Fatalf("Error establishing Firestore Client...")
	}

	userService := &db.UserImpl{Client: client}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		User: userService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
