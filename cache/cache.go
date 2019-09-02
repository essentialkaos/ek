// Package cache provides simple in-memory key:value store
package cache

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sync"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Store is cache store
type Store struct {
	expiration time.Duration
	data       map[string]interface{}
	expiry     map[string]int64
	mu         *sync.RWMutex
}

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates new store instance
func New(defaultExpiration, cleanupInterval time.Duration) *Store {
	store := &Store{
		expiration: defaultExpiration,
		data:       make(map[string]interface{}),
		expiry:     make(map[string]int64),
		mu:         &sync.RWMutex{},
	}

	go store.janitor(cleanupInterval)

	return store
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Set adds or updates item in store
func (s *Store) Set(key string, data interface{}) {
	if s == nil {
		return
	}

	s.mu.Lock()

	s.expiry[key] = time.Now().Add(s.expiration).UnixNano()
	s.data[key] = data

	s.mu.Unlock()
}

// Get returns item from cache or nil
func (s *Store) Get(key string) interface{} {
	if s == nil {
		return nil
	}

	s.mu.RLock()

	expiration, ok := s.expiry[key]

	if !ok {
		s.mu.RUnlock()
		return nil
	}

	if time.Now().UnixNano() > expiration {
		s.mu.RUnlock()
		return nil
	}

	item := s.data[key]

	s.mu.RUnlock()

	return item
}

// GetWithExpiration returns item from cache and expiration date or nil
func (s *Store) GetWithExpiration(key string) (interface{}, time.Time) {
	if s == nil {
		return nil, time.Time{}
	}

	s.mu.RLock()

	expiration, ok := s.expiry[key]

	if !ok {
		s.mu.RUnlock()
		return nil, time.Time{}
	}

	if time.Now().UnixNano() > expiration {
		s.mu.RUnlock()
		return nil, time.Time{}
	}

	item := s.data[key]

	s.mu.RUnlock()

	return item, time.Unix(0, expiration)
}

// Delete removes item from cache
func (s *Store) Delete(key string) {
	if s == nil {
		return
	}

	s.mu.Lock()

	delete(s.data, key)
	delete(s.expiry, key)

	s.mu.Unlock()
}

// Flush removes all data from cache
func (s *Store) Flush() {
	if s == nil {
		return
	}

	s.mu.Lock()

	s.data = make(map[string]interface{})
	s.expiry = make(map[string]int64)

	s.mu.Unlock()
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *Store) janitor(interval time.Duration) {
	for range time.NewTicker(interval).C {
		if len(s.data) == 0 {
			continue
		}

		now := time.Now().UnixNano()

		s.mu.Lock()

		for key, expiration := range s.expiry {
			if now > expiration {
				delete(s.data, key)
				delete(s.expiry, key)
			}
		}

		s.mu.Unlock()
	}
}
