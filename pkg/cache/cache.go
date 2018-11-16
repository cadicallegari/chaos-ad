package cache

import (
	"errors"
	"time"
)

var (
	ErrAlreadyExists = errors.New("already exists")
)

type Cacher interface {
	Get(key string) (string, bool)
	Add(key, value string) error
	Del(key string) error
}

type CacherHitter interface {
	Cacher
	Hit(key string, expiration time.Duration) (bool, error)
}

// get hash from body
// check in Cache if hash exists
// if no add to Cache and return
// if yes: check the timestamp
func commonHitter(c CacherHitter, key string, ttl time.Duration) (bool, error) {
	v, ok := c.Get(key)

	if !ok {
		if err := c.Add(key, time.Now().Format(time.RFC3339)); err != nil {
			return false, err
		}
		return true, nil
	}

	dt, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return false, err
	}

	duration := time.Since(dt)

	if duration < ttl {
		return false, nil
	}

	c.Del(key)

	if err := c.Add(key, time.Now().Format(time.RFC3339)); err != nil {
		return false, err
	}

	return true, nil
}
