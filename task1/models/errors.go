package models

import (
	"weather_service/adapters/domain/errors"
)

var (
	ErrInvalidDate  = errors.New("date is invalid")
	ErrCityNotFound = errors.New("city not found")
)
