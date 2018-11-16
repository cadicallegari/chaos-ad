// +build integration

package cache_test

import (
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis"

	"cadicallegari/chaos-ad/pkg/cache"
)

func newRedisCache(t *testing.T) (cache.CacherHitter, func()) {
	client, err := cache.NewRedis(
		os.Getenv("REDIS_URL"),
		"",
		0,
	)

	if err != nil {
		t.Errorf("Unexpected error connecting to redis: %s", err)
	}

	return client, func() {
		redisdb := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Password: "",
			DB:       0,
		})
		redisdb.FlushDB()
	}
}

func TestShouldReturnProperlyStatusWhenEmpty(t *testing.T) {
	c, teardown := newRedisCache(t)
	defer teardown()

	_, ok := c.Get("abc")
	if ok {
		t.Errorf("Should return false when key was not found")
	}
}

func TestShouldAddAndRecoveryProperly(t *testing.T) {
	c, teardown := newRedisCache(t)
	defer teardown()

	key := "thekey"

	_, ok := c.Get(key)
	if ok {
		t.Error("Should return false when key was not found")
	}

	if err := c.Add(key, time.Now().String()); err != nil {
		t.Errorf("Not expected error inserting value for key '%s'", key)
	}

	_, ok = c.Get(key)
	if !ok {
		t.Errorf("The key '%s' should exists, but don't =(", key)
	}

	if err := c.Del(key); err != nil {
		t.Errorf("Not expected error removing the key '%s': %s", key, err)
	}

	_, ok = c.Get(key)
	if ok {
		t.Error("the key was removed, why it was found?")
	}

}

func TestShouldHandleNewEntriesProperly(t *testing.T) {
	c, teardown := newRedisCache(t)
	defer teardown()

	key := "thekey"

	ok, err := c.Hit(key, time.Minute)
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}

	if !ok {
		t.Error("valid cache expected")
	}

	ok, err = c.Hit(key, time.Millisecond)
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}

	if !ok {
		t.Error("invalid cache expected")
	}

}

func TestShouldExpireCacheProperly(t *testing.T) {
	c, teardown := newRedisCache(t)
	defer teardown()

	key := "thekey"

	ok, err := c.Hit(key, time.Minute)
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}

	if !ok {
		t.Error("valid cache expected")
	}

	ok, err = c.Hit(key, time.Minute)
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}

	if ok {
		t.Error("invalid cache expected")
	}

}
