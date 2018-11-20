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

func (c *local) Hit(key string, ttl time.Duration) (bool, error) {
	return commonHitter(c, key, ttl)
}

func NewLocal() (*local, error) {
	return &local{data: make(map[string]string)}, nil
}
