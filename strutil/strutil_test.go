package strutil

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

type StrUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&StrUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *StrUtilSuite) TestHelpers(c *C) {
	c.Assert(Q(), Equals, "")
	c.Assert(Q("a"), Equals, "a")
	c.Assert(Q("", "", "1"), Equals, "1")

	c.Assert(B(true, "A", "B"), Equals, "A")
	c.Assert(B(false, "A", "B"), Equals, "B")
}

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
	c.Assert(Substr("test1234TEST", 4, 4), Equals, "1234")
	c.Assert(Substr("test1234TEST", 8, 16), Equals, "TEST")
	c.Assert(Substr("test1234TEST", -1, 4), Equals, "test")
	c.Assert(Substr("简单的消息", -1, 2), Equals, "简单")
	c.Assert(Substr("Пример", -1, 2), Equals, "Пр")
	c.Assert(Substr("test1234TEST", 12, 16), Equals, "")
	c.Assert(Substr("test1234TEST", 11, 16), Equals, "T")
}

func (s *StrUtilSuite) TestSubstring(c *C) {
	c.Assert(Substring("", 1, 2), Equals, "")
	c.Assert(Substring("test1234TEST", 30, 32), Equals, "")
	c.Assert(Substring("test1234TEST", 0, 999), Equals, "test1234TEST")
	c.Assert(Substring("test1234TEST", 4, 8), Equals, "1234")
	c.Assert(Substring("test1234TEST", 4, 4), Equals, "")
	c.Assert(Substring("test1234TEST", 8, 100), Equals, "TEST")
	c.Assert(Substring("test1234TEST", 2, -6), Equals, "st12")
	c.Assert(Substring("简单的消息", -1, 2), Equals, "简单")
	c.Assert(Substring("Пример", -1, 2), Equals, "Пр")
	c.Assert(Substring("test1234TEST", 12, 99), Equals, "")
	c.Assert(Substring("test1234TEST", 11, 99), Equals, "T")
}

