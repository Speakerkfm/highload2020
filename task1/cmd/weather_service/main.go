package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"weather_service/adapters/domain/aerospike"
	"weather_service/adapters/domain/config"
	"weather_service/adapters/domain/logger"
	"weather_service/adapters/providers/open_weather_map"
	"weather_service/api"
	"weather_service/internal"
	"weather_service/usecases"
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
	openWeatherMapClient := open_weather_map.NewWeatherClient(config.WConfig.OpenWeatherMapHost, config.WConfig.OpenWeatherMapApiKey, 5*time.Second)

	aeroStore, err := aerospike.New(config.WConfig.AerospikeDSN)
	if err != nil {
		logger.WLogger.Fatal().Err(err).Msg("fail to connect to aerospike")
	}

	weatherStore := internal.NewWeatherStore(aeroStore)

	weatherService := internal.NewWeatherService(openWeatherMapClient, weatherStore)

	weatherInformer := usecases.NewWeatherInformer(weatherService)

	router := api.NewRouter(weatherInformer)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.WConfig.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.WLogger.Fatal().Err(err).Msgf("listen: %s\n", err)
		}
	}()
	log.Println("server started listening on port:", config.WConfig.Port)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.WLogger.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.WLogger.Fatal().Err(err).Msg("server forced to shutdown:")
		os.Exit(1)
	}

	logger.WLogger.Info().Msg("server exiting")
}
