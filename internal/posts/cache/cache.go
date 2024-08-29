package postscache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	client *redis.Client
}

func NewRedisRepo(client *redis.Client) *RedisRepo {
	return &RedisRepo{
		client: client,
	}
}

func (r *RedisRepo) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// Marshal the value to JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %v", err)
	}

	// Store the JSON string in Redis
	return r.client.Set(ctx, key, jsonData, expiration).Err()
}

func (r *RedisRepo) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisRepo) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisRepo) PaginationAdd(ctx context.Context, key string, score float64, data interface{}) error {
	return r.client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: data,
	}).Err()
}

func (r *RedisRepo) PaginationGet(ctx context.Context, key string, start, end int) ([]string, error) {
	data, err := r.client.ZRevRange(ctx, key, int64(start), int64(end)).Result()

	if err != nil {
		fmt.Println("err in PaginationGet: ", err)
		return data, err
	}

	return data, nil
}
