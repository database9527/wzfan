package utils

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// SyncLocalDataToRedis syncs data from local text files to Redis sets.
// dataDir is the directory containing files like #tagName#[01].txt
// siteID is used in the Redis key.
// localDataDbIndex is the Redis database index to use.
func SyncLocalDataToRedis(redisClient *redis.Client, dataDir string, siteID string, localDataDbIndex int) error {
	ctx := context.Background()

	// It's crucial to ensure operations are on the correct DB.
	// This implementation creates a new client instance configured for the specific DB.
	// Alternatively, if the main client is passed, ensure SELECT is handled carefully.
	// For simplicity, let's assume the passed client is ready or we create a specific one.
	// If the passed client is a general client, we would need to select DB.
	// For this example, let's assume the passed client is already configured for the correct DB
	// or that its options allow this (not typical for go-redis where client is bound to options).
	// A more robust solution might involve creating a new client instance for this specific DB
	// using options from a config, or ensuring the provided client is specifically for this DB.

	// Let's use a temporary client that clones options and sets the DB.
	// This is safer than trying to SELECT and then SELECT back on a shared client.
	opts := redisClient.Options()
	dbSpecificClient := redis.NewClient(&redis.Options{
		Addr:     opts.Addr,
		Password: opts.Password,
		DB:       localDataDbIndex, // Target specific DB
	})
	defer dbSpecificClient.Close() // Ensure this temporary client is closed

	// Check connectivity for the DB-specific client
	if _, err := dbSpecificClient.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("failed to connect to Redis DB %d: %w", localDataDbIndex, err)
	}

	pattern := filepath.Join(dataDir, "#*#[01].txt")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("error matching files with glob pattern %s: %w", pattern, err)
	}

	for _, filePath := range matches {
		fileName := filepath.Base(filePath)
		parts := strings.Split(strings.TrimSuffix(fileName, ".txt"), "#")
		if len(parts) < 2 || parts[0] != "" { // Should start with #
			// log.Printf("Skipping file with unexpected format: %s", fileName)
			continue
		}
		tagName := parts[1]

		redisKey := fmt.Sprintf("site:%s:local_data:%s", siteID, tagName)
		redisTimestampKey := fmt.Sprintf("site:%s:local_data:%s:updated_at", siteID, tagName)

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			// log.Printf("Error getting file info for %s: %v", filePath, err)
			continue
		}
		fileModTime := fileInfo.ModTime()

		redisUpdatedAtStr, err := dbSpecificClient.Get(ctx, redisTimestampKey).Result()
		if err != nil && err != redis.Nil {
			// log.Printf("Error getting Redis timestamp for %s: %v", redisTimestampKey, err)
			continue
		}

		var redisUpdatedAt time.Time
		if err == nil && redisUpdatedAtStr != "" {
			ts, _ := strconv.ParseInt(redisUpdatedAtStr, 10, 64)
			redisUpdatedAt = time.Unix(ts, 0)
		}

		if err == redis.Nil || fileModTime.After(redisUpdatedAt) {
			// File is new or updated, proceed with sync
			file, err := os.Open(filePath)
			if err != nil {
				// log.Printf("Error opening file %s: %v", filePath, err)
				continue
			}

			var lines []string
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" {
					lines = append(lines, line)
				}
			}
			file.Close() // Close file immediately after reading

			if err := scanner.Err(); err != nil {
				// log.Printf("Error reading file %s: %v", filePath, err)
				continue
			}

			if len(lines) > 0 {
				pipe := dbSpecificClient.Pipeline()
				pipe.Del(ctx, redisKey)
				// SADD in go-redis expects []interface{} for multiple members
				saddArgs := make([]interface{}, len(lines))
				for i, v := range lines {
					saddArgs[i] = v
				}
				pipe.SAdd(ctx, redisKey, saddArgs...)
				pipe.Set(ctx, redisTimestampKey, fileModTime.Unix(), 0)
				_, err := pipe.Exec(ctx)
				if err != nil {
					// log.Printf("Error executing Redis pipeline for %s: %v", redisKey, err)
					continue
				}
				// log.Printf("Synced %s to Redis", fileName)
			} else {
				// If file is empty, remove existing keys
				pipe := dbSpecificClient.Pipeline()
				pipe.Del(ctx, redisKey)
				pipe.Del(ctx, redisTimestampKey)
				_, err := pipe.Exec(ctx)
				if err != nil {
					// log.Printf("Error deleting Redis keys for empty file %s: %v", fileName, err)
				}
				// log.Printf("Removed Redis keys for empty file %s", fileName)
			}
		}
	}
	return nil
}
