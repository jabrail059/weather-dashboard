package service

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/storage/redis"
)

func SaveInFavorites(ctx context.Context, session_id string, city string) error {
	key := fmt.Sprintf("favorites:%s", session_id)
	if err := redis.Client().SAdd(ctx, key, city).Err(); err != nil {
		return fmt.Errorf("Не удалось добавить в избранные: %w", err)
	}
	return nil
}

func DeleteCity(ctx context.Context, session_id string, city string) error {
	key := fmt.Sprintf("favorites:%s", session_id)
	if err := redis.Client().SRem(ctx, key, city).Err(); err != nil {
		return fmt.Errorf("Не удалось удалить из избранных: %w", err)
	}
	return nil
}

func GetCities(ctx context.Context, session_id string) ([]string, error) {
	key := fmt.Sprintf("favorites:%s", session_id)
	value, err := redis.Client().SMembers(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить данные об избранных: %w", err)
	}
	return value, nil
}
