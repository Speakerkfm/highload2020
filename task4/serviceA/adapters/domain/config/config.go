package config

import (
	"os"
	"strconv"

	"serviceA/adapters/domain/errors"
)

var WConfig Config

type Config struct {
	Port         int
	ServiceBHost string
	ServiceBPort int
}

func InitConfig() error {
	portENV := os.Getenv("SERVICEA_PORT")
	serviceBHost := os.Getenv("SERVICEB_HOST")
	serviceBPortENV := os.Getenv("SERVICEB_PORT")

	port, err := strconv.Atoi(portENV)
	if err != nil {
		return errors.Wrap(err, "fail to parse SERVICEA_PORT")
	}

	serviceBPort, err := strconv.Atoi(serviceBPortENV)
	if err != nil {
		return errors.Wrap(err, "fail to parse SERVICEB_PORT")
	}

	if serviceBHost == "" {
		return errors.New("SERVICEB_HOST is empty")
	}

	WConfig = Config{
		Port:         port,
		ServiceBHost: serviceBHost,
		ServiceBPort: serviceBPort,
	}

	return nil
}
