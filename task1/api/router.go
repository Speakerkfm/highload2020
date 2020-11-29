package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"weather_service/api/controller"
	"weather_service/models"
)

type weatherInformer interface {
	GetTemperatureForecast(ctx context.Context, city string, date time.Time) (models.TemperatureInfo, error)
	GetCurrentTemperature(ctx context.Context, city string) (models.TemperatureInfo, error)
	SaveWeatherInfo(ctx context.Context, weather models.WeatherInfo) error
}

func NewRouter(weather weatherInformer) http.Handler {
	router := chi.NewRouter()
	ctrl := controller.New(weather)

	router.Handle("/metrics", promhttp.Handler())
	router.Route("/v1", func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(middleware.Logger)
		r.Use(MiddlewareMetrics)
		r.Get("/forecast/", ctrl.GetWeatherForecast)
		r.Get("/current/", ctrl.GetCurrentWeather)
		r.Put("/current/", ctrl.SaveCurrentWeather)
	})

	return router
}
