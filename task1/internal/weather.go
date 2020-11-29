package internal

import (
	"context"
	"time"

	"weather_service/adapters/domain/errors"
	"weather_service/adapters/domain/logger"
	"weather_service/models"
)

const (
	celsius = "celsius"
)

type weatherProvider interface {
	GetCurrentWeather(ctx context.Context, query, units string) (models.WeatherInfo, error)
	GetWeatherForecast(ctx context.Context, query, units string, date time.Time) (models.WeatherInfo, error)
}

type store interface {
	GetWeatherInfo(ctx context.Context, city, date string) (*models.WeatherInfo, error)
	SetWeatherInfo(ctx context.Context, city, date string, info models.WeatherInfo) error
}

type weatherService struct {
	provider weatherProvider
	st       store
}

func NewWeatherService(provider weatherProvider, st store) *weatherService {
	return &weatherService{
		provider: provider,
		st:       st,
	}
}

func (ws *weatherService) GetTemperatureForecast(ctx context.Context, city string, date time.Time) (models.TemperatureInfo, error) {
	if weatherInfo, err := ws.st.GetWeatherInfo(ctx, city, date.String()); err == nil && weatherInfo != nil {
		return weatherInfo.Temperature, nil
	}
	weather, err := ws.provider.GetWeatherForecast(ctx, city, celsius, date)
	if err != nil {
		return models.TemperatureInfo{}, errors.Wrap(err, "fail to get weather forecast")
	}
	if err := ws.st.SetWeatherInfo(ctx, city, date.String(), weather); err != nil {
		logger.WLogger.Err(err).Msgf("fail to set weather info for city: %s", city)
	}

	return weather.Temperature, nil
}

func (ws *weatherService) GetCurrentTemperature(ctx context.Context, city string) (models.TemperatureInfo, error) {
	if weatherInfo, err := ws.st.GetWeatherInfo(ctx, city, models.CurrentWeather); err == nil && weatherInfo != nil {
		return weatherInfo.Temperature, nil
	}
	weather, err := ws.provider.GetCurrentWeather(ctx, city, celsius)
	if err != nil {
		return models.TemperatureInfo{}, errors.Wrap(err, "fail to get current weather")
	}
	if err := ws.st.SetWeatherInfo(ctx, city, models.CurrentWeather, weather); err != nil {
		logger.WLogger.Err(err).Msgf("fail to set weather info for city: %s", city)
	}

	return weather.Temperature, nil
}

func (ws *weatherService) SaveWeatherInfo(ctx context.Context, weather models.WeatherInfo) error {
	if err := ws.st.SetWeatherInfo(ctx, weather.CityName, models.CurrentWeather, weather); err != nil {
		return errors.Wrap(err, "fail to set weather info")
	}

	return nil
}
