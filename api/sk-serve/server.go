package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()
	//client := firestore.NewClient(ctx)
	//graph.Initialize(ctx)

	//userService := &db.UserImpl{Client: client}
	//leagueService := &db.LeagueImpl{Client: client}
	//teamService := &db.TeamImpl{Client: client}
	//contractService := &db.ContractImpl{Client: client}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.Initialize(ctx))) //generated.Config{Resolvers: &graph.Resolver{
	//User: userService, League: leagueService, Team: teamService, Contract: contractService,
	//}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
