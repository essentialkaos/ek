//go:build !windows
// +build !windows

// Package fs provides cache with file system storage
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"encoding/gob"
	"fmt"
	"os"
	"path"
	"regexp"
	"time"

	"github.com/essentialkaos/ek/v13/cache"
	"github.com/essentialkaos/ek/v13/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DEFAULT_EXPIRATION is default expiration
const DEFAULT_EXPIRATION = cache.HOUR

// DEFAULT_VALIDATION_REGEXP is default regular expression pattern to validate keys
const DEFAULT_VALIDATION_REGEXP = `^[a-zA-Z0-9_-]{1,}$`

// MIN_EXPIRATION is minimal expiration duration
const MIN_EXPIRATION = cache.SECOND

// MIN_CLEANUP_INTERVAL is minimal cleanup interval
const MIN_CLEANUP_INTERVAL = cache.SECOND

// ////////////////////////////////////////////////////////////////////////////////// //

// Cache is fs cache instance
type Cache struct {
	dir             string
	expiration      cache.Duration
	validationRegex *regexp.Regexp
	isJanitorWorks  bool
}

// Config is cache configuration
type Config struct {
	Dir               string
	ValidationRegexp  string
	DefaultExpiration cache.Duration
	CleanupInterval   cache.Duration
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
		expiration: DEFAULT_EXPIRATION,
	}

	if config.DefaultExpiration != 0 {
		c.expiration = config.DefaultExpiration
	}

	if config.ValidationRegexp != "" {
		c.validationRegex = regexp.MustCompile(config.ValidationRegexp)
	} else {
		c.validationRegex = regexp.MustCompile(DEFAULT_VALIDATION_REGEXP)
	}

	if config.CleanupInterval != 0 {
		c.isJanitorWorks = true
		go c.janitor(config.CleanupInterval)
	}

	return c, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

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

// Has returns true if cache contains data for given key
func (c *Cache) Has(key string) bool {
	if c == nil || !c.isValidKey(key) {
		return false
	}

	return fsutil.IsExist(c.getItemPath(key, false))
}

// Set adds or updates item in cache
func (c *Cache) Set(key string, data any, expiration ...cache.Duration) bool {
	if c == nil || data == nil || !c.isValidKey(key) {
		return false
	}

	tmpFile := c.getItemPath(key, true)

	if !writeItem(tmpFile, data) {
		return false
	}

	itemFile := c.getItemPath(key, false)

	if os.Rename(tmpFile, itemFile) != nil {
		os.Remove(tmpFile)
		return false
	}

	expr := c.expiration

	if len(expiration) > 0 && expiration[0] >= MIN_EXPIRATION {
		expr = expiration[0]
	}

	return os.Chtimes(itemFile, time.Time{}, time.Now().Add(expr)) == nil
}

// GetWithExpiration returns item from cache
func (c *Cache) Get(key string) any {
	if c == nil || !c.isValidKey(key) || !c.Has(key) {
		return nil
	}

	expr := c.GetExpiration(key)

	if !expr.IsZero() && expr.Before(time.Now()) {
		c.Delete(key)
		return nil
	}

	return readItem(c.getItemPath(key, false))
}

// GetWithExpiration returns item expiration date
func (c *Cache) GetExpiration(key string) time.Time {
	if c == nil || !c.isValidKey(key) || !c.Has(key) {
		return time.Time{}
	}

	mt, _ := fsutil.GetMTime(c.getItemPath(key, false))

	return mt
}

// GetWithExpiration returns item from cache and expiration date or nil
func (c *Cache) GetWithExpiration(key string) (any, time.Time) {
	if c == nil || !c.isValidKey(key) || !c.Has(key) {
		return nil, time.Time{}
	}

	data := c.Get(key)

	if data != nil {
		return data, c.GetExpiration(key)
	}

	return nil, time.Time{}
}

// Keys is an iterator over cache keys
func (c *Cache) Keys(yield func(k string) bool) {
	if c == nil {
		return
	}

	for _, k := range fsutil.List(c.dir, true) {
		if !yield(k) {
			return
		}
	}
}

// All is an iterator over cache items
func (c *Cache) All(yield func(k string, v any) bool) {
	if c == nil {
		return
	}

	for _, k := range fsutil.List(c.dir, true) {
		if !yield(k, c.Get(k)) {
			return
		}
	}
}

// Delete removes item from cache
func (c *Cache) Delete(key string) bool {
	if c == nil || !c.isValidKey(key) {
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
	if c.DefaultExpiration != 0 && c.DefaultExpiration < MIN_EXPIRATION {
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

// isValidKey returns true if key is valid
func (c *Cache) isValidKey(key string) bool {
	return key != "" && c.validationRegex.MatchString(key)
}

// getItemPath returns path to cache item
func (c *Cache) getItemPath(key string, temporary bool) string {
	if temporary {
		return path.Join(c.dir, "."+key)
	}

	return path.Join(c.dir, key)
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

// ////////////////////////////////////////////////////////////////////////////////// //

// writeItem encodes data into GOB format and writes it into the file
func writeItem(file string, data any) bool {
	fd, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)

	if err != nil {
		return false
	}

	err = gob.NewEncoder(fd).Encode(&cacheItem{data})

	fd.Close()

	if err != nil {
		os.Remove(file)
	}

	return err == nil
}

// readItem reads GOB-encoded data from the file
func readItem(file string) any {
	fd, err := os.Open(file)

	if err != nil {
		return nil
	}

	item := &cacheItem{}
	err = gob.NewDecoder(fd).Decode(item)

	fd.Close()

	if err != nil {
		return nil
	}

	return item.Data
}
