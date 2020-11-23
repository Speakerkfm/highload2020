package aerospike

import (
	"context"

	"github.com/aerospike/aerospike-client-go"

	"weather_service/adapters/domain/errors"
	"weather_service/models"
)

const (
	serviceName = "test"
)

type client struct {
	*aerospike.Client
}

func New(host string, port int) (*client, error) {
	c, err := aerospike.NewClient(host, port)
	if err != nil {
		return nil, errors.Wrap(err, "fail to connect to aerospike")
	}

	return &client{Client: c}, nil
}

func (c *client) SetWeatherInfo(ctx context.Context, key string, info models.WeatherInfo) error {
	if ctx.Err() == context.Canceled {
		return errors.New("ctx done")
	}

	aeroKey, err := aerospike.NewKey(serviceName, "aerospike", key)
	if err != nil {
		return errors.Wrap(err, "fail to create key")
	}

	if err := c.Client.PutObject(nil, aeroKey, info); err != nil {
		return errors.Wrap(err, "fail to put object")
	}

	return nil
}

func (c *client) GetWeatherInfo(ctx context.Context, key string) (models.WeatherInfo, error) {
	if ctx.Err() == context.Canceled {
		return models.WeatherInfo{}, errors.New("ctx done")
	}

	aeroKey, err := aerospike.NewKey(serviceName, "aerospike", key)
	if err != nil {
		return models.WeatherInfo{}, errors.Wrap(err, "fail to create key")
	}

	info := models.WeatherInfo{}
	if err := c.Client.GetObject(nil, aeroKey, &info); err != nil {
		return models.WeatherInfo{}, errors.Wrap(err, "fail to put object")
	}

	return info, nil
}
