package internal

import (
	"context"
	"fmt"
	"strings"

	"weather_service/adapters/domain/errors"
	"weather_service/models"
)

const (
	weatherTTL = 30 * 60
)

type db interface {
	GetWeatherInfo(ctx context.Context, key string) (*models.WeatherInfo, error)
	Set(ctx context.Context, key string, value interface{}, ttl uint32) error
}

type weatherStore struct {
	db db
}

func NewWeatherStore(db db) *weatherStore {
	return &weatherStore{db: db}
}

func (ws *weatherStore) GetWeatherInfo(ctx context.Context, city, date string) (*models.WeatherInfo, error) {
	info, err := ws.db.GetWeatherInfo(ctx, weatherKey(city, date))
	if err != nil {
		return nil, errors.Wrap(err, "fail to get value from db")
	}
	return info, nil
}

func (ws *weatherStore) SetWeatherInfo(ctx context.Context, city, date string, info models.WeatherInfo) error {
	if err := ws.db.Set(ctx, weatherKey(city, date), info, weatherTTL); err != nil {
		return errors.Wrap(err, "fail to set value from db")
	}

	return nil
}

func weatherKey(city, date string) string {
	return strings.ToLower(fmt.Sprintf("%s:%s", city, date))
}
