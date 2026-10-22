package favorites_redis_repository

import core_redis_client "github.com/jabrail059/weather-dashboard/internal/core/repository/redis/client"

type FavoritesRedisRepository struct {
	client core_redis_client.Client
}

func NewFavoritesRedisRepository(
	client core_redis_client.Client,
) *FavoritesRedisRepository {
	return &FavoritesRedisRepository{
		client: client,
	}
}
