package controller

import (
	"context"
	"time"

	"weather_service/models"
)

type weatherInformer interface {
	GetTemperatureForecast(ctx context.Context, city string, date time.Time) (models.TemperatureInfo, error)
	GetCurrentTemperature(ctx context.Context, city string) (models.TemperatureInfo, error)
	SaveWeatherInfo(ctx context.Context, weather models.WeatherInfo) error
}

type controller struct {
	weather weatherInformer
}

func New(weather weatherInformer) *controller {
	return &controller{
		weather: weather,
	}
}
