package redis

import (
	"context"

	"github.com/akmuhammetakmyradov/test/pkg/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Configs) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
