package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jabrail059/weather-dashboard/storage"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RedisConnection() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	storage.SetRedisClient(client)

	if _, err := client.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("Не удалось подключиться к Redis")
	}

	slog.Info("Redis успешно подключен")
	return nil
}
