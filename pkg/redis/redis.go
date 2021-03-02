package redis

import (
	"context"
	"time"

	"github.com/AleksK1NG/nats-streaming/config"
	"github.com/go-redis/redis/v8"
)

// NewRedisClient Returns new redis client
func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	redisHost := cfg.Redis.RedisAddr

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: cfg.Redis.MinIdleConn,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.RedisPassword, // no password set
		DB:           cfg.Redis.DB,            // use default DB
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
