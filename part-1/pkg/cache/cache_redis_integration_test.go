// +build integration

package cache_test

import (
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis"

	"cadicallegari/chaos-ad/pkg/cache"
)

func shouldHit(
	t *testing.T,
	expectation bool,
	c cache.CacherHitter,
	key string,
	ttl time.Duration,
) {
	ok, err := c.Hit(key, ttl)
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}

	if ok != expectation {
		t.Errorf("Expecting cache %t, got %t", expectation, ok)
	}

}

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

func TestShouldReturnReturnFalseWhenKeyAbsent(t *testing.T) {
	c, teardown := newRedisCache(t)
	defer teardown()

	_, ok := c.Get("not_exists")
	if ok {
		t.Errorf("Should return false when key was not found")
	}
}
func TestShouldAddAndGetKeyProperly(t *testing.T) {
	c, teardown := newRedisCache(t)
	defer teardown()

	key, value := "thekey", "thevalue"

	_, ok := c.Get(key)
	if ok {
		t.Error("Should return false when key was not found")
	}

	if err := c.Add(key, value); err != nil {
		t.Errorf("Not expected error inserting value for key '%s'", key)
	}

	got, ok := c.Get(key)
	if !ok {
		t.Errorf("The key '%s' should exists, but don't =(", key)
	}

	if value != got {
		t.Errorf("Expecting: '%s', got '%s'", value, got)
	}

	if err := c.Del(key); err != nil {
		t.Errorf("Not expected error removing the key '%s': %s", key, err)
	}

	_, ok = c.Get(key)
	if ok {
		t.Error("the key was removed, why it was found?")
	}

}

func TestShouldHitOkIfTTLHasPassed(t *testing.T) {
	c, teardown := newRedisCache(t)
	defer teardown()
	key := "thekey"
	shouldHit(t, true, c, key, time.Minute)
	shouldHit(t, true, c, key, time.Millisecond)
	shouldHit(t, true, c, key, time.Millisecond)
}

func TestShouldHitFalseIfTTLNotEnded(t *testing.T) {
	c, teardown := newRedisCache(t)
	defer teardown()
	key := "thekey"
	shouldHit(t, true, c, key, time.Millisecond)
	shouldHit(t, false, c, key, time.Minute)
	shouldHit(t, true, c, key, time.Millisecond)
}
