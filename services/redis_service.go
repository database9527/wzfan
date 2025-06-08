package services

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisService holds the Redis client.
type RedisService struct {
	client *redis.Client
}

// RedisConfig holds the configuration for Redis.
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// NewRedisService creates a new RedisService.
func NewRedisService(cfg RedisConfig) (*RedisService, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisService{client: rdb}, nil
}

// GetClient returns the underlying Redis client.
func (s *RedisService) GetClient() *redis.Client {
	return s.client
}
