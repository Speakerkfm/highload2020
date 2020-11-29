package config

type ConsumerConfig struct {
	LogLevel          string `long:"log-level" env:"MF_LOG_LEVEL" required:"false" default:"info"`
}