package prefixed

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"

	"github.com/essentialkaos/ek/v13/uuid"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type PrefixedSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PrefixedSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PrefixedSuite) TestEncode(c *C) {
	var u uuid.UUID

	c.Assert(Encode("", u), Equals, "")
	c.Assert(Encode("test", u), Equals, "")

	u = uuid.UUID7()

	c.Assert(Encode("test", u), Not(Equals), "")
}

func (s *PrefixedSuite) TestDecode(c *C) {
	_, _, err := Decode("")
	c.Assert(err, Equals, ErrNoPrefix)
	_, _, err = Decode("test")
	c.Assert(err, Equals, ErrNoPrefix)
	_, _, err = Decode(".ABCD")
	c.Assert(err, Equals, ErrEmptyPrefix)
	_, _, err = Decode("test.")
	c.Assert(err, Equals, ErrEmptyUUID)
	_, _, err = Decode("test.####")
	c.Assert(err, ErrorMatches, `can't decode UUID data: illegal base64 data at input byte 0`)

	uu := uuid.UUID7()
	pu := Encode("test", uu)
	prf, u, err := Decode(pu)

	c.Assert(prf, Equals, "test")
	c.Assert(u.String(), Equals, uu.String())
	c.Assert(err, IsNil)
}
