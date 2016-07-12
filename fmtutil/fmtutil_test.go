package fmtutil

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

type FmtUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&FmtUtilSuite{})

func (s *FmtUtilSuite) TestPretyNum(c *C) {
	c.Assert(PrettyNum(999), Equals, "999")
	c.Assert(PrettyNum(1000), Equals, "1,000")
	c.Assert(PrettyNum(1234567890), Equals, "1,234,567,890")
	c.Assert(PrettyNum(100000), Equals, "100,000")
	c.Assert(PrettyNum(0), Equals, "0")
	c.Assert(PrettyNum(2500.50), Equals, "2,500.50")
	c.Assert(PrettyNum(1.23), Equals, "1.23")
	c.Assert(PrettyNum(-1000), Equals, "-1,000")
}

func (s *FmtUtilSuite) TestPretySize(c *C) {
	c.Assert(PrettySize(0), Equals, "0B")
	c.Assert(PrettySize(345), Equals, "345B")
	c.Assert(PrettySize(1025), Equals, "1KB")
	c.Assert(PrettySize(3000125), Equals, "2.86MB")
	c.Assert(PrettySize(1024*1024), Equals, "1MB")
	c.Assert(PrettySize(1024*1024*1024), Equals, "1GB")
	c.Assert(PrettySize(1024*1024*1024*1024), Equals, "1TB")
	c.Assert(PrettySize(1052500000), Equals, "1003.7MB")

	c.Assert(PrettySize(int32(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(int64(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(uint(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(uint32(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(uint64(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(float32(3000125)), Equals, "2.86MB")
	c.Assert(PrettySize(float64(3000125)), Equals, "2.86MB")
}

func (s *FmtUtilSuite) TestParseSize(c *C) {
	c.Assert(ParseSize("1 MB"), Equals, uint64(1024*1024))
	c.Assert(ParseSize("2tb"), Equals, uint64(2*1024*1024*1024*1024))
	c.Assert(ParseSize("5gB"), Equals, uint64(5*1024*1024*1024))
	c.Assert(ParseSize("13kb"), Equals, uint64(13*1024))
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
consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?
`
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
}

func (s *FmtUtilSuite) TestSeparator(c *C) {
	Separator(true)
	Separator(false)
	Separator(true, "test")
	Separator(false, "test")
}
