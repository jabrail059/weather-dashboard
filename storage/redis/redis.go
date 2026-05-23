package redis

import "github.com/redis/go-redis/v9"

var client *redis.Client

func SetClient(redis *redis.Client) {
	client = redis
}

func Client() *redis.Client {
	return client
}
