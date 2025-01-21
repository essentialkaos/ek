package protip

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type TipSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TipSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *TipSuite) TestAdd(c *C) {
	c.Assert(collection, IsNil)

	var tip *Tip
	Add(tip)
	c.Assert(collection, NotNil)
	c.Assert(collection.data, HasLen, 0)

	tip = &Tip{}
	Add(tip)
	c.Assert(collection.data, HasLen, 0)

	tip.Title = "TEST"
	Add(tip)
	c.Assert(collection.data, HasLen, 0)

	tip.Message = "Test message"

	disabled = true

	Add(tip)
	c.Assert(collection.data, HasLen, 0)

	disabled = false

	Add(tip)
	c.Assert(collection.data, HasLen, 1)
	c.Assert(collection.data[0].Weight, Equals, 0.5)
}

func (s *TipSuite) TestShow(c *C) {
	collection = nil

	c.Assert(Show(true), Equals, false)

	Add(&Tip{Title: "Test #1", Message: "Test", Weight: 0.99, ColorTag: "{r}"})
	Add(&Tip{Title: "Test #2", Message: "Test", Weight: 0.01})

	Probability = 0.0

	c.Assert(Show(false), Equals, false)

	Probability = 1.0

	c.Assert(Show(true), Equals, true)
}

func (s *TipSuite) TestIntSearch(c *C) {
	k := []int{1, 3, 5}

	c.Assert(searchInts(k, 1), Equals, 0)
	c.Assert(searchInts(k, 3), Equals, 1)
	c.Assert(searchInts(k, 5), Equals, 2)
}
