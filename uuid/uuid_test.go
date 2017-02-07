package uuid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type UUIDSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&UUIDSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *UUIDSuite) TestGenUUID(c *C) {
	c.Assert(GenUUID(), HasLen, 36)
	c.Assert(GenUUID(), Not(Equals), "00000000-0000-0000-0000-000000000000")
}

func (s *UUIDSuite) TestGenUUID4(c *C) {
	c.Assert(GenUUID4(), HasLen, 36)
	c.Assert(GenUUID4(), Not(Equals), "00000000-0000-0000-0000-000000000000")
}

func (s *UUIDSuite) TestGenUUID5(c *C) {
	c.Assert(GenUUID5(NsURL, "TEST"), HasLen, 36)
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
