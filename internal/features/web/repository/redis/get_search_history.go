package web_redis_repository

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

func (r *WebRedisRepository) GetSearchHistory(
	ctx context.Context,
	value string,
) (domain.HistoryList, error) {
	key := fmt.Sprintf("searchhistory:%s", value)

	cities, err := r.client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return domain.HistoryList{}, fmt.Errorf(
			"redis client lrange: %w",
			err,
		)
	}

	return domain.HistoryList{
		Cities:   cities,
		Presence: len(cities) > 0,
	}, nil
}
