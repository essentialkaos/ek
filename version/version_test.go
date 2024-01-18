package version

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type VersionSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&VersionSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *VersionSuite) TestParsing(c *C) {
	_, err1 := Parse("1")
	_, err2 := Parse("2.1")
	_, err3 := Parse("3.4.5")
	_, err4 := Parse("4-alpha1")
	_, err5 := Parse("5-beta1+sha:5114f85")
	_, err6 := Parse("6.12.1-beta2+exp.sha.5114f85")

	c.Assert(err1, IsNil)
	c.Assert(err2, IsNil)
	c.Assert(err3, IsNil)
	c.Assert(err4, IsNil)
	c.Assert(err5, IsNil)
	c.Assert(err6, IsNil)
}

func (s *VersionSuite) TestMajor(c *C) {
	v1, _ := Parse("1")
	v2, _ := Parse("2.0.0")
	v3, _ := Parse("3.4.5")
	v4, _ := Parse("4-beta1")
	v5, _ := Parse("5-beta2+exp.sha.5114f85")

	c.Assert(v1.Major(), Equals, 1)
	c.Assert(v2.Major(), Equals, 2)
	c.Assert(v3.Major(), Equals, 3)
	c.Assert(v4.Major(), Equals, 4)
	c.Assert(v5.Major(), Equals, 5)
}

func (s *VersionSuite) TestMinor(c *C) {
	v1, _ := Parse("1")
	v2, _ := Parse("2.1")
	v3, _ := Parse("3.4.5")
	v4, _ := Parse("4-alpha1")
	v5, _ := Parse("5-beta1+sha:5114f85")
	v6, _ := Parse("6.12.1-beta2+exp.sha.5114f85")

	c.Assert(v1.Minor(), Equals, 0)
	c.Assert(v2.Minor(), Equals, 1)
	c.Assert(v3.Minor(), Equals, 4)
	c.Assert(v4.Minor(), Equals, 0)
	c.Assert(v5.Minor(), Equals, 0)
	c.Assert(v6.Minor(), Equals, 12)
}

func (s *VersionSuite) TestPatch(c *C) {
	v1, _ := Parse("1")
	v2, _ := Parse("2.1")
	v3, _ := Parse("3.4.5")
	v4, _ := Parse("4-alpha1")
	v5, _ := Parse("5-beta1+sha:5114f85")
	v6, _ := Parse("6.12.1-beta2+exp.sha.5114f85")

	c.Assert(v1.Patch(), Equals, 0)
	c.Assert(v2.Patch(), Equals, 0)
	c.Assert(v3.Patch(), Equals, 5)
	c.Assert(v4.Patch(), Equals, 0)
	c.Assert(v5.Patch(), Equals, 0)
	c.Assert(v6.Patch(), Equals, 1)
}

func (s *VersionSuite) TestPre(c *C) {
	v1, _ := Parse("1")
	v2, _ := Parse("2.1")
	v3, _ := Parse("3.4.5")
	v4, _ := Parse("4-alpha1")
	v5, _ := Parse("5-beta1+sha:5114f85")
	v6, _ := Parse("6.12.1-beta2+exp.sha.5114f85")

	c.Assert(v1.PreRelease(), Equals, "")
	c.Assert(v2.PreRelease(), Equals, "")
	c.Assert(v3.PreRelease(), Equals, "")
	c.Assert(v4.PreRelease(), Equals, "alpha1")
	c.Assert(v5.PreRelease(), Equals, "beta1")
	c.Assert(v6.PreRelease(), Equals, "beta2")
}

func (s *VersionSuite) TestBuild(c *C) {
	v1, _ := Parse("1")
	v2, _ := Parse("2.1")
	v3, _ := Parse("3.4.5")
	v4, _ := Parse("4-alpha1")
	v5, _ := Parse("5-beta1+sha:5114f85")
	v6, _ := Parse("6.12.1-beta2+exp.sha.5114f85")

	c.Assert(v1.Build(), Equals, "")
	c.Assert(v2.Build(), Equals, "")
	c.Assert(v3.Build(), Equals, "")
	c.Assert(v4.Build(), Equals, "")
	c.Assert(v5.Build(), Equals, "sha:5114f85")
	c.Assert(v6.Build(), Equals, "exp.sha.5114f85")
}

