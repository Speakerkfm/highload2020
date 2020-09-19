package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var WLogger *WeatherLogger

type WeatherLogger struct {
	zerolog.Logger
}

func NewLogger() {
	WLogger = &WeatherLogger{
		Logger: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}
