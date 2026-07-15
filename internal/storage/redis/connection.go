package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/config"
	"github.com/redis/go-redis/v9"
)

func Connection(Config *config.Config) error {
	client := redis.NewClient(&redis.Options{
		Addr:     Config.RedisAddr,
		Password: Config.RedisPassword,
		DB:       Config.RedisDB,
	})
	SetClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("не удалось подключиться к Redis: %w", err)
	}

	slog.Info("Redis успешно подключен")
	return nil
}

func Close() error {
	if client == nil {
		return nil
	}

	return client.Close()
}
