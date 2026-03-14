package barcode

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

type BarcodeSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&BarcodeSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *BarcodeSuite) TestDots(c *C) {
	c.Assert(Dots([]byte("")), Equals, "\x1b[38;5;1m⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒\x1b[0m")
	c.Assert(Dots([]byte("Test1234!")), Equals, "\x1b[38;5;1m⠮⠮\x1b[0m\x1b[38;5;3m⠒\x1b[0m\x1b[38;5;4m⠤\x1b[0m\x1b[38;5;3m⠒\x1b[0m\x1b[38;5;1m⠮⠮\x1b[0m\x1b[38;5;3m⠒\x1b[0m\x1b[38;5;1m⠮\x1b[0m\x1b[38;5;4m⠤\x1b[0m\x1b[38;5;2m⠉\x1b[0m\x1b[38;5;3m⠒\x1b[0m\x1b[38;5;1m⠮\x1b[0m\x1b[38;5;2m⠉⠉\x1b[0m\x1b[38;5;3m⠒\x1b[0m")
}

func (s *BarcodeSuite) TestLines(c *C) {
	c.Assert(Lines([]byte("")), Equals, "\x1b[38;5;1m╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬\x1b[0m")
	c.Assert(Lines([]byte("Test1234!")), Equals, "\x1b[38;5;1m╬╬\x1b[0m\x1b[38;5;3m╩\x1b[0m\x1b[38;5;4m╧\x1b[0m\x1b[38;5;3m╩\x1b[0m\x1b[38;5;1m╬╬\x1b[0m\x1b[38;5;3m╩\x1b[0m\x1b[38;5;1m╬\x1b[0m\x1b[38;5;4m╧\x1b[0m\x1b[38;5;2m╪\x1b[0m\x1b[38;5;3m╩\x1b[0m\x1b[38;5;1m╬\x1b[0m\x1b[38;5;2m╪╪\x1b[0m\x1b[38;5;3m╩\x1b[0m")
}

func (s *BarcodeSuite) TestBoxes(c *C) {
	c.Assert(Boxes([]byte("")), Equals, "\x1b[38;5;1m████████████████\x1b[0m")
	c.Assert(Boxes([]byte("Test1234!")), Equals, "\x1b[38;5;1m▁▁\x1b[0m\x1b[38;5;3m▃\x1b[0m\x1b[38;5;4m▄\x1b[0m\x1b[38;5;3m▃\x1b[0m\x1b[38;5;1m▁▁\x1b[0m\x1b[38;5;3m▃\x1b[0m\x1b[38;5;1m▁\x1b[0m\x1b[38;5;4m▄\x1b[0m\x1b[38;5;2m▂\x1b[0m\x1b[38;5;3m▃\x1b[0m\x1b[38;5;1m▁\x1b[0m\x1b[38;5;2m▂▂\x1b[0m\x1b[38;5;3m▃\x1b[0m")
}
