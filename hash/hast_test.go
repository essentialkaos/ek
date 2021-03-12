package hash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type HashSuite struct {
	TmpDir string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&HashSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *HashSuite) SetUpSuite(c *C) {
	s.TmpDir = c.MkDir()
}

func (s *HashSuite) TestJCHash(c *C) {
	c.Assert(JCHash(1, 1), Equals, 0)
	c.Assert(JCHash(36, 49), Equals, 8)
	c.Assert(JCHash(0xDEAD10CC, 1), Equals, 0)
	c.Assert(JCHash(0xDEAD10CC, 1000), Equals, 361)
	c.Assert(JCHash(128, 1024), Equals, 267)
}

func (s *HashSuite) TestJCHashNegative(c *C) {
	c.Assert(JCHash(0, -10), Equals, 0)
	c.Assert(JCHash(0xDEAD10CC, -1000), Equals, 0)
}

func (s *HashSuite) TestFileHash(c *C) {
	tempFile := s.TmpDir + "/test.log"

	err := ioutil.WriteFile(tempFile, []byte("ABCDEF12345\n\n"), 0644)

	hash1 := FileHash(tempFile)
	hash2 := FileHash(s.TmpDir + "/not-exist.log")

	c.Assert(err, IsNil)

	c.Assert(hash1, Equals, "2d7ec20906125cd23fee7b628b98463d554b1105b141b2d39a19bac5f3274dec")
	c.Assert(hash2, Equals, "")
}
