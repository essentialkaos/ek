package cache

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"
	"time"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type CacheSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&CacheSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *CacheSuite) TestStore(c *C) {
	store := New(time.Second/16, time.Second/32)

	store.Set("1", "TEST")
	store.Set("2", "TEST")

	c.Assert(store.Get("1"), Equals, "TEST")
	c.Assert(store.Get("2"), Equals, "TEST")

	c.Assert(store.Has("2"), Equals, true)
	c.Assert(store.Has("3"), Equals, false)

	item, exp := store.GetWithExpiration("1")

	c.Assert(item, Equals, "TEST")
	c.Assert(exp.IsZero(), Not(Equals), true)

	store.Delete("1")

	c.Assert(store.Get("1"), Equals, nil)

	time.Sleep(time.Second / 8)

	item, _ = store.GetWithExpiration("2")

	c.Assert(store.Get("2"), Equals, nil)
	c.Assert(item, Equals, nil)

	store.Flush()
}

func (s *CacheSuite) TestStoreWithoutJanitor(c *C) {
	store := New(time.Second/32, 0)

	store.Set("1", "TEST")
	store.Set("2", "TEST")
	store.Set("3", "TEST")

	c.Assert(store.Has("2"), Equals, true)
	c.Assert(store.Has("0"), Equals, false)

	time.Sleep(time.Second / 16)

	c.Assert(store.Has("1"), Equals, false)
	c.Assert(store.Get("2"), Equals, nil)

	data, _ := store.GetWithExpiration("3")

	c.Assert(data, Equals, nil)
}

func (s *CacheSuite) TestExpiration(c *C) {
	store := New(time.Second/16, time.Minute)

	store.Set("1", "TEST")

	time.Sleep(time.Second / 8)

	item, _ := store.GetWithExpiration("1")
	c.Assert(item, Equals, nil)

	c.Assert(store.Get("1"), Equals, nil)
}

func (s *CacheSuite) TestNil(c *C) {
	var store *Store

	store.Set("1", "TEST")
	store.Delete("1")
	store.Flush()

	c.Assert(store.Get("1"), Equals, nil)
	c.Assert(store.Has("1"), Equals, false)

	item, exp := store.GetWithExpiration("1")
	c.Assert(item, Equals, nil)
	c.Assert(exp.IsZero(), Equals, true)
}
