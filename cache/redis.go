package cache

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

type RedisCache interface {
	GetRedisClient() (*redis.Client, context.Context, error)
}

type redisCache struct {
	dbNumber   int
	dbURI      string
	dbPassword string
}

func NewRedis(dbNumber int, dbURI, dbPassword string) *redisCache {
	return &redisCache{dbNumber: dbNumber, dbURI: dbURI, dbPassword: dbPassword}
}

func (r redisCache) GetRedisClient() (*redis.Client, context.Context, error) {

	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     r.dbURI,
		Password: r.dbPassword,
		DB:       r.dbNumber,
	})

	return client, ctx, nil

}