func (s *VersionSuite) TestSimple(c *C) {
	v1, _ := Parse("1")
	v2, _ := Parse("2.1")
	v3, _ := Parse("3.4.5")
	v4, _ := Parse("4-alpha1")
	v5, _ := Parse("5-beta1+sha:5114f85")
	v6, _ := Parse("6.12.1-beta2+exp.sha.5114f85")

	c.Assert(v1.Simple(), Equals, "1.0.0")
	c.Assert(v2.Simple(), Equals, "2.1.0")
	c.Assert(v3.Simple(), Equals, "3.4.5")
	c.Assert(v4.Simple(), Equals, "4.0.0")
	c.Assert(v5.Simple(), Equals, "5.0.0")
	c.Assert(v6.Simple(), Equals, "6.12.1")
}

func (s *VersionSuite) TestString(c *C) {
	v1, _ := Parse("1")
	v2, _ := Parse("2.1")
	v3, _ := Parse("3.4.5")
	v4, _ := Parse("5-beta1+sha:5114f85")
	v5, _ := Parse("6.12.1-beta2+exp.sha.5114f85")

	c.Assert(v1.String(), Equals, "1")
	c.Assert(v2.String(), Equals, "2.1")
	c.Assert(v3.String(), Equals, "3.4.5")
	c.Assert(v4.String(), Equals, "5-beta1+sha:5114f85")
	c.Assert(v5.String(), Equals, "6.12.1-beta2+exp.sha.5114f85")
}

func (s *VersionSuite) TestIsZero(c *C) {
	var v1 Version

	v2, _ := Parse("1")

	c.Assert(v1.IsZero(), Equals, true)
	c.Assert(v2.IsZero(), Equals, false)
}

func (s *VersionSuite) TestErrors(c *C) {
	var err error

	var v1 Version
	var v2 = Version{}

	_, err = Parse("A")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `strconv.Atoi: parsing "A": invalid syntax`)

	_, err = Parse(" ")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `strconv.Atoi: parsing " ": invalid syntax`)

	_, err = Parse("")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Version can't be empty`)

	_, err = Parse("1.2.B")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `strconv.Atoi: parsing "B": invalid syntax`)

	_, err = Parse("1.2.8-")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Prerelease number is empty`)

	_, err = Parse("1.2.8-1+")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Build number is empty`)

	c.Assert(v1.Major(), Equals, -1)
	c.Assert(v1.Minor(), Equals, -1)
	c.Assert(v1.Patch(), Equals, -1)
	c.Assert(v1.PreRelease(), Equals, "")
	c.Assert(v1.Build(), Equals, "")
	c.Assert(v1.String(), Equals, "")
	c.Assert(v1.Simple(), Equals, "0.0.0")

	c.Assert(v2.Major(), Equals, -1)
	c.Assert(v2.Minor(), Equals, -1)
	c.Assert(v2.Patch(), Equals, -1)
	c.Assert(v2.PreRelease(), Equals, "")
	c.Assert(v2.Build(), Equals, "")
	c.Assert(v2.String(), Equals, "")
	c.Assert(v2.Simple(), Equals, "0.0.0")
}

