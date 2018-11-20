// +build unit

package cache_test

import (
	"testing"
	"time"

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

func TestShouldReturnReturnFalseWhenKeyAbsent(t *testing.T) {
	c, err := cache.NewLocal()
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}

	_, ok := c.Get("not_exists")
	if ok {
		t.Errorf("Should return false when key was not found")
	}
}

func TestShouldAddAndGetKeyProperly(t *testing.T) {
	c, err := cache.NewLocal()
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}

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
		t.Errorf("The key '%s' should exist, but don't =(", key)
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
	local, _ := cache.NewLocal()
	key := "thekey"
	shouldHit(t, true, local, key, time.Minute)
	shouldHit(t, true, local, key, time.Millisecond)
	shouldHit(t, true, local, key, time.Millisecond)
}

func TestShouldHitFalseIfTTLNotEnded(t *testing.T) {
	local, _ := cache.NewLocal()
	key := "thekey"
	shouldHit(t, true, local, key, time.Millisecond)
	shouldHit(t, false, local, key, time.Minute)
	shouldHit(t, true, local, key, time.Millisecond)
}
