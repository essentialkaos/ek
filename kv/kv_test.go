// Package kv provides simple key-value structs
package kv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type KVSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&KVSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *KVSuite) TestTypeConversion(c *C) {
	kv1 := KV{"test1", "TEST1234ABCD"}
	kv2 := KV{"test2", 1234}
	kv3 := KV{"test3", 1234.12}

	c.Assert(kv1.Value.(string), Equals, "TEST1234ABCD")
	c.Assert(kv1.String(), Equals, "TEST1234ABCD")

	c.Assert(kv2.Value.(int), Equals, 1234)
	c.Assert(kv2.Int(), Equals, 1234)

	c.Assert(kv3.Value.(float64), Equals, 1234.12)
	c.Assert(kv3.Float(), Equals, 1234.12)
}

func (s *KVSuite) TestSorting(c *C) {
	kvs := []KV{
		{"test1", "1"},
		{"test5", "2"},
		{"test3", "3"},
		{"test2", "4"},
	}

	Sort(kvs)

	c.Assert(kvs, HasLen, 4)
	c.Assert(kvs[0].Key, Equals, "test1")
	c.Assert(kvs[1].Key, Equals, "test2")
	c.Assert(kvs[2].Key, Equals, "test3")
	c.Assert(kvs[3].Key, Equals, "test5")
}
