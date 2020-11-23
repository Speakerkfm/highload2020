package aerospike_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"weather_service/adapters/domain/aerospike"
	"weather_service/models"
)

func TestNew(t *testing.T) {
	client, err := aerospike.New("35.228.110.108", 3000)
	assert.Nil(t, err)

	err = client.SetWeatherInfo(context.Background(), "test", models.WeatherInfo{CityName: "Perm"})
	assert.Nil(t, err)

	info, err  := client.GetWeatherInfo(context.Background(), "test")
	assert.Nil(t, err)
	fmt.Println(info)
}
