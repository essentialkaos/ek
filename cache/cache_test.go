package cache

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
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
	cache := New(time.Second/16, time.Second/32)

	cache.Set("1", "TEST")
	cache.Set("2", "TEST")

	c.Assert(cache.Size(), Equals, 2)

	c.Assert(cache.Get("1"), Equals, "TEST")
	c.Assert(cache.Get("2"), Equals, "TEST")

	c.Assert(cache.Has("2"), Equals, true)
	c.Assert(cache.Has("3"), Equals, false)

	item, exp := cache.GetWithExpiration("1")

	c.Assert(item, Equals, "TEST")
	c.Assert(exp.IsZero(), Not(Equals), true)

	cache.Delete("1")

	c.Assert(cache.Get("1"), Equals, nil)

	time.Sleep(time.Second / 8)

	item, _ = cache.GetWithExpiration("2")

	c.Assert(cache.Get("2"), Equals, nil)
	c.Assert(item, Equals, nil)

	cache.Flush()
}

func (s *CacheSuite) TestCacheWithoutJanitor(c *C) {
	cache := New(time.Second/32, 0)

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
	cache := New(time.Second/16, time.Minute)

	cache.Set("1", "TEST")

	time.Sleep(time.Second / 8)

	item, _ := cache.GetWithExpiration("1")
	c.Assert(item, Equals, nil)

	c.Assert(cache.Get("1"), Equals, nil)
}

func (s *CacheSuite) TestNil(c *C) {
	var cache *Cache

	cache.Set("1", "TEST")
	cache.Delete("1")
	cache.Flush()

	c.Assert(cache.Size(), Equals, 0)
	c.Assert(cache.Expired(), Equals, 0)
	c.Assert(cache.Get("1"), Equals, nil)
	c.Assert(cache.Has("1"), Equals, false)

	item, exp := cache.GetWithExpiration("1")
	c.Assert(item, Equals, nil)
	c.Assert(exp.IsZero(), Equals, true)
}
