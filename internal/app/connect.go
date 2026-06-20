package app

import (
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/config"
	"github.com/jabrail059/weather-dashboard/internal/storage/redis"
	"github.com/jabrail059/weather-dashboard/internal/storage/sqlite"
)

func Connect(cfg *config.Config) error {
	cityStorage, err := sqlite.New(cfg.SqlitePath)
	if err != nil {
		return fmt.Errorf("Не удалось подключиться к sqlite: %w", err)
	}
	SetStorage(cityStorage)

	if err := redis.Connection(cfg); err != nil {
		return err
	}

	return nil
}
