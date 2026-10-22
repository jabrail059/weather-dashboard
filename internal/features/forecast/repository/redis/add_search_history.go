package forecast_redis_repository

import (
	"context"
	"fmt"
)

func (r *ForecastRedisRepository) AddSearchHistory(
	ctx context.Context,
	city string,
	value string,
) error {
	key := fmt.Sprintf("searchhistory:%s", value)

	cities, err := r.client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return fmt.Errorf(
			"redis client lrange: %w",
			err,
		)
	}

	for _, c := range cities {
		if c == city {
			if err := r.client.LRem(ctx, key, 0, c).Err(); err != nil {
				return fmt.Errorf("redis client lrem, %w", err)
			}
			break
		}
	}

	if err := r.client.LPush(ctx, key, city).Err(); err != nil {
		return fmt.Errorf("redis client lpush, %w", err)
	}

	if err := r.client.LTrim(ctx, key, 0, 4).Err(); err != nil {
		return fmt.Errorf("redis client ltrim: %w", err)
	}

	return nil
}
