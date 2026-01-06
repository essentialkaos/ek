//go:build windows

// Package fs provides cache with file system storage
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"

	"github.com/essentialkaos/ek/v13/cache"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ MIN_EXPIRATION is minimal expiration duration
const MIN_EXPIRATION = cache.SECOND

// ❗ MIN_CLEANUP_INTERVAL is minimal cleanup interval
const MIN_CLEANUP_INTERVAL = cache.SECOND

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Cache is fs cache instance
type Cache struct{}

// ❗ Config is cache configuration
type Config struct {
	Dir               string
	DefaultExpiration cache.Duration
	CleanupInterval   cache.Duration
}

// ////////////////////////////////////////////////////////////////////////////////// //

// validate storage interface
var _ cache.Cache = (*Cache)(nil)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ New creates new cache instance
func New(config Config) (*Cache, error) {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Has returns true if cache contains data for given key
func (c *Cache) Has(key string) bool {
	panic("UNSUPPORTED")
}

// ❗ Size returns number of items in cache
func (c *Cache) Size() int {
	panic("UNSUPPORTED")
}

// ❗ Expired returns number of expired items in cache
func (c *Cache) Expired() int {
	panic("UNSUPPORTED")
}

// ❗ Set adds or updates item in cache
func (c *Cache) Set(key string, data any, expiration ...cache.Duration) bool {
	panic("UNSUPPORTED")
}

// ❗ GetWithExpiration returns item from cache
func (c *Cache) Get(key string) any {
	panic("UNSUPPORTED")
}

// ❗ GetWithExpiration returns item expiration date
func (c *Cache) GetExpiration(key string) time.Time {
	panic("UNSUPPORTED")
}

// ❗ GetWithExpiration returns item from cache and expiration date or nil
func (c *Cache) GetWithExpiration(key string) (any, time.Time) {
	panic("UNSUPPORTED")
}

// ❗ Keys is an iterator over cache keys
func (c *Cache) Keys(yield func(k string) bool) {
	panic("UNSUPPORTED")
}

// ❗ All is an iterator over cache items
func (c *Cache) All(yield func(k string, v any) bool) {
	panic("UNSUPPORTED")
}

// ❗ Invalidate deletes all expired records
func (c *Cache) Invalidate() bool {
	panic("UNSUPPORTED")
}

// ❗ Delete removes item from cache
func (c *Cache) Delete(key string) bool {
	panic("UNSUPPORTED")
}

// ❗ Flush removes all data from cache
func (c *Cache) Flush() bool {
	panic("UNSUPPORTED")
}
