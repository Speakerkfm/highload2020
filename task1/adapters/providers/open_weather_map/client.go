package open_weather_map

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"weather_service/adapters/domain/errors"
	"weather_service/models"
)

const (
	currentWeatherURL  = "https://%s/weather"
	weatherForecastURL = "https://%s/forecast"
)

type weatherClient struct {
	restyClient        *resty.Client
	openWeatherMapHost string
	apiKey             string
	timeout            time.Duration
	unitsMap           map[string]string
}

func NewWeatherClient(openWeatherMapHost, apiKey string, timeout time.Duration) *weatherClient {
	return &weatherClient{
		restyClient:        resty.New().SetTransport(&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}),
		openWeatherMapHost: openWeatherMapHost,
		apiKey:             apiKey,
		timeout:            timeout,
		unitsMap: map[string]string{
			"celsius": "metric",
		},
	}
}

func (wc *weatherClient) GetCurrentWeather(query, units string) (models.WeatherInfo, error) {
	result := struct {
		Name string `json:"name"`
		Main struct {
			Temp float32 `json:"temp"`
		} `json:"main"`
	}{}

	resp, err := wc.restyClient.R().
		SetHeader("X-RapidAPI-Host", wc.openWeatherMapHost).
		SetHeader("X-RapidAPI-Key", wc.apiKey).
		SetQueryParam("q", query).
		SetQueryParam("units", wc.convertUnits(units)).
		SetResult(&result).
		Get(fmt.Sprintf(currentWeatherURL, wc.openWeatherMapHost))
	if err != nil {
		return models.WeatherInfo{}, errors.Wrap(err, "fail to get current weather")
	}

	if resp.StatusCode() != http.StatusOK {
		if resp.StatusCode() == http.StatusNotFound {
			return models.WeatherInfo{}, errors.WithStack(models.ErrCityNotFound)
		}
		return models.WeatherInfo{}, errors.New(fmt.Sprintf("fail to get current weather with status: %d", resp.StatusCode()))
	}

	return models.WeatherInfo{
		CityName: result.Name,
		Temperature: models.TemperatureInfo{
			Temperature: result.Main.Temp,
			City:        result.Name,
			Unit:        units,
		},
	}, nil
}

func (wc *weatherClient) GetWeatherForecast(query, units string, date time.Time) (models.WeatherInfo, error) {
	now := time.Now()
	if date.After(now.Add(5*24*time.Hour)) || date.Before(now) {
		return models.WeatherInfo{}, errors.WithStack(models.ErrInvalidDate)
	}

	result := struct {
		City struct {
			Name string `json:"name"`
		} `json:"city"`
		List []struct {
			TimeUnix int `json:"dt"`
			Main     struct {
				Temp float32 `json:"temp"`
			} `json:"main"`
		}
	}{}

	resp, err := wc.restyClient.R().
		SetHeader("X-RapidAPI-Host", wc.openWeatherMapHost).
		SetHeader("X-RapidAPI-Key", wc.apiKey).
		SetQueryParam("q", query).
		SetQueryParam("units", wc.convertUnits(units)).
		SetResult(&result).
		Get(fmt.Sprintf(weatherForecastURL, wc.openWeatherMapHost))
	if err != nil {
		return models.WeatherInfo{}, errors.Wrap(err, "fail to get current weather")
	}

	if resp.StatusCode() != http.StatusOK {
		if resp.StatusCode() == http.StatusNotFound {
			return models.WeatherInfo{}, errors.WithStack(models.ErrCityNotFound)
		}
		return models.WeatherInfo{}, errors.New(fmt.Sprintf("fail to get current weather with status: %d", resp.StatusCode()))
	}

	for _, res := range result.List {
		// ищем самое ближайшее время из списка прогноза (прогноз на 5 дней)
		if date.Unix()-int64(res.TimeUnix) < 0 {
			return models.WeatherInfo{
				CityName: result.City.Name,
				Temperature: models.TemperatureInfo{
					Temperature: res.Main.Temp,
					City:        result.City.Name,
					Unit:        units,
				},
			}, nil
		}
	}

	return models.WeatherInfo{}, models.ErrInvalidDate
}

func (wc *weatherClient) convertUnits(units string) string {
	return wc.unitsMap[units]
}