func (s *StrUtilSuite) TestExtract(c *C) {
	c.Assert(Extract("", 1, 10), Equals, "")
	c.Assert(Extract("test1234TEST", -10, 4), Equals, "test")
	c.Assert(Extract("test1234TEST", 8, 100), Equals, "TEST")
	c.Assert(Extract("test1234TEST", 4, 8), Equals, "1234")
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

func (s *StrUtilSuite) TestExclude(c *C) {
	c.Assert(Exclude("ABCD1234abcd1234ABCD", ""), Equals, "ABCD1234abcd1234ABCD")
	c.Assert(Exclude("", "1234"), Equals, "")
	c.Assert(Exclude("ABCD1234abcd1234ABCD", "5678"), Equals, "ABCD1234abcd1234ABCD")
	c.Assert(Exclude("ABCD1234abcd1234ABCD", "1234"), Equals, "ABCDabcdABCD")
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

func (s *StrUtilSuite) TestReplaceAll(c *C) {
	c.Assert(ReplaceAll("ABCDABCD12341234", "AB12", "?"), Equals, "??CD??CD??34??34")
	c.Assert(ReplaceAll("", "AB12", "?"), Equals, "")
}

func (s *StrUtilSuite) TestReplaceIgnoreCase(c *C) {
	c.Assert(ReplaceIgnoreCase("ABCD1234abcd1234AbCd11ABcd", "abcd", "????"), Equals, "????1234????1234????11????")
	c.Assert(ReplaceIgnoreCase("TESTtestTEST", "abcd", "????"), Equals, "TESTtestTEST")
	c.Assert(ReplaceIgnoreCase("", "abcd", "????"), Equals, "")
	c.Assert(ReplaceIgnoreCase("ABCD1234abcd1234AbCd11ABcd", "abcd", ""), Equals, "1234123411")
}

func (s *StrUtilSuite) TestFields(c *C) {
	c.Assert(Fields(""), IsNil)
	c.Assert(Fields(""), HasLen, 0)
	c.Assert(Fields("1 2 3 4 5"), DeepEquals, []string{"1", "2", "3", "4", "5"})
	c.Assert(Fields(`abc 123 'k:"42"'`), DeepEquals, []string{"abc", "123", `k:"42"`})
	c.Assert(Fields("1,2,3,4,5"), DeepEquals, []string{"1", "2", "3", "4", "5"})
	c.Assert(Fields("1;2;3;4;5"), DeepEquals, []string{"1", "2", "3", "4", "5"})
	c.Assert(Fields("1,  2, 3,   4, 5"), DeepEquals, []string{"1", "2", "3", "4", "5"})
	c.Assert(Fields(`"1 2" 3 "4 5"`), DeepEquals, []string{"1 2", "3", "4 5"})
	c.Assert(Fields("'1 2' 3 '4 5'"), DeepEquals, []string{"1 2", "3", "4 5"})
	c.Assert(Fields("‘1 2’ 3 ‘4 5’"), DeepEquals, []string{"1 2", "3", "4 5"})
	c.Assert(Fields("“1 2” 3 “4 5”"), DeepEquals, []string{"1 2", "3", "4 5"})
	c.Assert(Fields("„1 2“ 3 «4 5»"), DeepEquals, []string{"1 2", "3", "4 5"})
	c.Assert(Fields("«1 '2'» 3 «4 “5”»"), DeepEquals, []string{"1 '2'", "3", "4 “5”"})
	c.Assert(Fields(`Bob  Alice, 'Mary Key', "John \"Big John\" 'Dow'"`), DeepEquals, []string{"Bob", "Alice", "Mary Key", "John \"Big John\" 'Dow'"})
}

func (s *StrUtilSuite) TestReadField(c *C) {
	c.Assert(ReadField("abc 1234 DEF", -1, true), Equals, "")
	c.Assert(ReadField("abc 1234 DEF", 0, true), Equals, "abc")
	c.Assert(ReadField("abc 1234 DEF", 1, true), Equals, "1234")
	c.Assert(ReadField("abc 1234 DEF", 2, true), Equals, "DEF")
	c.Assert(ReadField("abc 1234 DEF", 3, true), Equals, "")

	c.Assert(ReadField("abc|||||1234||DEF", 1, true, '|'), Equals, "1234")
	c.Assert(ReadField("abc+1234|DEF", 1, true, '|', '+'), Equals, "1234")
	c.Assert(ReadField("abc::1234:::DEF:", 5, false, ':'), Equals, "DEF")
}

func (s *StrUtilSuite) TestCopy(c *C) {
	c.Assert(Copy(""), Equals, "")
	c.Assert(Copy("ABCD1234"), Equals, "ABCD1234")
}

func (s *StrUtilSuite) TestLen(c *C) {
	c.Assert(Len("ABCDABCD12341234"), Equals, 16)
	c.Assert(Len(""), Equals, 0)
	c.Assert(Len("✶✈12AB例例子예"), Equals, 10)
}

func (s *StrUtilSuite) TestLenVisual(c *C) {
	c.Assert(LenVisual("ABCDABCD12341234"), Equals, 16)
	c.Assert(LenVisual(""), Equals, 0)
	c.Assert(LenVisual("✶✈12AB例例子예"), Equals, 14)
}

func (s *StrUtilSuite) TestBefore(c *C) {
	c.Assert(Before("", "@"), Equals, "")
	c.Assert(Before("test::1234", "@"), Equals, "test::1234")
	c.Assert(Before("test::1234", "::"), Equals, "test")
}

func (s *StrUtilSuite) TestAfter(c *C) {
	c.Assert(After("", "@"), Equals, "")
	c.Assert(After("test::1234", "@"), Equals, "test::1234")
	c.Assert(After("test::1234", "::"), Equals, "1234")
}

func (s *StrUtilSuite) TestHasPrefixAny(c *C) {
	c.Assert(HasPrefixAny("#abcd", "#", "@"), Equals, true)
	c.Assert(HasPrefixAny("#abcd", "$", "@"), Equals, false)
}

func (s *StrUtilSuite) TestHasSuffixAny(c *C) {
	c.Assert(HasSuffixAny("abcd#", "#", "@"), Equals, true)
	c.Assert(HasSuffixAny("abcd#", "$", "@"), Equals, false)
}

func (s *StrUtilSuite) TestIndexByteSkip(c *C) {
	c.Assert(IndexByteSkip("", '/', 0), Equals, -1)
	c.Assert(IndexByteSkip("", '/', 1), Equals, -1)
	c.Assert(IndexByteSkip("/", '/', 1), Equals, -1)
	c.Assert(IndexByteSkip("/home/john/projects/test.log", '/', 2), Equals, 10)
	c.Assert(IndexByteSkip("/home/john/projects/test.log", '/', -1), Equals, 10)
}

func (s *StrUtilSuite) TestSqueezeRepeats(c *C) {
	c.Assert(SqueezeRepeats("", ""), Equals, "")
	c.Assert(SqueezeRepeats("test", ""), Equals, "test")
	c.Assert(SqueezeRepeats("1  ----  2 - 3 --     4", " -"), Equals, "1 - 2 - 3 - 4")
}

func (s *StrUtilSuite) TestMask(c *C) {
	c.Assert(Mask("", 0, 1000, '*'), Equals, "")
	c.Assert(Mask("Test1234test", 4, 8, '*'), Equals, "Test****test")
	c.Assert(Mask("Test1234test", 0, -4, '*'), Equals, "********test")
}

func (s *StrUtilSuite) TestJoinFunc(c *C) {
	c.Assert(
		JoinFunc([]string{}, " ", func(s string) string { return "" }),
		Equals, "",
	)

	c.Assert(
		JoinFunc([]string{"1"}, " ", func(s string) string { return "'" + s + "'" }),
		Equals, "'1'",
	)

	c.Assert(
		JoinFunc([]string{"1", "2", "3", "4"}, ", ", func(s string) string { return "'" + s + "'" }),
		Equals, "'1', '2', '3', '4'",
	)
}

func (s *StrUtilSuite) TestWrap(c *C) {
	c.Assert(Wrap("", 100), Equals, "")
	c.Assert(Wrap("TEST", -1), Equals, "TEST")
	c.Assert(Wrap("TEST", 4), Equals, "TEST")
	c.Assert(Wrap("TEST", 10), Equals, "TEST")
	c.Assert(Wrap("AAAABBBBCCCCDDDD", 4), Equals, "AAAA\nBBBB\nCCCC\nDDDD")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *StrUtilSuite) BenchmarkSubstr(c *C) {
	for range c.N {
		Substr("test1234TEST", 4, 8)
	}
}

func (s *StrUtilSuite) BenchmarkEllipsis(c *C) {
	for range c.N {
		Ellipsis("Test1234test", 8)
	}
}

func (s *StrUtilSuite) BenchmarkHead(c *C) {
	for range c.N {
		Head("ABCD1234ABCD1234", 4)
	}
}

func (s *StrUtilSuite) BenchmarkTail(c *C) {
	for range c.N {
		Tail("ABCD1234ABCD1234", 4)
	}
}

func (s *StrUtilSuite) BenchmarkExclude(c *C) {
	for range c.N {
		Exclude("ABCD1234abcd1234ABCD", "1234")
	}
}

func (s *StrUtilSuite) BenchmarkExtract(c *C) {
	for range c.N {
		Extract("test1234TEST", 4, 8)
	}
}

func (s *StrUtilSuite) BenchmarkSize(c *C) {
	for range c.N {
		PrefixSize("    abcd", ' ')
	}
}

func (s *StrUtilSuite) BenchmarkFields(c *C) {
	for range c.N {
		Fields(`"123" 59 "31" '2'`)
	}
}

func (s *StrUtilSuite) BenchmarkReadField(c *C) {
	for range c.N {
		ReadField("abc 1234 DEF", 2, false)
	}
}

func (s *StrUtilSuite) BenchmarkLen(c *C) {
	for range c.N {
		Len("✶✈12AB例例子예")
	}
}

func (s *StrUtilSuite) BenchmarkReplaceAll(c *C) {
	for range c.N {
		ReplaceAll("ABCDABCD12341234", "AB12", "?")
	}
}

func (s *StrUtilSuite) BenchmarkReplaceIgnoreCase(c *C) {
	for range c.N {
		ReplaceIgnoreCase("ABCD1234abcd1234AbCd11ABcd", "abcd", "????")
	}
}
