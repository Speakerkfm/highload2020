package config

import (
	"os"
	"strconv"

	"weather_service/adapters/domain/errors"
)

var WConfig Config

type Config struct {
	Port                 int
	OpenWeatherMapHost   string
	OpenWeatherMapApiKey string
}

func InitConfig() error {
	portENV := os.Getenv("WEATHER_PORT")
	openWeatherMapHost := os.Getenv("OPEN_WEATHER_MAP_HOST")
	openWeatherMapApiKey := os.Getenv("OPEN_WEATHER_MAP_API_KEY")

	port, err := strconv.Atoi(portENV)
	if err != nil {
		return errors.Wrap(err, "fail to parse WEATHER_PORT")
	}

	if openWeatherMapHost == "" {
		return errors.New("OPEN_WEATHER_MAP_HOST is empty")
	}

	if openWeatherMapApiKey == "" {
		return errors.New("OPEN_WEATHER_MAP_API_KEY is empty")
	}

	WConfig = Config{
		Port:                 port,
		OpenWeatherMapHost:   openWeatherMapHost,
		OpenWeatherMapApiKey: openWeatherMapApiKey,
	}

	return nil
}
