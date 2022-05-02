package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/auth"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rs/cors"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/generated"
)

const defaultPort = "8080"

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	ctx := context.Background()

	firestoreClient := firestore.NewClient(ctx)

	srv := configureGql(firestoreClient)

	router := configureRouter(srv, firestoreClient)

	startServer(router)
}

func configureGql(client firestore.Client) *handler.Server {
	graphConfig := graph.Initialize(client)
	graphConfig.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (res interface{}, err error) {
		if !auth.GetUserRolesFromContext(ctx).ContainsRole(role, ctx) {
			// block calling the next resolver
			return nil, fmt.Errorf("Access denied")
		}

		// or let it pass through
		return next(ctx)
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graphConfig))

	return srv
}

func configureRouter(server *handler.Server, client firestore.Client) *chi.Mux {
	router := chi.NewRouter()

	// Setting up cors config
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   getAllowedOrigins(),
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", middleware.CheckJWT(auth.LoadUserRoles(client, server)))

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
