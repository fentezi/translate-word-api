package database

import (
	"github.com/redis/go-redis/v9"
)

type Database interface {
	GetDB() *redis.Client
}
