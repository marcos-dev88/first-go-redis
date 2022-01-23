package cache

import "time"

type Cache interface {
	GetByAddr(addrName string) (string, error)
	Create(addrName string, data interface{}, ttl time.Duration) error
}

type cache struct {
	rediscache RedisCache
}

func NewCache(rediscache RedisCache) *cache {
	return &cache{rediscache: rediscache}
}

func (c cache) GetByAddr(addrName string) (string, error) {
	client, ctx, err := c.rediscache.GetRedisClient()

	defer client.Close()
	if err != nil {
		return "", err
	}

	return client.Get(ctx, addrName).Result()
}

func (c cache) Create(addrName string, data interface{}, ttl time.Duration) error {
	client, ctx, err := c.rediscache.GetRedisClient()
	defer client.Close()
	err = client.Set(ctx, addrName, data, ttl).Err()

	if err != nil {
		return err
	}

	return nil
}
