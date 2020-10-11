package main

import (
	"log"
	"net/http"
	"os"

	"serviceA/adapters/service_b"
	"serviceA/graph"
	"serviceA/graph/generated"
	"serviceA/internal"
	"serviceA/usecases"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	phoneService := service_b.NewPhoneService()
	codeVerifier := internal.NewCodeVerifier()
	phoneVerifier := usecases.NewPhoneVerifier(phoneService, codeVerifier)
	resolver := graph.NewResolver(phoneVerifier)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
