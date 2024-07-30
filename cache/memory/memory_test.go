package memory

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type CacheSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&CacheSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *CacheSuite) TestCache(c *C) {
	cache, err := New(Config{
		DefaultExpiration: time.Second / 16,
		CleanupInterval:   time.Second / 32,
	})

	c.Assert(err, IsNil)
	c.Assert(cache, NotNil)

	cache.Set("1", "TEST")
	cache.Set("2", "TEST")
	cache.Set("3", "TEST", time.Minute)

	c.Assert(cache.Size(), Equals, 3)

	c.Assert(cache.Get("1"), Equals, "TEST")
	c.Assert(cache.Get("2"), Equals, "TEST")

	c.Assert(cache.Has("2"), Equals, true)
	c.Assert(cache.Has("4"), Equals, false)

	item, exp := cache.GetWithExpiration("1")

	c.Assert(item, Equals, "TEST")
	c.Assert(exp.IsZero(), Not(Equals), true)

	cache.Delete("1")

	c.Assert(cache.Get("1"), Equals, nil)
	c.Assert(cache.GetExpiration("3").IsZero(), Not(Equals), true)

	time.Sleep(time.Second / 8)

	item, _ = cache.GetWithExpiration("2")

	c.Assert(cache.Get("2"), Equals, nil)
	c.Assert(item, Equals, nil)

	c.Assert(cache.Flush(), Equals, true)
}

func (s *CacheSuite) TestCacheWithoutJanitor(c *C) {
	cache, err := New(Config{DefaultExpiration: time.Second / 32})

	c.Assert(err, IsNil)
	c.Assert(cache, NotNil)

	cache.Set("1", "TEST")
	cache.Set("2", "TEST")
	cache.Set("3", "TEST")

	c.Assert(cache.Has("2"), Equals, true)
	c.Assert(cache.Has("0"), Equals, false)

	time.Sleep(time.Second / 16)

	c.Assert(cache.Expired(), Equals, 3)
	c.Assert(cache.Has("1"), Equals, false)
	c.Assert(cache.Get("2"), Equals, nil)

	data, _ := cache.GetWithExpiration("3")

	c.Assert(data, Equals, nil)
}

func (s *CacheSuite) TestExpiration(c *C) {
	cache, err := New(Config{
		DefaultExpiration: time.Second / 16,
		CleanupInterval:   time.Minute,
	})

	c.Assert(err, IsNil)
	c.Assert(cache, NotNil)

	cache.Set("1", "TEST")

	time.Sleep(time.Second / 8)

	item, _ := cache.GetWithExpiration("1")
	c.Assert(item, Equals, nil)

	c.Assert(cache.GetExpiration("2").IsZero(), Equals, true)

	c.Assert(cache.Get("1"), Equals, nil)
}

func (s *CacheSuite) TestNil(c *C) {
	var cache *Cache

	c.Assert(func() { cache.Set("1", "TEST") }, NotPanics)
	c.Assert(func() { cache.Delete("1") }, NotPanics)
	c.Assert(func() { cache.Flush() }, NotPanics)

	c.Assert(cache.Size(), Equals, 0)
	c.Assert(cache.Expired(), Equals, 0)
	c.Assert(cache.Get("1"), Equals, nil)
	c.Assert(cache.Has("1"), Equals, false)
	c.Assert(cache.GetExpiration("1").IsZero(), Equals, true)

	item, exp := cache.GetWithExpiration("1")
	c.Assert(item, Equals, nil)
	c.Assert(exp.IsZero(), Equals, true)
}

func (s *CacheSuite) TestConfig(c *C) {
	_, err := New(Config{DefaultExpiration: 1})

	c.Assert(err.Error(), Equals, "Invalid configuration: Expiration is too short (< 1ms)")

	_, err = New(Config{DefaultExpiration: time.Minute, CleanupInterval: 1})

	c.Assert(err.Error(), Equals, "Invalid configuration: Cleanup interval is too short (< 1ms)")
}
