//go:build windows
// +build windows

// Package fs provides cache with file system storage
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ MIN_EXPIRATION is minimal expiration duration
const MIN_EXPIRATION = time.Second

// ❗ MIN_CLEANUP_INTERVAL is minimal cleanup interval
const MIN_CLEANUP_INTERVAL = time.Second

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Cache is fs cache instance
type Cache struct{}

// ❗ Config is cache configuration
type Config struct {
	Dir               string
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

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
func (c *Cache) Set(key string, data any, expiration ...time.Duration) bool {
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

// ❗ Delete removes item from cache
func (c *Cache) Delete(key string) bool {
	panic("UNSUPPORTED")
}

// ❗ Flush removes all data from cache
func (c *Cache) Flush() bool {
	panic("UNSUPPORTED")
}
