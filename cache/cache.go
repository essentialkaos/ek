// Package cache provides methods and structs for caching data
package cache

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "time"

// ////////////////////////////////////////////////////////////////////////////////// //

// Duration is time.Duration alias
type Duration = time.Duration

// ////////////////////////////////////////////////////////////////////////////////// //

// Cache is cache backend interface
type Cache interface {
	// Has returns true if cache contains data for given key
	Has(key string) bool

	// Size returns number of items in cache
	Size() int

	// Expired returns number of expired items in cache
	Expired() int

	// Set adds or updates item in cache
	Set(key string, data any, expiration ...time.Duration) bool

	// GetWithExpiration returns item from cache
	Get(key string) any

	// GetWithExpiration returns item expiration date
	GetExpiration(key string) time.Time

	// GetWithExpiration returns item from cache and expiration date or nil
	GetWithExpiration(key string) (any, time.Time)

	// Keys is an iterator over cache keys
	Keys(yield func(k string) bool)

	// All is an iterator over cache items
	All(yield func(k string, v any) bool)

	// Delete removes item from cache
	Delete(key string) bool

	// Flush removes all data from cache
	Flush() bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	MILLISECOND = time.Millisecond // 1 ms
	SECOND      = time.Second      // 1 sec
	MINUTE      = time.Minute      // 1 min
	HOUR        = time.Hour        // 1 hr
	DAY         = 24 * HOUR        // 24 hr
	WEEK        = 7 * DAY          // 7 d
	MONTH       = 30 * DAY         // 30 d
	YEAR        = 365 * DAY        // 365 d
)
