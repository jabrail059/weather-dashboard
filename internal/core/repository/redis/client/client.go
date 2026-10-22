package core_redis_client

import (
	"context"
	"errors"
	"time"
)

var ErrKeyNotFound = errors.New("redis key not found")

type Client interface {
	SAdd(ctx context.Context, key string, members ...interface{}) IntCmd
	Del(ctx context.Context, keys ...string) IntCmd
	SMembers(ctx context.Context, key string) StringSliceCmd
	Expire(ctx context.Context, key string, expiration time.Duration) StatusCmd

	LRange(ctx context.Context, key string, start, stop int64) StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) IntCmd
	LPush(ctx context.Context, key string, values ...interface{}) IntCmd
	LTrim(ctx context.Context, key string, start, stop int64) StatusCmd

	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) StatusCmd
	Get(ctx context.Context, key string) StringCmd
}

type IntCmd interface {
	Result() (int64, error)
	Err() error
}

type StringSliceCmd interface {
	Result() ([]string, error)
	Err() error
}

type StatusCmd interface {
	Err() error
}

type StringCmd interface {
	Result() (string, error)
	Err() error
}
