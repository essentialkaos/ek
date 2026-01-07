package emoji

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
