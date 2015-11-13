package spellcheck

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type SpellcheckSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SpellcheckSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SpellcheckSuite) TestSpellcheck(c *C) {
	m := Train([]string{"test", "MAGIC", "TeStInG", "", "Random"})

	c.Assert(m.Correct("test"), Equals, "test")
	c.Assert(m.Correct("tes"), Equals, "test")
	c.Assert(m.Correct("magic"), Equals, "MAGIC")
	c.Assert(m.Correct("testin"), Equals, "TeStInG")
	c.Assert(m.Correct("rANDOM"), Equals, "Random")
}
