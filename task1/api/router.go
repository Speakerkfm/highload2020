package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"weather_service/api/controller"
	"weather_service/models"
)

type WeatherService interface {
	GetTemperatureForecast(city string, date time.Time) (models.TemperatureInfo, error)
	GetCurrentTemperature(city string) (models.TemperatureInfo, error)
}

func NewRouter(weather WeatherService) http.Handler {
	router := chi.NewRouter()
	ctrl := controller.New(weather)

	router.Route("/v1", func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Get("/forecast/", ctrl.GetWeatherForecast)
		r.Get("/current/", ctrl.GetCurrentWeather)
	})

	return router
}
