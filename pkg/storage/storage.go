package storage

import (
	"errors"
	"time"
)

type Storage struct {
	data map[string]time.Time
}

var (
	ErrAlreadyExists = errors.New("already exists")
)

func (s *Storage) Lookup(key string) (time.Time, bool) {
	v, ok := s.data[key]
	return v, ok
}

func (s *Storage) Add(key string, value time.Time) error {
	_, ok := s.Lookup(key)
	if ok {
		return ErrAlreadyExists
	}
	s.data[key] = value
	return nil
}

func (s *Storage) Del(key string) error {
	delete(s.data, key)
	return nil
}

// get hash from body
// check in storage if hash exists
// if no add to storage and return
// if yes: check the timestamp
func (s *Storage) CheckCache(hash string, ttl time.Duration) (bool, error) {
	v, ok := s.Lookup(hash)

	if !ok {
		if err := s.Add(hash, time.Now()); err != nil {
			return false, err
		}
		return true, nil
	}

	duration := time.Since(v)

	if duration < ttl {
		return false, nil
	}

	s.Del(hash)

	if err := s.Add(hash, time.Now()); err != nil {
		return false, err
	}

	return false, nil
}

func New() (*Storage, error) {
	return &Storage{data: make(map[string]time.Time)}, nil
}
