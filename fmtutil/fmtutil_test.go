package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type FmtUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&FmtUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FmtUtilSuite) TestPretyNum(c *C) {
	c.Assert(PrettyNum(999), Equals, "999")
	c.Assert(PrettyNum(1000), Equals, "1,000")
	c.Assert(PrettyNum(1234567890), Equals, "1,234,567,890")
	c.Assert(PrettyNum(100000), Equals, "100,000")
	c.Assert(PrettyNum(100000, "|"), Equals, "100|000")
	c.Assert(PrettyNum(0), Equals, "0")
	c.Assert(PrettyNum(2500.50), Equals, "2,500.5")
	c.Assert(PrettyNum(2500.00), Equals, "2,500")
	c.Assert(PrettyNum(1.23), Equals, "1.23")
	c.Assert(PrettyNum(-1000), Equals, "-1,000")
	c.Assert(PrettyNum(-123456789), Equals, "-123,456,789")
	c.Assert(PrettyNum(-123456789.15), Equals, "-123,456,789.15")
	c.Assert(PrettyNum(math.NaN()), Equals, "0")
	c.Assert(PrettyNum("abcd"), Equals, "abcd")
}

func (s *FmtUtilSuite) TestPretyDiff(c *C) {
	c.Assert(PrettyDiff(0), Equals, "0")
	c.Assert(PrettyDiff(100), Equals, "+100")
	c.Assert(PrettyDiff(15620), Equals, "+15,620")
	c.Assert(PrettyDiff(-15620), Equals, "-15,620")
}

func (s *FmtUtilSuite) TestPretyPerc(c *C) {
	c.Assert(PrettyPerc(0.12), Equals, "0.12%")
	c.Assert(PrettyPerc(1), Equals, "1%")
	c.Assert(PrettyPerc(1.123), Equals, "1.12%")
	c.Assert(PrettyPerc(12.100), Equals, "12.1%")
	c.Assert(PrettyPerc(123.123), Equals, "123.1%")
	c.Assert(PrettyPerc(0.0002), Equals, "< 0.01%")
}

