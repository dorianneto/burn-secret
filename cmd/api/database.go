package api

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type redisClient struct {
	client *redis.Client
	logger *slog.Logger
}

func NewDatabase(logger *slog.Logger) (*redisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		logger.Error("database error connection", "error", err.Error())
		return nil, err
	}

	return &redisClient{client: rdb, logger: logger}, nil
}

func (c *redisClient) Get(key string) (interface{}, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *redisClient) Set(key string, value string) error {
	err := c.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
