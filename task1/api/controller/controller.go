package controller

import (
	"time"

	"weather_service/models"
)

type WeatherService interface {
	GetTemperatureForecast(city string, date time.Time) (models.TemperatureInfo, error)
	GetCurrentTemperature(city string) (models.TemperatureInfo, error)
}

type controller struct {
	weather WeatherService
}

func New(weather WeatherService) *controller {
	return &controller{
		weather: weather,
	}
}
