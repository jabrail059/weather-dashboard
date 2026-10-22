package favorites_redis_repository

import (
	"context"
	"fmt"
	"time"
)

const favoritesCacheTTL = 15 * time.Minute

func (r *FavoritesRedisRepository) AddFavorites(
	ctx context.Context,
	value string,
	cities []string,
) error {
	if len(cities) == 0 {
		return nil
	}

	key := fmt.Sprintf("favorites:v2:%s", value)
	members := make([]interface{}, len(cities))
	for i, city := range cities {
		members[i] = city
	}
	if err := r.client.SAdd(ctx, key, members...).Err(); err != nil {
		return fmt.Errorf("redis client sAdd: %w", err)
	}
	if err := r.client.Expire(ctx, key, favoritesCacheTTL).Err(); err != nil {
		_ = r.client.Del(ctx, key).Err()
		return fmt.Errorf("redis client expire: %w", err)
	}

	return nil
}
