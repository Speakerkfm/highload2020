package usecases

import (
	"context"
	"time"

	"weather_service/adapters/domain/errors"
	"weather_service/models"
)

type weatherService interface {
	GetCurrentTemperature(ctx context.Context, city string) (models.TemperatureInfo, error)
	GetTemperatureForecast(ctx context.Context, city string, date time.Time) (models.TemperatureInfo, error)
	SaveWeatherInfo(ctx context.Context, weather models.WeatherInfo) error
}

type weatherInformer struct {
	svc weatherService
}

func NewWeatherInformer(svc weatherService) *weatherInformer {
	return &weatherInformer{
		svc: svc,
	}
}

func (ws *weatherInformer) GetTemperatureForecast(ctx context.Context, city string, date time.Time) (models.TemperatureInfo, error) {
	weather, err := ws.svc.GetTemperatureForecast(ctx, city, date)
	if err != nil {
		return models.TemperatureInfo{}, errors.Wrap(err, "fail to get weather forecast")
	}

	return weather, nil
}

func (ws *weatherInformer) GetCurrentTemperature(ctx context.Context, city string) (models.TemperatureInfo, error) {
	weather, err := ws.svc.GetCurrentTemperature(ctx, city)
	if err != nil {
		return models.TemperatureInfo{}, errors.Wrap(err, "fail to get current weather")
	}

	return weather, nil
}

func (ws *weatherInformer) SaveWeatherInfo(ctx context.Context, weather models.WeatherInfo) error {
	if err := ws.svc.SaveWeatherInfo(ctx, weather); err != nil {
		return errors.Wrap(err, "fail to save current weather")
	}

	return nil
}
