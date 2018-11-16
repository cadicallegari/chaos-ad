package cache

import (
	"time"
)

type local struct {
	data map[string]string
}

func (c *local) Get(key string) (string, bool) {
	v, ok := c.data[key]
	return v, ok
}

func (c *local) Add(key string, value string) error {
	_, ok := c.Get(key)
	if ok {
		return ErrAlreadyExists
	}
	c.data[key] = value
	return nil
}

func (c *local) Del(key string) error {
	delete(c.data, key)
	return nil
}

// get hash from body
// check in Cache if hash exists
// if no add to Cache and return
// if yes: check the timestamp
func (c *local) Hit(key string, ttl time.Duration) (bool, error) {
	v, ok := c.data[key]

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

	return false, nil
}

func NewLocal() (*local, error) {
	return &local{data: make(map[string]string)}, nil
}
