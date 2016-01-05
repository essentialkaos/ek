package spellcheck

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

func Test(t *testing.T) { TestingT(t) }

type SpellcheckSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SpellcheckSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SpellcheckSuite) TestSpellcheck(c *C) {
	model := Train([]string{})

	c.Assert(model, NotNil)
	c.Assert(model.Correct("test-1234"), Equals, "test-1234")
	c.Assert(model.Suggest("test-1234", 10), DeepEquals, []string{"test-1234"})

	model = Train([]string{"test", "MAGIC", "TeStInG", "", "Random"})

	c.Assert(model, NotNil)

	c.Assert(model.Correct("test"), Equals, "test")
	c.Assert(model.Correct(""), Equals, "")
	c.Assert(model.Correct("test123test123"), Equals, "test123test123")
	c.Assert(model.Correct("tes"), Equals, "test")
	c.Assert(model.Correct("magic"), Equals, "MAGIC")
	c.Assert(model.Correct("testin"), Equals, "TeStInG")
	c.Assert(model.Correct("rANDOM"), Equals, "Random")

	c.Assert(model.Suggest("tes", 3), DeepEquals, []string{"test", "", "TeStInG"})
	c.Assert(model.Suggest("tes", 1), DeepEquals, []string{"test"})
}
