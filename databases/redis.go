package databases

import (
	"module-crud/config"

	"github.com/go-redis/redis/v8"
)

func ConnectionRedisDB(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword, // no password set
		DB:       0,                 // use default DB
	})
}
