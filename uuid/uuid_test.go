package uuid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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

func (s *UUIDSuite) TestGenUUID4(c *C) {
	c.Assert(UUID4(), HasLen, 16)
	c.Assert(UUID4().String(), Not(Equals), "00000000-0000-0000-0000-000000000000")
}

func (s *UUIDSuite) TestGenUUID5(c *C) {
	c.Assert(UUID5(NsURL, "TEST"), HasLen, 16)
	c.Assert(UUID5(NsURL, "TEST").String(), Not(Equals), "00000000-0000-0000-0000-000000000000")
}

func (s *UUIDSuite) TestDeprecated(c *C) {
	c.Assert(GenUUID(), Not(Equals), "00000000-0000-0000-0000-000000000000")
	c.Assert(GenUUID4(), Not(Equals), "00000000-0000-0000-0000-000000000000")
	c.Assert(GenUUID5(NsURL, "TEST"), Not(Equals), "00000000-0000-0000-0000-000000000000")
}

func (s *UUIDSuite) BenchmarkGenUUID4(c *C) {
	for i := 0; i < c.N; i++ {
		GenUUID4()
	}
}

func (s *UUIDSuite) BenchmarkGenUUID5(c *C) {
	for i := 0; i < c.N; i++ {
		GenUUID5(NsURL, "TEST")
	}
}
