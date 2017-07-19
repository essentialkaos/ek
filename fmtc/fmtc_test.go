package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type FormatSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&FormatSuite{})

func (s *FormatSuite) TestColors(c *C) {
	c.Assert(Sprint("{r}W{!}"), Equals, "\x1b[0;31;49mW\x1b[0m")
	c.Assert(Sprint("{g}W{!}"), Equals, "\x1b[0;32;49mW\x1b[0m")
	c.Assert(Sprint("{y}W{!}"), Equals, "\x1b[0;33;49mW\x1b[0m")
	c.Assert(Sprint("{b}W{!}"), Equals, "\x1b[0;34;49mW\x1b[0m")
	c.Assert(Sprint("{m}W{!}"), Equals, "\x1b[0;35;49mW\x1b[0m")
	c.Assert(Sprint("{c}W{!}"), Equals, "\x1b[0;36;49mW\x1b[0m")
	c.Assert(Sprint("{s}W{!}"), Equals, "\x1b[0;37;49mW\x1b[0m")
	c.Assert(Sprint("{w}W{!}"), Equals, "\x1b[0;97;49mW\x1b[0m")
	c.Assert(Sprint("{r-}W{!}"), Equals, "\x1b[0;91;49mW\x1b[0m")
	c.Assert(Sprint("{g-}W{!}"), Equals, "\x1b[0;92;49mW\x1b[0m")
	c.Assert(Sprint("{y-}W{!}"), Equals, "\x1b[0;93;49mW\x1b[0m")
	c.Assert(Sprint("{b-}W{!}"), Equals, "\x1b[0;94;49mW\x1b[0m")
	c.Assert(Sprint("{m-}W{!}"), Equals, "\x1b[0;95;49mW\x1b[0m")
	c.Assert(Sprint("{c-}W{!}"), Equals, "\x1b[0;96;49mW\x1b[0m")
	c.Assert(Sprint("{s-}W{!}"), Equals, "\x1b[0;90;49mW\x1b[0m")
	c.Assert(Sprint("{w-}W{!}"), Equals, "\x1b[0;97;49mW\x1b[0m")
}

func (s *FormatSuite) TestBackgrounds(c *C) {
	c.Assert(Sprint("{R}W{!}"), Equals, "\x1b[0;39;41mW\x1b[0m")
	c.Assert(Sprint("{G}W{!}"), Equals, "\x1b[0;39;42mW\x1b[0m")
	c.Assert(Sprint("{Y}W{!}"), Equals, "\x1b[0;39;43mW\x1b[0m")
	c.Assert(Sprint("{B}W{!}"), Equals, "\x1b[0;39;44mW\x1b[0m")
	c.Assert(Sprint("{M}W{!}"), Equals, "\x1b[0;39;45mW\x1b[0m")
	c.Assert(Sprint("{C}W{!}"), Equals, "\x1b[0;39;46mW\x1b[0m")
	c.Assert(Sprint("{S}W{!}"), Equals, "\x1b[0;39;47mW\x1b[0m")
	c.Assert(Sprint("{W}W{!}"), Equals, "\x1b[0;39;107mW\x1b[0m")
}

func (s *FormatSuite) TestSpecial(c *C) {
	c.Assert(Sprint("{_}W{!}"), Equals, "\x1b[4;39;49mW\x1b[0m")
	c.Assert(Sprint("{*}W{!}"), Equals, "\x1b[1;39;49mW\x1b[0m")
}

func (s *FormatSuite) TestParsing(c *C) {
	c.Assert(Sprint(""), Equals, "")
	c.Assert(Sprint("W"), Equals, "W")
	c.Assert(Sprint("{"), Equals, "{")
	c.Assert(Sprint("{r"), Equals, "{r")
	c.Assert(Sprint("{J}W"), Equals, "{J}W")
	c.Assert(Sprint("{r}W"), Equals, "\x1b[0;31;49mW\x1b[0m")
	c.Assert(Sprint("{{r}W{!}}"), Equals, "{\x1b[0;31;49mW\x1b[0m}")
	c.Assert(Sprint("Test"+string(rune(65533))), Equals, "Test")
}

func (s *FormatSuite) TestZDisable(c *C) {
	DisableColors = true

	c.Assert(Sprint("{r}W{!}"), Equals, "W")
	c.Assert(Sprint("{g}W{!}"), Equals, "W")
	c.Assert(Sprint("{y}W{!}"), Equals, "W")
	c.Assert(Sprint("{b}W{!}"), Equals, "W")
	c.Assert(Sprint("{m}W{!}"), Equals, "W")
	c.Assert(Sprint("{c}W{!}"), Equals, "W")
	c.Assert(Sprint("{s}W{!}"), Equals, "W")
	c.Assert(Sprint("{R}W{!}"), Equals, "W")
	c.Assert(Sprint("{G}W{!}"), Equals, "W")
	c.Assert(Sprint("{Y}W{!}"), Equals, "W")
	c.Assert(Sprint("{B}W{!}"), Equals, "W")
	c.Assert(Sprint("{M}W{!}"), Equals, "W")
	c.Assert(Sprint("{C}W{!}"), Equals, "W")
	c.Assert(Sprint("{S}W{!}"), Equals, "W")
	c.Assert(Sprint("{S*_}W{!}"), Equals, "W")

	c.Assert(Sprint("Test {config} value"), Equals, "Test {config} value")

	DisableColors = false
}

func (s *FormatSuite) TestClean(c *C) {
	c.Assert(Clean("{r}W{!}"), Equals, "W")
	c.Assert(Clean("{g}W{!}"), Equals, "W")
	c.Assert(Clean("{y}W{!}"), Equals, "W")
	c.Assert(Clean("{b}W{!}"), Equals, "W")
	c.Assert(Clean("{m}W{!}"), Equals, "W")
	c.Assert(Clean("{c}W{!}"), Equals, "W")
	c.Assert(Clean("{s}W{!}"), Equals, "W")
	c.Assert(Clean("{R}W{!}"), Equals, "W")
	c.Assert(Clean("{G}W{!}"), Equals, "W")
	c.Assert(Clean("{Y}W{!}"), Equals, "W")
	c.Assert(Clean("{B}W{!}"), Equals, "W")
	c.Assert(Clean("{M}W{!}"), Equals, "W")
	c.Assert(Clean("{C}W{!}"), Equals, "W")
	c.Assert(Clean("{S}W{!}"), Equals, "W")
	c.Assert(Clean("{S*_}W{!}"), Equals, "W")
}

func (s *FormatSuite) TestMethods(c *C) {
	c.Assert(Errorf("Test %s", "OK"), DeepEquals, errors.New("Test OK"))
	c.Assert(Sprintf("Test %s", "OK"), Equals, "Test OK")

	w := bytes.NewBufferString("")

	Fprint(w, "TEST")

	c.Assert(w.String(), Equals, "TEST")

	w = bytes.NewBufferString("")

	Fprintln(w, "TEST")

	c.Assert(w.String(), Equals, "TEST\n")

	w = bytes.NewBufferString("")

	Fprintf(w, "TEST %s", "OK")

	c.Assert(w.String(), Equals, "TEST OK")

	Printf("TEST %s\n", "OK")
}

func (s *FormatSuite) TestAux(c *C) {
	t := NewT()

	t.Printf("TEST %s", "OK")
	t.Printf("TEST %s", "OK")

	t.Println("TEST OK")

	Bell()
	NewLine()
}
