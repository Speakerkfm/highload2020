package usecases

import (
	"time"

	"weather_service/adapters/domain/errors"
	"weather_service/models"
)

const (
	celsius = "celsius"
)

type WeatherInformer interface {
	GetCurrentWeather(query, units string) (models.WeatherInfo, error)
	GetWeatherForecast(query, units string, date time.Time) (models.WeatherInfo, error)
}

type weatherService struct {
	informer WeatherInformer
}

func NewWeatherService(informer WeatherInformer) *weatherService {
	return &weatherService{
		informer: informer,
	}
}

func (ws *weatherService) GetTemperatureForecast(city string, date time.Time) (models.TemperatureInfo, error) {
	weather, err := ws.informer.GetWeatherForecast(city, celsius, date)
	if err != nil {
		return models.TemperatureInfo{}, errors.Wrap(err, "fail to get weather forecast")
	}

	return weather.Temperature, nil
}

func (ws *weatherService) GetCurrentTemperature(city string) (models.TemperatureInfo, error) {
	weather, err := ws.informer.GetCurrentWeather(city, celsius)
	if err != nil {
		return models.TemperatureInfo{}, errors.Wrap(err, "fail to get current weather")
	}

	return weather.Temperature, nil
}
