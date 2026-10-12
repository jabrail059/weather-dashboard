package core_config

import (
	"fmt"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TimeZone         *time.Location
	GeocodingBaseURL string `envconfig:"GEOCODING_BASE_URL" required:"true"`
	ForecastBaseURL  string `envconfig:"FORECAST_BASE_URL" required:"true"`
	IpApiBaseURL     string `envconfig:"IP_API_BASE_URL" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config

	tz := os.Getenv("TIME_ZONE")
	if tz == "" {
		tz = "UTC"
	}

	if err := envconfig.Process("", &config); err != nil {
		return nil, fmt.Errorf("process envconfig: %w", err)
	}

	zone, err := time.LoadLocation(tz)
	if err != nil {
		return nil, fmt.Errorf("load time zone: %s: %w", tz, err)
	}

	config.TimeZone = zone

	return &config, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get core config: %w", err)
		panic(err)
	}

	return config
}
