package luhn

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type LuhnSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&LuhnSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *LuhnSuite) TestIsValid(c *C) {
	c.Assert(IsValid(""), Equals, false)
	c.Assert(IsValid("12345B"), Equals, false)
	c.Assert(IsValid("1"), Equals, false)
	c.Assert(IsValid("8970199240502489916"), Equals, true)
	c.Assert(IsValid("4513349830885029"), Equals, true)
}

func (s *LuhnSuite) TestCalculate(c *C) {
	_, err := Calculate("")
	c.Assert(err, Equals, ErrEmptyData)

	_, err = Calculate("12345B")
	c.Assert(err, Equals, ErrInvalidData)

	_, err = Calculate("1")
	c.Assert(err, Equals, ErrInvalidData)

	crc, err := Calculate("8970199240502489916")
	c.Assert(err, IsNil)
	c.Assert(crc, Equals, "6")

	crc, err = Calculate("897019924050248991")
	c.Assert(err, IsNil)
	c.Assert(crc, Equals, "6")
}

func (s *LuhnSuite) TestNormalize(c *C) {
	_, err := Normalize("")
	c.Assert(err, Equals, ErrEmptyData)

	_, err = Normalize("12345B")
	c.Assert(err, Equals, ErrInvalidData)

	_, err = Normalize("1")
	c.Assert(err, Equals, ErrInvalidData)

	v, err := Normalize("8970199240502489916")
	c.Assert(err, IsNil)
	c.Assert(v, Equals, "8970199240502489916")

	v, err = Normalize("897019924050248991")
	c.Assert(err, IsNil)
	c.Assert(v, Equals, "8970199240502489916")
}
