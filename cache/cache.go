// Package cache provides methods and structs for caching data
package cache

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "time"

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
	Set(key string, data any) bool

	// GetWithExpiration returns item from cache and expiration date or nil
	GetWithExpiration(key string) (any, time.Time)

	// Delete removes item from cache
	Delete(key string) bool

	// Flush removes all data from cache
	Flush() bool
}
