//go:build !windows
// +build !windows

// Package fs provides cache with file system storage
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"hash"
	"os"
	"path"
	"time"

	"github.com/essentialkaos/ek/v13/cache"
	"github.com/essentialkaos/ek/v13/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MIN_EXPIRATION is minimal expiration duration
const MIN_EXPIRATION = time.Second

// MIN_CLEANUP_INTERVAL is minimal cleanup interval
const MIN_CLEANUP_INTERVAL = time.Second

// ////////////////////////////////////////////////////////////////////////////////// //

// Cache is fs cache instance
type Cache struct {
	dir            string
	hasher         hash.Hash
	expiration     time.Duration
	isJanitorWorks bool
}

// Config is cache configuration
type Config struct {
	Dir               string
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

// ////////////////////////////////////////////////////////////////////////////////// //

// cacheItem is cache item
type cacheItem struct {
	Data any
}

// ////////////////////////////////////////////////////////////////////////////////// //

// validate storage interface
var _ cache.Cache = (*Cache)(nil)

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates new cache instance
func New(config Config) (*Cache, error) {
	err := config.Validate()

	if err != nil {
		return nil, fmt.Errorf("Invalid configuration: %w", err)
	}

	c := &Cache{
		dir:        config.Dir,
		expiration: config.DefaultExpiration,
		hasher:     sha256.New(),
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
	if c == nil || key == "" {
		return false
	}

	return fsutil.IsExist(c.getItemPath(key, false))
}

// Size returns number of items in cache
func (c *Cache) Size() int {
	if c == nil {
		return 0
	}

	return len(fsutil.List(c.dir, true))
}

// Expired returns number of expired items in cache
func (c *Cache) Expired() int {
	if c == nil {
		return 0
	}

	expired := 0
	now := time.Now()
	items := fsutil.List(c.dir, true)

	fsutil.ListToAbsolute(c.dir, items)

	for _, item := range items {
		mtime, _ := fsutil.GetMTime(item)

		if mtime.Before(now) {
			expired++
		}
	}

	return expired
}

// Set adds or updates item in cache
func (c *Cache) Set(key string, data any, expiration ...time.Duration) bool {
	if c == nil || data == nil || key == "" {
		return false
	}

	tmpFile := c.getItemPath(key, true)
	fd, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)

	if err != nil {
		return false
	}

	err = gob.NewEncoder(fd).Encode(&cacheItem{data})

	fd.Close()

	if err != nil {
		os.Remove(tmpFile)
		return false
	}

	itemFile := c.getItemPath(key, false)

	err = os.Rename(tmpFile, itemFile)

	if err != nil {
		os.Remove(tmpFile)
		return false
	}

	expr := c.expiration

	if len(expiration) > 0 {
		expr = expiration[0]
	}

	return os.Chtimes(itemFile, time.Time{}, time.Now().Add(expr)) == err
}

// GetWithExpiration returns item from cache
func (c *Cache) Get(key string) any {
	if c == nil || key == "" {
		return nil
	}

	fd, err := os.Open(c.getItemPath(key, false))

	if err != nil {
		return nil
	}

	item := &cacheItem{}
	err = gob.NewDecoder(fd).Decode(item)

	if err != nil {
		return nil
	}

	return item.Data
}

// GetWithExpiration returns item expiration date
func (c *Cache) GetExpiration(key string) time.Time {
	if c == nil || key == "" {
		return time.Time{}
	}

	mt, _ := fsutil.GetMTime(c.getItemPath(key, false))

	return mt
}

// GetWithExpiration returns item from cache and expiration date or nil
func (c *Cache) GetWithExpiration(key string) (any, time.Time) {
	if c == nil || key == "" {
		return nil, time.Time{}
	}

	return c.Get(key), c.GetExpiration(key)
}

// Delete removes item from cache
func (c *Cache) Delete(key string) bool {
	if c == nil {
		return false
	}

	return os.Remove(c.getItemPath(key, false)) == nil
}

// Flush removes all data from cache
func (c *Cache) Flush() bool {
	if c == nil {
		return false
	}

	items := fsutil.List(c.dir, true)
	fsutil.ListToAbsolute(c.dir, items)

	for _, item := range items {
		os.Remove(item)
	}

	return true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Validate validates cache configuration
func (c Config) Validate() error {
	if c.DefaultExpiration < MIN_EXPIRATION {
		return fmt.Errorf("Expiration is too short (< 1s)")
	}

	if c.CleanupInterval != 0 && c.CleanupInterval < MIN_CLEANUP_INTERVAL {
		return fmt.Errorf("Cleanup interval is too short (< 1s)")
	}

	err := fsutil.ValidatePerms("DRWX", c.Dir)

	if err != nil {
		return fmt.Errorf("Can't use given directory for cache: %w", err)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getItemPath returns path to cache item
func (c *Cache) getItemPath(key string, temporary bool) string {
	if temporary {
		return path.Join(c.dir, "."+c.hashKey(key))
	}

	return path.Join(c.dir, c.hashKey(key))
}

// hashKey generates SHA-1 hash for given key
func (c *Cache) hashKey(key string) string {
	return fmt.Sprintf("%64x", c.hasher.Sum([]byte(key)))
}

// janitor is cache cleanup job
func (c *Cache) janitor(interval time.Duration) {
	for range time.NewTicker(interval).C {
		now := time.Now()

		items := fsutil.List(c.dir, true)
		fsutil.ListToAbsolute(c.dir, items)

		for _, item := range items {
			mtime, _ := fsutil.GetMTime(item)

			if mtime.Before(now) {
				os.Remove(item)
			}
		}
	}
}
