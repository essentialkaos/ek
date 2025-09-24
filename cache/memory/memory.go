// Package memory provides cache with in-memory storage
package memory

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"sync"
	"time"

	"github.com/essentialkaos/ek/v13/cache"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DEFAULT_EXPIRATION is default expiration
const DEFAULT_EXPIRATION = cache.MINUTE

// MIN_EXPIRATION is minimal expiration duration
const MIN_EXPIRATION = cache.MILLISECOND

// MIN_CLEANUP_INTERVAL is minimal cleanup interval
const MIN_CLEANUP_INTERVAL = cache.MILLISECOND

// ////////////////////////////////////////////////////////////////////////////////// //

// validate storage interface
var _ cache.Cache = (*Cache)(nil)

// ////////////////////////////////////////////////////////////////////////////////// //

// Cache is in-memory cache instance
type Cache struct {
	data           map[string]any
	expiry         map[string]int64
	mu             *sync.RWMutex
	expiration     cache.Duration
	isJanitorWorks bool
}

// Config is cache configuration
type Config struct {
	DefaultExpiration cache.Duration
	CleanupInterval   cache.Duration
}

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates new cache instance
func New(config Config) (*Cache, error) {
	err := config.Validate()

	if err != nil {
		return nil, err
	}

	c := &Cache{
		expiration: DEFAULT_EXPIRATION,
		data:       make(map[string]any),
		expiry:     make(map[string]int64),
		mu:         &sync.RWMutex{},
	}

	if config.DefaultExpiration != 0 {
		c.expiration = config.DefaultExpiration
	}

	if config.CleanupInterval != 0 {
		c.isJanitorWorks = true
		go c.janitor(config.CleanupInterval)
	}

	return c, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Has returns true if cache contains data for given key
func (c *Cache) Has(key string) bool {
	if c == nil {
		return false
	}

	c.mu.RLock()

	expiration, ok := c.expiry[key]

	if !ok {
		c.mu.RUnlock()
		return false
	}

	if time.Now().UnixNano() > expiration {
		c.mu.RUnlock()

		if !c.isJanitorWorks {
			c.Delete(key)
		}

		return false
	}

	c.mu.RUnlock()

	return ok
}

// Size returns number of items in cache
func (c *Cache) Size() int {
	if c == nil {
		return 0
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.data)
}

// Expired returns number of expired items in cache
func (c *Cache) Expired() int {
	if c == nil {
		return 0
	}

	items := 0
	now := time.Now().UnixNano()

	c.mu.Lock()

	for _, expiration := range c.expiry {
		if now > expiration {
			items++
		}
	}

	c.mu.Unlock()

	return items
}

// Set adds or updates item in cache
func (c *Cache) Set(key string, data any, expiration ...time.Duration) bool {
	if c == nil || data == nil {
		return false
	}

	c.mu.Lock()

	if len(expiration) > 0 && expiration[0] > MIN_EXPIRATION {
		c.expiry[key] = time.Now().Add(expiration[0]).UnixNano()
	} else {
		c.expiry[key] = time.Now().Add(c.expiration).UnixNano()
	}

	c.data[key] = data

	c.mu.Unlock()

	return true
}

// Get returns item from cache or nil
func (c *Cache) Get(key string) any {
	if c == nil {
		return nil
	}

	c.mu.RLock()

	expiration, ok := c.expiry[key]

	if !ok {
		c.mu.RUnlock()
		return nil
	}

	if time.Now().UnixNano() > expiration {
		c.mu.RUnlock()

		if !c.isJanitorWorks {
			c.Delete(key)
		}

		return nil
	}

	item := c.data[key]

	c.mu.RUnlock()

	return item
}

// GetWithExpiration returns item expiration date
func (c *Cache) GetExpiration(key string) time.Time {
	if c == nil {
		return time.Time{}
	}

	c.mu.RLock()

	expiration, ok := c.expiry[key]

	if !ok {
		c.mu.RUnlock()
		return time.Time{}
	}

	c.mu.RUnlock()

	return time.Unix(0, expiration)
}

// GetWithExpiration returns item from cache and expiration date or nil
func (c *Cache) GetWithExpiration(key string) (any, time.Time) {
	if c == nil {
		return nil, time.Time{}
	}

	c.mu.RLock()

	expiration, ok := c.expiry[key]

	if !ok {
		c.mu.RUnlock()
		return nil, time.Time{}
	}

	if time.Now().UnixNano() > expiration {
		c.mu.RUnlock()

		if !c.isJanitorWorks {
			c.Delete(key)
		}

		return nil, time.Time{}
	}

	item := c.data[key]

	c.mu.RUnlock()

	return item, time.Unix(0, expiration)
}

// Keys is an iterator over cache keys
func (c *Cache) Keys(yield func(k string) bool) {
	if c == nil {
		return
	}

	c.mu.RLock()

	for k := range c.data {
		if !yield(k) {
			c.mu.RUnlock()
			return
		}
	}

	c.mu.RUnlock()
}

// All is an iterator over cache items
func (c *Cache) All(yield func(k string, v any) bool) {
	if c == nil {
		return
	}

	c.mu.RLock()

	for k, v := range c.data {
		if !yield(k, v) {
			c.mu.RUnlock()
			return
		}
	}

	c.mu.RUnlock()
}

// Delete removes item from cache
func (c *Cache) Delete(key string) bool {
	if c == nil {
		return false
	}

	c.mu.Lock()

	delete(c.data, key)
	delete(c.expiry, key)

	c.mu.Unlock()

	return true
}

// Flush removes all data from cache
func (c *Cache) Flush() bool {
	if c == nil {
		return false
	}

	c.mu.Lock()

	c.data = make(map[string]any)
	c.expiry = make(map[string]int64)

	c.mu.Unlock()

	return true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Validate validates cache configuration
func (c Config) Validate() error {
	if c.DefaultExpiration != 0 && c.DefaultExpiration < MIN_EXPIRATION {
		return errors.New("Invalid configuration: Expiration is too short (< 1ms)")
	}

	if c.CleanupInterval != 0 && c.CleanupInterval < MIN_CLEANUP_INTERVAL {
		return errors.New("Invalid configuration: Cleanup interval is too short (< 1ms)")
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// janitor is cache cleanup job
func (c *Cache) janitor(interval time.Duration) {
	for range time.NewTicker(interval).C {
		if len(c.data) == 0 {
			continue
		}

		now := time.Now().UnixNano()

		c.mu.Lock()

		for key, expiration := range c.expiry {
			if now > expiration {
				delete(c.data, key)
				delete(c.expiry, key)
			}
		}

		c.mu.Unlock()
	}
}
