package web_redis_repository

import (
	"github.com/redis/go-redis/v9"
)

type WebRedisRepository struct {
	client *redis.Client
}

func NewWebRedisRepository(
	client *redis.Client,
) *WebRedisRepository {
	return &WebRedisRepository{
		client: client,
	}
}

func NewRedisClient(
	config Config,
) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	return client
}
