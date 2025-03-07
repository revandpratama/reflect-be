package helper

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_POSTS_PATTERN = "posts:page:*"
)

func GetFromRedis(ctx context.Context, client *redis.Client, key string) ([]byte, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}

func SetToRedis(ctx context.Context, client *redis.Client, key string, value any) error {
	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	expirationTime := time.Minute * 5

	if err :=client.Set(ctx, key, byteValue, expirationTime).Err(); err != nil {
		return err
	}

	NewLog().Info("set data to redis")

	return nil
}

func DeleteRedisCache(ctx context.Context, client *redis.Client, pattern string) error {
	keys, err := client.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		NewLog().Info("deleting data from redis")
		return client.Del(ctx, keys...).Err()
	}


	return nil
}

func DeleteAllRedisCache(ctx context.Context, client *redis.Client) error {
	return client.FlushAll(ctx).Err()
}
