package redis

import (
	"context"

	"github.com/fentezi/translete-word/internal/repositories"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisRepository(client *redis.Client) repositories.Repository {
	return &RedisRepository{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisRepository) Set(key string, value string) error {
	return r.client.Set(r.ctx, key, value, 0).Err()
}

func (r *RedisRepository) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisRepository) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
