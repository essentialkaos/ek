package emoji

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type EmojiSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&EmojiSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *EmojiSuite) TestGet(c *C) {
	c.Assert(Get("100"), Equals, "💯")
	c.Assert(Get("_unknown_"), Equals, "")
}

func (s *EmojiSuite) TestGetName(c *C) {
	c.Assert(GetName("⚡️"), Equals, "zap")
	c.Assert(GetName("_unknown_"), Equals, "")
}

func (s *EmojiSuite) TestFind(c *C) {
	c.Assert(Find("bikin"), HasLen, 3)
}

func (s *EmojiSuite) TestEmojize(c *C) {
	c.Assert(Emojize("Hi :smile: emoji: :zap:!"), Equals, "Hi 😄 emoji: ⚡️!")
	c.Assert(Emojize("Hi :smile__1: emoji: :zap:!"), Equals, "Hi :smile__1: emoji: ⚡️!")
	c.Assert(Emojize(""), Equals, "")
}
