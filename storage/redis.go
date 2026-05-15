package storage

import "github.com/redis/go-redis/v9"

var client *redis.Client

func SetRedisClient(redis *redis.Client) {
	client = redis
}

func Redis() *redis.Client {
	return client
}
