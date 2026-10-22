package favorites_redis_repository

import (
	"context"
	"fmt"
)

func (r *FavoritesRedisRepository) InvalidateFavorites(
	ctx context.Context,
	value string,
) error {
	key := fmt.Sprintf("favorites:v2:%s", value)
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis client del: %w", err)
	}

	return nil
}
