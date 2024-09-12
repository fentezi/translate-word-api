package repositories

import (
	"context"
	"fmt"

	"github.com/fentezi/translete-word/internal/database"
	"github.com/fentezi/translete-word/internal/utils/errors"
	"github.com/redis/go-redis/v9"
)

type wordRedisRepository struct {
	db  database.Database
	ctx context.Context
}

func NewWordRedisRepository(db database.Database) WordRepository {
	return &wordRedisRepository{
		db:  db,
		ctx: context.Background(),
	}
}

func (r *wordRedisRepository) Set(key string, value string) error {
	const op = "redis.Set"
	err := r.db.GetDB().Set(r.ctx, key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *wordRedisRepository) Get(key string) (string, error) {
	const op = "redis.Get"
	result, err := r.db.GetDB().Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("%s: %w", op, errors.ErrKeyNotFound)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}

func (r *wordRedisRepository) Delete(key string) error {
	const op = "redis.Delete"
	err := r.db.GetDB().Del(r.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
