package core_goredis_client

import (
	"context"
	"fmt"
	"time"

	core_redis_client "github.com/jabrail059/weather-dashboard/internal/core/repository/redis/client"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	*redis.Client
}

func NewClient(
	ctx context.Context,
	config Config,
) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("client ping: %w", err)
	}

	return &Client{
		Client: client,
	}, nil
}

func (c *Client) SMembers(
	ctx context.Context,
	key string,
) core_redis_client.StringSliceCmd {
	return StringSliceCmd{c.Client.SMembers(ctx, key)}
}

func (c *Client) SAdd(
	ctx context.Context,
	key string,
	members ...interface{},
) core_redis_client.IntCmd {
	return IntCmd{c.Client.SAdd(ctx, key, members...)}
}

func (c *Client) Del(
	ctx context.Context,
	keys ...string,
) core_redis_client.IntCmd {
	return IntCmd{c.Client.Del(ctx, keys...)}
}

func (c *Client) Expire(
	ctx context.Context,
	key string,
	expiration time.Duration,
) core_redis_client.StatusCmd {
	return c.Client.Expire(ctx, key, expiration)
}

func (c *Client) LRange(
	ctx context.Context,
	key string,
	start, stop int64,
) core_redis_client.StringSliceCmd {
	return StringSliceCmd{c.Client.LRange(ctx, key, start, stop)}
}

func (c *Client) LPush(
	ctx context.Context,
	key string,
	values ...interface{},
) core_redis_client.IntCmd {
	return IntCmd{c.Client.LPush(ctx, key, values...)}
}

func (c *Client) LRem(
	ctx context.Context,
	key string,
	count int64,
	value interface{},
) core_redis_client.IntCmd {
	return IntCmd{c.Client.LRem(ctx, key, count, value)}
}

func (c *Client) LTrim(
	ctx context.Context,
	key string,
	start, stop int64,
) core_redis_client.StatusCmd {
	return StatusCmd{c.Client.LTrim(ctx, key, start, stop)}
}

func (c *Client) Set(
	ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration,
) core_redis_client.StatusCmd {
	return StatusCmd{c.Client.Set(ctx, key, value, expiration)}
}

func (c *Client) Get(
	ctx context.Context,
	key string,
) core_redis_client.StringCmd {
	return StringCmd{c.Client.Get(ctx, key)}
}
