package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

func Connection() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	SetClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("Не удалось подключиться к Redis")
	}

	slog.Info("Redis успешно подключен")
	return nil
}
