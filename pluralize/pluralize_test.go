package pluralize

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

type PluralizeSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PluralizeSuite{})

func (s *PluralizeSuite) TestBasic(c *C) {
	data := []string{"A", "B"}

	c.Assert(P("%d %s", 1, data...), Equals, "1 A")
	c.Assert(P("%s %d", 2, data...), Equals, "B 2")
	c.Assert(P("%d %s", -1, data...), Equals, "-1 A")
	c.Assert(P("%s %d", -2, data...), Equals, "B -2")
	c.Assert(PS(En, "%s %d", 2, data...), Equals, "B 2")

	c.Assert(P("%d %s", int8(1), data...), Equals, "1 A")
	c.Assert(P("%d %s", int16(1), data...), Equals, "1 A")
	c.Assert(P("%d %s", int32(1), data...), Equals, "1 A")
	c.Assert(P("%d %s", int64(1), data...), Equals, "1 A")
	c.Assert(P("%d %s", uint(1), data...), Equals, "1 A")
	c.Assert(P("%d %s", uint8(1), data...), Equals, "1 A")
	c.Assert(P("%d %s", uint16(1), data...), Equals, "1 A")
	c.Assert(P("%d %s", uint32(1), data...), Equals, "1 A")
	c.Assert(P("%d %s", uint64(1), data...), Equals, "1 A")
	c.Assert(P("%g %s", 1.1, data...), Equals, "1.1 A")
	c.Assert(P("%g %s", float32(1.1), data...), Equals, "1.1 A")

	c.Assert(convertNumber("abcd"), Equals, uint64(0))
}

func (s *PluralizeSuite) TestAf(c *C) {
	data := []string{"A", "B"}

	c.Assert(Pluralize(1, data...), Equals, "A")
	c.Assert(Pluralize(2), Equals, "")
	c.Assert(PluralizeSpecial(Af, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Af, 2, data...), Equals, "B")
}

func (s *PluralizeSuite) TestAch(c *C) {
	data := []string{"A", "B"}

	c.Assert(PluralizeSpecial(Ach, 0, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Ach, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Ach, 3, data...), Equals, "B")
}

func (s *PluralizeSuite) TestBe(c *C) {
	data := []string{"A", "B", "C"}

	c.Assert(PluralizeSpecial(Be, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Be, 2, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Be, 9, data...), Equals, "C")
}

func (s *PluralizeSuite) TestAr(c *C) {
	data := []string{"A", "B", "C", "D", "E", "F"}

	c.Assert(PluralizeSpecial(Ar, 0, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Ar, 1, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Ar, 2, data...), Equals, "C")
	c.Assert(PluralizeSpecial(Ar, 3, data...), Equals, "D")
	c.Assert(PluralizeSpecial(Ar, 12, data...), Equals, "E")
	c.Assert(PluralizeSpecial(Ar, 100, data...), Equals, "F")
}

func (s *PluralizeSuite) TestAy(c *C) {
	data := []string{"A"}

	c.Assert(PluralizeSpecial(Ay, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Ay, 2, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Ay, 9, data...), Equals, "A")
}

func (s *PluralizeSuite) TestCs(c *C) {
	data := []string{"A", "B", "C"}

	c.Assert(PluralizeSpecial(Cs, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Cs, 3, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Cs, 5, data...), Equals, "C")
}

func (s *PluralizeSuite) TestCsb(c *C) {
	data := []string{"A", "B", "C"}

	c.Assert(PluralizeSpecial(Csb, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Csb, 3, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Csb, 25, data...), Equals, "C")
}

func (s *PluralizeSuite) TestCy(c *C) {
	data := []string{"A", "B", "C", "D"}

	c.Assert(PluralizeSpecial(Cy, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Cy, 2, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Cy, 5, data...), Equals, "C")
	c.Assert(PluralizeSpecial(Cy, 8, data...), Equals, "D")
}

func (s *PluralizeSuite) TestGa(c *C) {
	data := []string{"A", "B", "C", "D", "E"}

	c.Assert(PluralizeSpecial(Ga, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Ga, 2, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Ga, 5, data...), Equals, "C")
	c.Assert(PluralizeSpecial(Ga, 8, data...), Equals, "D")
	c.Assert(PluralizeSpecial(Ga, 14, data...), Equals, "E")
}

func (s *PluralizeSuite) TestGd(c *C) {
	data := []string{"A", "B", "C", "D"}

	c.Assert(PluralizeSpecial(Gd, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Gd, 11, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Gd, 2, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Gd, 12, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Gd, 3, data...), Equals, "C")
	c.Assert(PluralizeSpecial(Gd, 25, data...), Equals, "D")
}

func (s *PluralizeSuite) TestIs(c *C) {
	data := []string{"A", "B"}

	c.Assert(PluralizeSpecial(Is, 11, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Is, 1, data...), Equals, "A")
}

func (s *PluralizeSuite) TestKw(c *C) {
	data := []string{"A", "B", "C", "D"}

	c.Assert(PluralizeSpecial(Kw, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Kw, 2, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Kw, 3, data...), Equals, "C")
	c.Assert(PluralizeSpecial(Kw, 4, data...), Equals, "D")
}

func (s *PluralizeSuite) TestLv(c *C) {
	data := []string{"A", "B", "C"}

	c.Assert(PluralizeSpecial(Lv, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Lv, 3, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Lv, 0, data...), Equals, "C")
}

func (s *PluralizeSuite) TestMk(c *C) {
	data := []string{"A", "B"}

	c.Assert(PluralizeSpecial(Mk, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Mk, 11, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Mk, 2, data...), Equals, "B")
}

func (s *PluralizeSuite) TestMnk(c *C) {
	data := []string{"A", "B", "C"}

	c.Assert(PluralizeSpecial(Mnk, 0, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Mnk, 1, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Mnk, 5, data...), Equals, "C")
}

func (s *PluralizeSuite) TestMt(c *C) {
	data := []string{"A", "B", "C", "D"}

	c.Assert(PluralizeSpecial(Mt, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Mt, 0, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Mt, 5, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Mt, 13, data...), Equals, "C")
	c.Assert(PluralizeSpecial(Mt, 22, data...), Equals, "D")
}

func (s *PluralizeSuite) TestRo(c *C) {
	data := []string{"A", "B", "C"}

	c.Assert(PluralizeSpecial(Ro, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Ro, 0, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Ro, 15, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Ro, 25, data...), Equals, "C")
}

func (s *PluralizeSuite) TestSk(c *C) {
	data := []string{"A", "B", "C"}

	c.Assert(PluralizeSpecial(Sk, 1, data...), Equals, "A")
	c.Assert(PluralizeSpecial(Sk, 3, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Sk, 5, data...), Equals, "C")
}

func (s *PluralizeSuite) TestSl(c *C) {
	data := []string{"A", "B", "C", "D"}

	c.Assert(PluralizeSpecial(Sl, 1, data...), Equals, "B")
	c.Assert(PluralizeSpecial(Sl, 2, data...), Equals, "C")
	c.Assert(PluralizeSpecial(Sl, 3, data...), Equals, "D")
	c.Assert(PluralizeSpecial(Sl, 4, data...), Equals, "A")
}
