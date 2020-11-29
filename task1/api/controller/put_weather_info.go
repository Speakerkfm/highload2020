package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"weather_service/adapters/domain/logger"
	"weather_service/models"
)

func (c *controller) SaveCurrentWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var weatherInfo models.WeatherInfo
	if err := json.NewDecoder(r.Body).Decode(&weatherInfo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Invalid body"}`)
		return
	}
	if err := c.weather.SaveWeatherInfo(r.Context(), weatherInfo); err != nil {
		logger.WLogger.Err(err).Msg("fail to save weather info")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message": "Fail to save weather info"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "Success"}`)
	return
}
