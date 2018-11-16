package cache

import (
	"time"
)

type LocalCache struct {
	data map[string]time.Time
}

func (c *LocalCache) Get(key string) (time.Time, bool) {
	v, ok := c.data[key]
	return v, ok
}

func (c *LocalCache) Add(key string, value time.Time) error {
	_, ok := c.Get(key)
	if ok {
		return ErrAlreadyExists
	}
	c.data[key] = value
	return nil
}

func (s *LocalCache) Del(key string) error {
	delete(s.data, key)
	return nil
}

// get hash from body
// check in Cache if hash exists
// if no add to Cache and return
// if yes: check the timestamp
func (c *LocalCache) Hit(key string, ttl time.Duration) (bool, error) {
	v, ok := c.data[key]

	if !ok {
		v = time.Now()
		if err := c.Add(key, time.Now()); err != nil {
			return false, err
		}
		return true, nil
	}

	duration := time.Since(v)

	if duration < ttl {
		return false, nil
	}

	c.Del(key)

	if err := c.Add(key, time.Now()); err != nil {
		return false, err
	}

	return false, nil
}

func NewLocal() (*LocalCache, error) {
	return &LocalCache{data: make(map[string]time.Time)}, nil
}
