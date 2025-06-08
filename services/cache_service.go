package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// CacheService holds the Redis client and cache configuration.
type CacheService struct {
	redisClient *redis.Client
	ttl         time.Duration
	prefix      string
	enabled     bool
}

// CacheConfig holds the configuration for the cache.
type CacheConfig struct {
	TTLSeconds int
	Prefix     string
	Enabled    bool
}

// NewCacheService creates a new CacheService.
func NewCacheService(redisService *RedisService, cfg CacheConfig) *CacheService {
	return &CacheService{
		redisClient: redisService.GetClient(), // Assumes GetClient() returns a non-nil client
		ttl:         time.Duration(cfg.TTLSeconds) * time.Second,
		prefix:      cfg.Prefix, // This will be added later if needed
		enabled:     cfg.Enabled,
	}
}

// Get retrieves a value from the cache.
func (s *CacheService) Get(key string) (string, error) {
	if !s.enabled {
		return "", redis.Nil // Or a custom error indicating cache is disabled
	}
	fullKey := s.prefix + key
	val, err := s.redisClient.Get(context.Background(), fullKey).Result()
	if err == redis.Nil {
		return "", redis.Nil // Key does not exist
	} else if err != nil {
		return "", err // Other Redis error
	}
	return val, nil
}

// Set adds a value to the cache.
func (s *CacheService) Set(key string, value string, ttl time.Duration) error {
	if !s.enabled {
		return nil // Cache is disabled
	}
	fullKey := s.prefix + key
	effectiveTTL := s.ttl // Default TTL
	if ttl > 0 {
		effectiveTTL = ttl // Use specific TTL if provided
	}
	err := s.redisClient.Set(context.Background(), fullKey, value, effectiveTTL).Err()
	return err
}
