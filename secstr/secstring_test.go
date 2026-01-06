package secstr

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

type SecstrSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SecstrSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SecstrSuite) TestSlice(c *C) {
	k := []byte("Test1234")
	ss, err := NewSecureString(k)

	c.Assert(err, IsNil)
	c.Assert(ss, NotNil)
	c.Assert(ss.IsEmpty(), Equals, false)
	c.Assert(ss.Data, HasLen, 8)
	c.Assert(k, DeepEquals, []byte{0, 0, 0, 0, 0, 0, 0, 0})
	c.Assert(string(ss.Data), Equals, "Test1234")

	err = ss.Destroy()

	c.Assert(err, IsNil)
	c.Assert(ss.Data, IsNil)

	err = ss.Destroy()

	c.Assert(err, IsNil)
	c.Assert(ss.Data, IsNil)
}

func (s *SecstrSuite) TestString(c *C) {
	k1 := "Test1234"

	ss, err := NewSecureString(k1)

	c.Assert(err, IsNil)
	c.Assert(ss, NotNil)
	c.Assert(ss.IsEmpty(), Equals, false)
	c.Assert(ss.Data, HasLen, 8)
	c.Assert(string(ss.Data), Equals, "Test1234")

	k2 := "Test1234"

	ss, err = NewSecureString(&k2)

	c.Assert(err, IsNil)
	c.Assert(ss, NotNil)
	c.Assert(ss.IsEmpty(), Equals, false)
	c.Assert(ss.Data, HasLen, 8)
	c.Assert(k2, Equals, "")
	c.Assert(string(ss.Data), Equals, "Test1234")
}

func (s *SecstrSuite) TestErrors(c *C) {
	_, err := NewSecureString(123)
	c.Assert(err, NotNil)
}

func (s *SecstrSuite) TestNil(c *C) {
	var k *String

	c.Assert(k.IsEmpty(), Equals, true)
	c.Assert(k.Destroy(), IsNil)
}
