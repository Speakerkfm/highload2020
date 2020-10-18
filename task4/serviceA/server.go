package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"serviceA/adapters/domain/config"
	"serviceA/adapters/domain/logger"
	"serviceA/adapters/service_b"
	"serviceA/graph"
	"serviceA/graph/generated"
	"serviceA/internal"
	"serviceA/usecases"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func init() {
	logger.NewLogger()

	// config
	if err := config.InitConfig(); err != nil {
		logger.WLogger.Fatal().Err(err).Msg("fail to init config")
		os.Exit(1)
	}
}

func main() {
	phoneService := service_b.NewPhoneService(fmt.Sprintf("%s:%d", config.WConfig.ServiceBHost, config.WConfig.ServiceBPort))
	codeVerifier := internal.NewCodeVerifier()
	phoneVerifier := usecases.NewPhoneVerifier(phoneService, codeVerifier)
	resolver := graph.NewResolver(phoneVerifier)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", config.WConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.WConfig.Port), nil))
}
