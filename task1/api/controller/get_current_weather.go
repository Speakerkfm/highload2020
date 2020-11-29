package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"weather_service/adapters/domain/errors"
	"weather_service/adapters/domain/logger"
	"weather_service/models"
)

func (c *controller) GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	city := r.URL.Query().Get("city")
	if city == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "City is empty"}`)
		return
	}

	temperature, err := c.weather.GetCurrentTemperature(r.Context(), city)
	if err != nil {
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

		logger.WLogger.Err(err).Msg("fail to get current temperature")

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
