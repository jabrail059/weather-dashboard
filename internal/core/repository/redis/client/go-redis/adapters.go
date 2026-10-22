package core_goredis_client

import (
	"errors"

	core_redis_client "github.com/jabrail059/weather-dashboard/internal/core/repository/redis/client"
	"github.com/redis/go-redis/v9"
)

// тут опиши импементации интерфейсов возвращаемых типов если таковые будут
type StringSliceCmd struct {
	*redis.StringSliceCmd
}

type IntCmd struct {
	*redis.IntCmd
}

type StatusCmd struct {
	*redis.StatusCmd
}

type StringCmd struct {
	*redis.StringCmd
}

func (c StringCmd) Result() (string, error) {
	value, err := c.StringCmd.Result()
	if errors.Is(err, redis.Nil) {
		return "", core_redis_client.ErrKeyNotFound
	}
	return value, err
}

func (c StringCmd) Err() error {
	err := c.StringCmd.Err()
	if errors.Is(err, redis.Nil) {
		return core_redis_client.ErrKeyNotFound
	}
	return err
}
