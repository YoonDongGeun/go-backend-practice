package db

import "github.com/redis/go-redis/v9"

func NewRedisClient(addr, password string, database int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	})
}
