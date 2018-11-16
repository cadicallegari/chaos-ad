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
