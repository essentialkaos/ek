// Package cache provides a simple in-memory key:value cache
package cache

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sync"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Cache is cache instance
type Cache struct {
	expiration     time.Duration
	data           map[string]any
	expiry         map[string]int64
	mu             *sync.RWMutex
	isJanitorWorks bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates new cache instance
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	s := &Cache{
		expiration: defaultExpiration,
		data:       make(map[string]any),
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

// Has returns true if cache contains data for given key
func (s *Cache) Has(key string) bool {
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

// Size returns number of items in cache
func (s *Cache) Size() int {
	if s == nil {
		return 0
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}

// Expired returns number of expired items in cache
func (s *Cache) Expired() int {
	if s == nil {
		return 0
	}

	items := 0
	now := time.Now().UnixNano()

	s.mu.Lock()

	for _, expiration := range s.expiry {
		if now > expiration {
			items++
		}
	}

	s.mu.Unlock()

	return items
}

// Set adds or updates item in cache
func (s *Cache) Set(key string, data any) {
	if s == nil {
		return
	}

	s.mu.Lock()

	s.expiry[key] = time.Now().Add(s.expiration).UnixNano()
	s.data[key] = data

	s.mu.Unlock()
}

// Get returns item from cache or nil
func (s *Cache) Get(key string) any {
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
func (s *Cache) GetWithExpiration(key string) (any, time.Time) {
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
func (s *Cache) Delete(key string) {
	if s == nil {
		return
	}

	s.mu.Lock()

	delete(s.data, key)
	delete(s.expiry, key)

	s.mu.Unlock()
}

// Flush removes all data from cache
func (s *Cache) Flush() {
	if s == nil {
		return
	}

	s.mu.Lock()

	s.data = make(map[string]any)
	s.expiry = make(map[string]int64)

	s.mu.Unlock()
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *Cache) janitor(interval time.Duration) {
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
