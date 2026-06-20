package service

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/storage/redis"
)

func AddSearchHistory(ctx context.Context, city string, cookieValue string) error {
	key := fmt.Sprintf("searchhistory:%s", cookieValue)

	cities, err := GetSearchHistory(ctx, cookieValue)
	if err != nil {
		return err
	}

	for _, c := range cities.Cities {
		if c == city {
			if err := redis.Client().LRem(ctx, key, 0, c).Err(); err != nil {
				return fmt.Errorf("Не удалось изменить историю поиска")
			}
			break
		}
	}

	if err := redis.Client().LPush(ctx, key, city).Err(); err != nil {
		return fmt.Errorf("Не удалось сохранить город в истории поиска")
	}

	if err := redis.Client().LTrim(ctx, key, 0, 4).Err(); err != nil {
		return fmt.Errorf("Ошибка лимита истории поиска")
	}

	return nil
}

func GetSearchHistory(ctx context.Context, cookieValue string) (*models.CitiesData, error) {
	key := fmt.Sprintf("searchhistory:%s", cookieValue)

	cities, err := redis.Client().LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить список городов")
	}

	return &models.CitiesData{
		Cities:   cities,
		Presence: len(cities) > 0,
	}, nil
}