func (s *FmtUtilSuite) TestPretySize(c *C) {
	c.Assert(PrettySize(0), Equals, "0B")
	c.Assert(PrettySize(345), Equals, "345B")
	c.Assert(PrettySize(1025), Equals, "1KB")
	c.Assert(PrettySize(1025, " "), Equals, "1 KB")
	c.Assert(PrettySize(3000125), Equals, "2.86MB")
	c.Assert(PrettySize(1024*1024), Equals, "1MB")
	c.Assert(PrettySize(1024*1024*1024), Equals, "1GB")
	c.Assert(PrettySize(int64(1024*1024*1024*1024)), Equals, "1TB")
	c.Assert(PrettySize(1052500000), Equals, "1003.7MB")

	c.Assert(PrettySize(int32(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(int64(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(uint(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(uint32(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(uint64(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(float32(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(float64(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(float64(1.1)), Equals, "1B")
	c.Assert(PrettySize(float64(1.6)), Equals, "2B")
	c.Assert(PrettySize(math.NaN()), Equals, "0B")
}

func (s *FmtUtilSuite) TestPretyBool(c *C) {
	c.Assert(PrettyBool(true), Equals, "Y")
	c.Assert(PrettyBool(false), Equals, "N")
	c.Assert(PrettyBool(true, "Yes"), Equals, "Yes")
	c.Assert(PrettyBool(false, "Yes"), Equals, "N")
	c.Assert(PrettyBool(false, "Yes", "No"), Equals, "No")
}

func (s *FmtUtilSuite) TestParseSize(c *C) {
	c.Assert(ParseSize("1 MB"), Equals, uint64(1024*1024))
	c.Assert(ParseSize("1 M"), Equals, uint64(1000*1000))
	c.Assert(ParseSize("2tb"), Equals, uint64(2*1024*1024*1024*1024))
	c.Assert(ParseSize("2t"), Equals, uint64(2*1000*1000*1000*1000))
	c.Assert(ParseSize("5gB"), Equals, uint64(5*1024*1024*1024))
	c.Assert(ParseSize("5g"), Equals, uint64(5*1000*1000*1000))
	c.Assert(ParseSize("13kb"), Equals, uint64(13*1024))
	c.Assert(ParseSize("13k"), Equals, uint64(13*1000))
	c.Assert(ParseSize("512"), Equals, uint64(512))
	c.Assert(ParseSize("kb"), Equals, uint64(0))
	c.Assert(ParseSize("123!"), Equals, uint64(0))

	c.Assert(ParseSize(PrettySize(345)), Equals, uint64(345))
	c.Assert(ParseSize(PrettySize(1025)), Equals, uint64(1024))
	c.Assert(ParseSize(PrettySize(1024*1024)), Equals, uint64(1024*1024))
}

func (s *FmtUtilSuite) TestFloat(c *C) {
	c.Assert(Float(1.0), Equals, 1.0)
	c.Assert(Float(0.1), Equals, 0.1)
	c.Assert(Float(0.01), Equals, 0.01)
	c.Assert(Float(0.001), Equals, 0.0)
	c.Assert(Float(0.0001), Equals, 0.0)
	c.Assert(Float(math.NaN()), Equals, 0.0)
}

func (s *FmtUtilSuite) TestWrap(c *C) {
	input := `Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, 
eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam 
voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione 
voluptatem sequi nesciunt, cum soluta nobis est caparet.

Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci 
velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.

Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea 
commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae 
consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?`

	result := `  Sed ut perspiciatis unde omnis iste
  natus error sit voluptatem accusantium
  doloremque laudantium, totam rem
  aperiam, eaque ipsa quae ab illo
  inventore veritatis et quasi
  architecto beatae vitae dicta sunt
  explicabo. Nemo enim ipsam voluptatem
  quia voluptas sit aspernatur aut odit
  aut fugit, sed quia consequuntur magni
  dolores eos qui ratione voluptatem
  sequi nesciunt, cum soluta nobis est
  caparet.

  Neque porro quisquam est, qui dolorem
  ipsum quia dolor sit amet,
  consectetur, adipisci velit, sed quia
  non numquam eius modi tempora incidunt
  ut labore et dolore magnam aliquam
  quaerat voluptatem.

  Ut enim ad minima veniam, quis nostrum
  exercitationem ullam corporis suscipit
  laboriosam, nisi ut aliquid ex ea
  commodi consequatur? Quis autem vel
  eum iure reprehenderit qui in ea
  voluptate velit esse quam nihil
  molestiae consequatur, vel illum qui
  dolorem eum fugiat quo voluptas nulla
  pariatur?`

	c.Assert(Wrap(input, "  ", 40), Equals, result)

	input = "Sed ut \033[40;38;5;82mperspiciatis \x1b[30;48;5;82munde\x1B[0m omnis iste natus error sit voluptatem accusantium doloremque"
	result = "  Sed ut \x1b[40;38;5;82mperspiciatis \x1b[30;48;5;82munde\x1b[0m\n" +
		"  omnis iste natus error sit\n" +
		"  voluptatem accusantium\n" +
		"  doloremque"

	c.Assert(Wrap(input, "  ", 30), Equals, result)

	input = "abcd1234abcd1234abcd1234abcd1234 abcd abcd abcd 1234 1234"
	result = "  abcd1234abcd1234abcd1234abcd1234\n" +
		"  abcd abcd abcd\n" +
		"  1234 1234"

	c.Assert(Wrap(input, "  ", 20), Equals, result)
}

func (s *FmtUtilSuite) TestSeparator(c *C) {
	SeparatorSize = 1

	Separator(true)
	Separator(false)
	Separator(true, "test")
	Separator(false, "test")
	Separator(false, "TEST1234TEST1234TEST1234TEST1234TEST1234TEST1234TEST1234TEST1234TEST1234TEST1234TEST1234")

	SeparatorFullscreen = true

	Separator(true)

	c.Assert(between(0, 1, 3), Equals, 1)
	c.Assert(between(2, 1, 3), Equals, 2)
	c.Assert(between(10, 1, 3), Equals, 3)
}

func (s *FmtUtilSuite) TestCountDigits(c *C) {
	c.Assert(CountDigits(1), Equals, 1)
	c.Assert(CountDigits(999), Equals, 3)
	c.Assert(CountDigits(45999), Equals, 5)
	c.Assert(CountDigits(-45999), Equals, 6)
}

func (s *FmtUtilSuite) TestColorizePassword(c *C) {
	p1 := "acbdabcd"
	p2 := "ABcd12AB"
	p3 := "AB[3a=c_"

	c.Assert(ColorizePassword(p1, "{r}", "{g}", "{y}"), Equals, "{r}acbdabcd{!}")
	c.Assert(ColorizePassword(p2, "{r}", "{g}", "{y}"), Equals, "{r}ABcd{g}12{r}AB{!}")
	c.Assert(ColorizePassword(p3, "{r}", "{g}", "{y}"), Equals, "{r}AB{y}[{g}3{r}a{y}={r}c{y}_{!}")

	c.Assert(ColorizePassword(p3, "{r}", "", ""), Equals, "{r}AB{!}[3{r}a{!}={r}c{!}_{!}")
	c.Assert(ColorizePassword(p3, "", "{g}", ""), Equals, "{!}AB[{g}3{!}a=c_{!}")
	c.Assert(ColorizePassword(p3, "", "", "{y}"), Equals, "{!}AB{y}[{!}3a{y}={!}c{y}_{!}")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FmtUtilSuite) BenchmarkPrettyNum(c *C) {
	for i := 0; i < c.N; i++ {
		PrettyNum(-123456)
	}
}

func (s *FmtUtilSuite) BenchmarkPrettyNumFloat(c *C) {
	for i := 0; i < c.N; i++ {
		PrettyNum(-123456.15)
	}
}

func (s *FmtUtilSuite) BenchmarkWrap(c *C) {
	for i := 0; i < c.N; i++ {
		Wrap("0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0", "  ", 10)
	}
}
