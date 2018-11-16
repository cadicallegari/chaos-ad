package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	client *redis.Client
}

func (c *redisCache) Get(key string) (string, bool) {
	v, err := c.client.Get(key).Result()
	if err != nil {
		return "", false
	}

	return v, true
}

func (c *redisCache) Add(key, value string) error {
	return c.client.Set(key, value, 0).Err()
}

func (c *redisCache) Del(key string) error {
	return c.client.Del(key).Err()
}

func (c *redisCache) Hit(key string, ttl time.Duration) (bool, error) {
	return commonHitter(c, key, ttl)
}

func NewRedis(addr string, password string, db int) (*redisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return &redisCache{client: client}, nil
}
