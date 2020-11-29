package aerospike

import (
	"context"
	"strconv"
	"strings"

	"github.com/aerospike/aerospike-client-go"

	"weather_service/adapters/domain/errors"
	"weather_service/models"
)

const (
	serviceName = "weather"
)

type client struct {
	*aerospike.Client
}

func New(dsn string) (*client, error) {
	hostPorts := strings.Split(dsn, ",")
	var aeroHosts []*aerospike.Host
	for _, h := range hostPorts {
		hostPort := strings.Split(h, ":")
		host := hostPort[0]
		port, err := strconv.Atoi(hostPort[1])
		if err != nil {
			return nil, err
		}

		aeroHosts = append(aeroHosts, aerospike.NewHost(host, port))
	}

	c, err := aerospike.NewClientWithPolicyAndHost(nil, aeroHosts...)
	if err != nil {
		return nil, errors.Wrap(err, "fail to connect to aerospike")
	}

	return &client{Client: c}, nil
}

func (c *client) Set(ctx context.Context, key string, value interface{}, ttl uint32) error {
	if ctx.Err() == context.Canceled {
		return errors.New("ctx done")
	}

	aeroKey, err := aerospike.NewKey(serviceName, models.WeatherInfoKey, key)
	if err != nil {
		return errors.Wrap(err, "fail to create key")
	}

	writePolicy := aerospike.NewWritePolicy(0, ttl)
	if err := c.Client.PutObject(writePolicy, aeroKey, value); err != nil {
		return errors.Wrap(err, "fail to put object")
	}

	return nil
}

func (c *client) GetWeatherInfo(ctx context.Context, key string) (*models.WeatherInfo, error) {
	if ctx.Err() == context.Canceled {
		return nil, errors.New("ctx done")
	}

	aeroKey, err := aerospike.NewKey(serviceName, models.WeatherInfoKey, key)
	if err != nil {
		return nil, errors.Wrap(err, "fail to create key")
	}

	info := models.WeatherInfo{}
	if err := c.Client.GetObject(nil, aeroKey, &info); err != nil {
		return nil, errors.Wrap(err, "fail to put object")
	}

	return &info, nil
}
