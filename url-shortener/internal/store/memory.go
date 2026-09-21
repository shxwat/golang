package main

import (
	"errors"

	"sync"
	"time"
)

type URL struct {
	ID          string
	OriginalURL string
	CreatedAt   time.Time
}
type Store interface {
	saveURL(u URL) error
	GetURL(shortCode string) (string, error)
}
type InMemoryStore struct {
	db   map[string]URL
	lock sync.Mutex
}

func (s *InMemoryStore) SaveURL(u URL) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.db[u.ID] = u
	return nil
}
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		db: make(map[string]URL),
	}
}
func (s *InMemoryStore) GetURL(shortCode string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	data, ok := s.db[shortCode]
	if !ok {
		return "", errors.New("URL not found")
	}
	return data.OriginalURL, nil
}
