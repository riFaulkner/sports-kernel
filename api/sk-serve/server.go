package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
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
	ctx := context.Background()

	srv := configureGql(ctx)

	router := configureRouter(srv)

	startServer(router)
}

func configureGql(context context.Context) *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.Initialize(context)))

	// placeholder, might try to use websockts at somepoint so I don't want to lose this config
	//srv.AddTransport(&transport.Websocket{
	//	Upgrader: websocket.Upgrader{
	//		CheckOrigin: func(r *http.Request) bool {
	//			// Check against desired domains
	//			return r.Host == "sports-kernel.com"
	//		},
	//		ReadBufferSize:  1024,
	//		WriteBufferSize: 1024,
	//	},
	//})

	return srv
}

func configureRouter(server *handler.Server) *chi.Mux {
	router := chi.NewRouter()
	// Setting up cors config
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   getAllowedOrigins(),
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", server)

	return router
}

func getAllowedOrigins() []string {
	if os.Getenv("ENV") == "PROD" {
		return []string{"https://sports-kernel.com", "https://api.sports-kernel.com"}
	}
	return []string{"http://localhost:3000"}
}

func startServer(router *chi.Mux) {
	port := getPort()
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}
