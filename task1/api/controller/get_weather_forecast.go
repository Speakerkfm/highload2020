package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"weather_service/adapters/domain/errors"
	"weather_service/adapters/domain/logger"
	"weather_service/models"
)

func (c *controller) GetWeatherForecast(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "City is empty"}`)
		return
	}

	dateParam := r.URL.Query().Get("timestamp")
	date, err := time.Parse("2006-01-02T15:04", dateParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Timestamp is invalid"}`)
		return
	}

	temperature, err := c.weather.GetTemperatureForecast(city, date)
	if err != nil {
		logger.WLogger.Err(err).Msg("fail to get temperature forecast")

		if errors.Is(err, models.ErrInvalidDate) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"message": "Date is invalid"}`)
			return
		}

		if errors.Is(err, models.ErrCityNotFound) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"message": "City not found"}`)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Internal server error"}`)
		return
	}

	response := struct {
		City        string  `json:"city"`
		Unit        string  `json:"unit"`
		Temperature float32 `json:"temperature"`
	}{
		City:        temperature.City,
		Unit:        temperature.Unit,
		Temperature: temperature.Temperature,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.WLogger.Err(err).Msg("fail to marshal response")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Internal server error"}`)
		return
	}

	return
}
