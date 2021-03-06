// Package cache provides simple in-memory key:value store
package cache

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sync"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Store is cache store
type Store struct {
	expiration     time.Duration
	data           map[string]interface{}
	expiry         map[string]int64
	mu             *sync.RWMutex
	isJanitorWorks bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates new store instance
func New(defaultExpiration, cleanupInterval time.Duration) *Store {
	s := &Store{
		expiration: defaultExpiration,
		data:       make(map[string]interface{}),
		expiry:     make(map[string]int64),
		mu:         &sync.RWMutex{},
	}

	if cleanupInterval != 0 {
		s.isJanitorWorks = true
		go s.janitor(cleanupInterval)
	}

	return s
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Has returns true if store contains data for given key
func (s *Store) Has(key string) bool {
	if s == nil {
		return false
	}

	s.mu.RLock()

	expiration, ok := s.expiry[key]

	if !ok {
		s.mu.RUnlock()
		return false
	}

	if time.Now().UnixNano() > expiration {
		s.mu.RUnlock()

		if !s.isJanitorWorks {
			s.Delete(key)
		}

		return false
	}

	s.mu.RUnlock()

	return ok
}

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

		if !s.isJanitorWorks {
			s.Delete(key)
		}

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

		if !s.isJanitorWorks {
			s.Delete(key)
		}

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
