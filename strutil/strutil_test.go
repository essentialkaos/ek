package strutil

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
	c.Assert(Substr("test1234TEST", 0, 999), Equals, "test1234TEST")
	c.Assert(Substr("test1234TEST", 4, 8), Equals, "1234")
	c.Assert(Substr("test1234TEST", 8, 16), Equals, "TEST")
	c.Assert(Substr("test1234TEST", -1, 4), Equals, "test")
}

func (s *StrUtilSuite) BenchmarkSubstr(c *C) {
	for i := 0; i < c.N; i++ {
		Substr("test1234TEST", 4, 8)
	}
}

func (s *StrUtilSuite) TestEllipsis(c *C) {
	c.Assert(Ellipsis("Test1234", 8), Equals, "Test1234")
	c.Assert(Ellipsis("Test1234test", 8), Equals, "Test1...")
}

func (s *StrUtilSuite) BenchmarkEllipsis(c *C) {
	for i := 0; i < c.N; i++ {
		Ellipsis("Test1234test", 8)
	}
}

func (s *StrUtilSuite) TestHead(c *C) {
	c.Assert(Head("", 1), Equals, "")
	c.Assert(Head("ABCD1234", 0), Equals, "")
	c.Assert(Head("ABCD1234", -10), Equals, "")
	c.Assert(Head("ABCD1234", 1), Equals, "A")
	c.Assert(Head("ABCD1234", 4), Equals, "ABCD")
	c.Assert(Head("ABCD1234", 100), Equals, "ABCD1234")
}

func (s *StrUtilSuite) BenchmarkHead(c *C) {
	for i := 0; i < c.N; i++ {
		Head("ABCD1234ABCD1234", 4)
	}
}

func (s *StrUtilSuite) TestTail(c *C) {
	c.Assert(Tail("", 1), Equals, "")
	c.Assert(Tail("ABCD1234", 0), Equals, "")
	c.Assert(Tail("ABCD1234", -10), Equals, "")
	c.Assert(Tail("ABCD1234", 1), Equals, "4")
	c.Assert(Tail("ABCD1234", 4), Equals, "1234")
	c.Assert(Tail("ABCD1234", 100), Equals, "ABCD1234")
}

func (s *StrUtilSuite) BenchmarkTail(c *C) {
	for i := 0; i < c.N; i++ {
		Tail("ABCD1234ABCD1234", 4)
	}
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

func (s *StrUtilSuite) BenchmarkSize(c *C) {
	for i := 0; i < c.N; i++ {
		PrefixSize("    abcd", ' ')
	}
}
func (s *StrUtilSuite) TestReplaceAll(c *C) {
	c.Assert(ReplaceAll("ABCDABCD12341234", "AB12", "?"), Equals, "??CD??CD??34??34")
	c.Assert(ReplaceAll("", "AB12", "?"), Equals, "")
}

func (s *StrUtilSuite) TestFields(c *C) {
	c.Assert(Fields(""), IsNil)
	c.Assert(Fields(""), HasLen, 0)
	c.Assert(Fields("1 2 3 4 5"), DeepEquals, []string{"1", "2", "3", "4", "5"})
	c.Assert(Fields("1,2,3,4,5"), DeepEquals, []string{"1", "2", "3", "4", "5"})
	c.Assert(Fields("1,  2, 3,   4, 5"), DeepEquals, []string{"1", "2", "3", "4", "5"})
	c.Assert(Fields("\"1 2\" 3 \"4 5\""), DeepEquals, []string{"1 2", "3", "4 5"})
	c.Assert(Fields("'1 2' 3 '4 5'"), DeepEquals, []string{"1 2", "3", "4 5"})
}

func (s *StrUtilSuite) TestReadField(c *C) {
	c.Assert(ReadField("abc 1234 DEF", -1), Equals, "")
	c.Assert(ReadField("abc 1234 DEF", 0), Equals, "abc")
	c.Assert(ReadField("abc 1234 DEF", 1), Equals, "1234")
	c.Assert(ReadField("abc 1234 DEF", 2), Equals, "DEF")
	c.Assert(ReadField("abc 1234 DEF", 3), Equals, "")

	c.Assert(ReadField("abc|1234|DEF", 1, "|"), Equals, "1234")
	c.Assert(ReadField("abc+1234|DEF", 1, "|+"), Equals, "1234")
}

func (s *StrUtilSuite) BenchmarkFields(c *C) {
	for i := 0; i < c.N; i++ {
		Fields("\"1 2\" 3 \"4 5\"")
	}
}

func (s *StrUtilSuite) BenchmarkReadField(c *C) {
	for i := 0; i < c.N; i++ {
		ReadField("abc 1234 DEF", 2)
	}
}

func (s *StrUtilSuite) TestLen(c *C) {
	c.Assert(Len("ABCDABCD12341234"), Equals, 16)
	c.Assert(Len(""), Equals, 0)
	c.Assert(Len("✶✈12AB例例子예"), Equals, 10)
}

func (s *StrUtilSuite) BenchmarkLen(c *C) {
	for i := 0; i < c.N; i++ {
		Len("✶✈12AB例例子예")
	}
}
