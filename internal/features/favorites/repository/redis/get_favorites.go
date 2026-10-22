package favorites_redis_repository

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

func (r *FavoritesRedisRepository) GetFavorites(
	ctx context.Context,
	value string,
) (domain.FavoritesList, error) {
	key := fmt.Sprintf("favorites:v2:%s", value)
	cities, err := r.client.SMembers(ctx, key).Result()
	if err != nil {
		return domain.FavoritesList{}, fmt.Errorf(
			"redis client sMembers: %w",
			err,
		)
	}

	return domain.FavoritesList{
		Cities:   cities,
		Presence: len(cities) > 0,
	}, nil
}
