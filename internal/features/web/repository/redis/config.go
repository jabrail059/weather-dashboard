package web_redis_repository

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr     string `envconfig:"ADDR" required:"true"`
	Password string `envconfig:"PASSWORD"`
	DB       int    `envconfig:"DB" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("REDIS", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get Redis config: %w", err)
		panic(err)
	}

	return config
}