func (s *VersionSuite) TestComparison(c *C) {
	var P = func(version string) Version {
		v, _ := Parse(version)
		return v
	}

	c.Assert(P("1").Equal(P("1")), Equals, true)
	c.Assert(P("1").Equal(P("2")), Equals, false)
	c.Assert(P("1").Equal(P("1.0")), Equals, true)
	c.Assert(P("1").Equal(P("1.1")), Equals, false)
	c.Assert(P("1").Equal(P("1.0.0")), Equals, true)
	c.Assert(P("1").Equal(P("1.0.1")), Equals, false)
	c.Assert(P("1").Equal(P("1.0.0-alpha1")), Equals, false)
	c.Assert(P("1").Equal(P("1.0.0+sha:5114f85")), Equals, false)
	c.Assert(P("1.0.0+sha:5114f85").Equal(P("1.0.0+sha:5114f85")), Equals, true)

	c.Assert(P("1").Less(P("1")), Equals, false)
	c.Assert(P("1").Less(P("1.0")), Equals, false)
	c.Assert(P("1").Less(P("1.0.0")), Equals, false)
	c.Assert(P("1").Less(P("2")), Equals, true)
	c.Assert(P("1").Less(P("1.1")), Equals, true)
	c.Assert(P("1").Less(P("1.0.1")), Equals, true)
	c.Assert(P("1.0.1-alpha").Less(P("1.0.1")), Equals, true)
	c.Assert(P("1.0.1").Less(P("1.0.1-alpha")), Equals, false)
	c.Assert(P("1.0.1-alpha").Less(P("1.0.1-beta")), Equals, true)
	c.Assert(P("1.0.1-gamma").Less(P("1.0.1-beta")), Equals, false)
	c.Assert(P("1.0.1-alpha").Less(P("1.0.1-alpha1")), Equals, true)
	c.Assert(P("1.0.1-a4").Less(P("1.0.1-a5")), Equals, true)
	c.Assert(P("1.0.1-a5").Less(P("1.0.1-a5")), Equals, false)
	c.Assert(P("1.11.0").Less(P("1.10.0")), Equals, false)
	c.Assert(P("1.0.11").Less(P("1.0.10")), Equals, false)
	c.Assert(P("2.0.0").Less(P("1.1.0")), Equals, false)

	c.Assert(P("1").Greater(P("1")), Equals, false)
	c.Assert(P("1").Greater(P("1.0")), Equals, false)
	c.Assert(P("1").Greater(P("1.0.0")), Equals, false)
	c.Assert(P("2").Greater(P("1")), Equals, true)
	c.Assert(P("1.1").Greater(P("1")), Equals, true)
	c.Assert(P("1.0.1").Greater(P("1")), Equals, true)
	c.Assert(P("1.0.1-alpha").Greater(P("1.0.1")), Equals, false)
	c.Assert(P("1.0.1").Greater(P("1.0.1-alpha")), Equals, true)
	c.Assert(P("1.0.1-alpha").Greater(P("1.0.1-beta")), Equals, false)
	c.Assert(P("1.0.1-gamma").Greater(P("1.0.1-beta")), Equals, true)
	c.Assert(P("1.0.1-alpha").Greater(P("1.0.1-alpha1")), Equals, false)
	c.Assert(P("1.0.1-a4").Greater(P("1.0.1-a5")), Equals, false)
	c.Assert(P("1.0.1-a5").Greater(P("1.0.1-a5")), Equals, false)
	c.Assert(P("1.10.0").Greater(P("1.11.0")), Equals, false)
	c.Assert(P("1.0.10").Greater(P("1.0.11")), Equals, false)
	c.Assert(P("2.0.0").Greater(P("1.1.0")), Equals, true)

	c.Assert(P("1").Contains(P("1")), Equals, true)
	c.Assert(P("1").Contains(P("1.1")), Equals, true)
	c.Assert(P("1").Contains(P("1.0.1")), Equals, true)
	c.Assert(P("2").Contains(P("1")), Equals, false)
	c.Assert(P("1.1").Contains(P("1.2")), Equals, false)
	c.Assert(P("1.0").Contains(P("1.0.2")), Equals, true)
	c.Assert(P("1.0.1").Contains(P("1.0.2")), Equals, false)
	c.Assert(P("1.0.1").Contains(P("1.0.1-alpha")), Equals, false)

	c.Assert(P("0.10.8").Greater(P("1.0.0")), Equals, false)
	c.Assert(P("1.0.0").Less(P("0.10.8")), Equals, false)
}
