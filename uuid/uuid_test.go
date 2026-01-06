package uuid

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

type UUIDSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&UUIDSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *UUIDSuite) TestUUID4(c *C) {
	c.Assert(UUID4(), HasLen, 16)
	c.Assert(UUID4().String(), Not(Equals), "00000000-0000-0000-0000-000000000000")
	c.Assert(UUID4().IsZero(), Equals, false)
}

func (s *UUIDSuite) TestUUID5(c *C) {
	c.Assert(UUID5(NsURL, "TEST"), HasLen, 16)
	c.Assert(UUID5(NsURL, "TEST").String(), Not(Equals), "00000000-0000-0000-0000-000000000000")
	c.Assert(UUID5(NsURL, "TEST").IsZero(), Equals, false)
}

func (s *UUIDSuite) TestUUID7(c *C) {
	c.Assert(UUID7(), HasLen, 16)
	c.Assert(UUID7().String(), Not(Equals), "00000000-0000-0000-0000-000000000000")
	c.Assert(UUID7().IsZero(), Equals, false)
}

func (s *UUIDSuite) BenchmarkUUID4(c *C) {
	for range c.N {
		UUID4()
	}
}

func (s *UUIDSuite) BenchmarkUUID5(c *C) {
	for range c.N {
		UUID5(NsURL, "TEST")
	}
}
