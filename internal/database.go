package internal

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type redisClient struct {
	client *redis.Client
	logger *slog.Logger
}

func NewDatabase(logger *slog.Logger) (*redisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT")),
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

func (c *redisClient) Insert(key string, data interface{}) error {
	err := c.client.HSet(ctx, key, data).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *redisClient) Select(key string, field string) (string, error) {
	cmd := c.client.HGet(ctx, key, field)
	if err := cmd.Err(); err != nil {
		return "", err
	}

	return cmd.Result()
}

func (c *redisClient) SelectAll(key string, output interface{}) error {
	cmd := c.client.HGetAll(ctx, key)
	if err := cmd.Err(); err != nil {
		return err
	}

	err := cmd.Scan(output)
	if err != nil {
		return err
	}

	return nil
}

func (c *redisClient) Delete(key string) (int64, error) {
	cmd := c.client.Del(ctx, key)
	if err := cmd.Err(); err != nil {
		return 0, err
	}

	return cmd.Result()
}
