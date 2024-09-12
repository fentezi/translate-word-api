package database

import (
	"github.com/fentezi/translete-word/internal/config"
	"github.com/redis/go-redis/v9"
)

type redisDatabase struct {
	Db *redis.Client
}

var (
	dbInstance *redisDatabase
)

func NewRedisDatabase(conf *config.Config) Database {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Database.Host + ":" + conf.Database.Port,
		Password: conf.Database.Password,
		DB:       conf.Database.Name,
	})

	dbInstance = &redisDatabase{
		Db: client,
	}

	return dbInstance
}

func (r *redisDatabase) GetDB() *redis.Client {
	return r.Db
}
