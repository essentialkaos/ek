package strutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type StrUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&StrUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *StrUtilSuite) TestConcat(c *C) {
	s1 := "abcdef"
	s2 := "123456"
	s3 := "ABCDEF"
	s4 := "!@#$%^"

	c.Assert(Concat(s1, s2, s3, s4), Equals, s1+s2+s3+s4)
}

func (s *StrUtilSuite) TestSubstr(c *C) {
	c.Assert(Substr("", 1, 2), Equals, "")
	c.Assert(Substr("test1234TEST", 30, 32), Equals, "")
	c.Assert(Substr("test1234TEST", 0, 8), Equals, "test1234")
	c.Assert(Substr("test1234TEST", 4, 8), Equals, "1234")
	c.Assert(Substr("test1234TEST", 8, 16), Equals, "TEST")
	c.Assert(Substr("test1234TEST", -1, 4), Equals, "test")

	c.Assert(Substr("test"+string(rune(65533))+"1234TEST", 0, 8), Equals, "test1234")
}

func (s *StrUtilSuite) TestEllipsis(c *C) {
	c.Assert(Ellipsis("Test1234", 8), Equals, "Test1234")
	c.Assert(Ellipsis("Test1234test", 8), Equals, "Test1...")
}

func (s *StrUtilSuite) TestHead(c *C) {
	c.Assert(Head("", 1), Equals, "")
	c.Assert(Head("ABCD1234", 0), Equals, "")
	c.Assert(Head("ABCD1234", -10), Equals, "")
	c.Assert(Head("ABCD1234", 1), Equals, "A")
	c.Assert(Head("ABCD1234", 4), Equals, "ABCD")
	c.Assert(Head("ABCD1234", 100), Equals, "ABCD1234")
}

func (s *StrUtilSuite) TestTail(c *C) {
	c.Assert(Tail("", 1), Equals, "")
	c.Assert(Tail("ABCD1234", 0), Equals, "")
	c.Assert(Tail("ABCD1234", -10), Equals, "")
	c.Assert(Tail("ABCD1234", 1), Equals, "4")
	c.Assert(Tail("ABCD1234", 4), Equals, "1234")
	c.Assert(Tail("ABCD1234", 100), Equals, "ABCD1234")
}

func (s *StrUtilSuite) TestSize(c *C) {
	c.Assert(PrefixSize("", ' '), Equals, 0)
	c.Assert(PrefixSize("abcd", ' '), Equals, 0)
	c.Assert(PrefixSize("    abcd", ' '), Equals, 4)
	c.Assert(PrefixSize("    ", ' '), Equals, 4)

	c.Assert(SuffixSize("", ' '), Equals, 0)
	c.Assert(SuffixSize("abcd", ' '), Equals, 0)
	c.Assert(SuffixSize("abcd    ", ' '), Equals, 4)
	c.Assert(SuffixSize("    ", ' '), Equals, 4)
}
