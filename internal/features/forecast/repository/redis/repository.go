package forecast_redis_repository

import core_redis_client "github.com/jabrail059/weather-dashboard/internal/core/repository/redis/client"

type ForecastRedisRepository struct {
	client core_redis_client.Client
}

func NewForecastRedisRepository(
	client core_redis_client.Client,
) *ForecastRedisRepository {
	return &ForecastRedisRepository{
		client: client,
	}
}
