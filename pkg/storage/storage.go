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

func New() (*Storage, error) {
	return &Storage{data: make(map[string]time.Time)}, nil
}
