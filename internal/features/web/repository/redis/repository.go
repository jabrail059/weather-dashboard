package web_redis_repository

import core_redis_client "github.com/jabrail059/weather-dashboard/internal/core/repository/redis/client"

type WebRedisRepository struct {
	client core_redis_client.Client
}

func NewWebRedisRepository(
	client core_redis_client.Client,
) *WebRedisRepository {
	return &WebRedisRepository{
		client: client,
	}
}
