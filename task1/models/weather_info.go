package models

const (
	WeatherInfoKey = "info"

	CurrentWeather = "current"
)

type WeatherInfo struct {
	CityName    string          `json:"city_name"`
	Temperature TemperatureInfo `json:"temperature"`
}
