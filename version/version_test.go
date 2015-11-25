package version

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "gopkg.in/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type VersionSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&VersionSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *VersionSuite) TestMajor(c *C) {
	c.Assert(Parse("1").Major(), Equals, 1)
	c.Assert(Parse("2.0.0").Major(), Equals, 2)
	c.Assert(Parse("3.4.5").Major(), Equals, 3)
	c.Assert(Parse("4-beta1").Major(), Equals, 4)
	c.Assert(Parse("5-beta2+exp.sha.5114f85").Major(), Equals, 5)
}

func (s *VersionSuite) TestMinor(c *C) {
	c.Assert(Parse("1").Minor(), Equals, 0)
	c.Assert(Parse("2.1").Minor(), Equals, 1)
	c.Assert(Parse("3.4.5").Minor(), Equals, 4)
	c.Assert(Parse("4-alpha1").Minor(), Equals, 0)
	c.Assert(Parse("5-beta1+sha:5114f85").Minor(), Equals, 0)
	c.Assert(Parse("6.12.1-beta2+exp.sha.5114f85").Minor(), Equals, 12)
}

func (s *VersionSuite) TestPatch(c *C) {
	c.Assert(Parse("1").Patch(), Equals, 0)
	c.Assert(Parse("2.1").Patch(), Equals, 0)
	c.Assert(Parse("3.4.5").Patch(), Equals, 5)
	c.Assert(Parse("4-alpha1").Patch(), Equals, 0)
	c.Assert(Parse("5-beta1+sha:5114f85").Patch(), Equals, 0)
	c.Assert(Parse("6.12.1-beta2+exp.sha.5114f85").Patch(), Equals, 1)
}

func (s *VersionSuite) TestPre(c *C) {
	c.Assert(Parse("1").PreRelease(), Equals, "")
	c.Assert(Parse("2.1").PreRelease(), Equals, "")
	c.Assert(Parse("3.4.5").PreRelease(), Equals, "")
	c.Assert(Parse("3.4.5-").PreRelease(), Equals, "")
	c.Assert(Parse("4-alpha1").PreRelease(), Equals, "alpha1")
	c.Assert(Parse("5-beta1+sha:5114f85").PreRelease(), Equals, "beta1")
	c.Assert(Parse("6.12.1-beta2+exp.sha.5114f85").PreRelease(), Equals, "beta2")
}

func (s *VersionSuite) TestBuild(c *C) {
	c.Assert(Parse("1").Build(), Equals, "")
	c.Assert(Parse("2.1").Build(), Equals, "")
	c.Assert(Parse("3.4.5").Build(), Equals, "")
	c.Assert(Parse("4-alpha1+").Build(), Equals, "")
	c.Assert(Parse("5-beta1+sha:5114f85").Build(), Equals, "sha:5114f85")
	c.Assert(Parse("6.12.1-beta2+exp.sha.5114f85").Build(), Equals, "exp.sha.5114f85")
}

func (s *VersionSuite) TestString(c *C) {
	c.Assert(Parse("1").String(), Equals, "1")
	c.Assert(Parse("2.1").String(), Equals, "2.1")
	c.Assert(Parse("3.4.5").String(), Equals, "3.4.5")
	c.Assert(Parse("5-beta1+sha:5114f85").String(), Equals, "5-beta1+sha:5114f85")
	c.Assert(Parse("6.12.1-beta2+exp.sha.5114f85").String(), Equals, "6.12.1-beta2+exp.sha.5114f85")
}

func (s *VersionSuite) TestErrors(c *C) {
	var v1 *Version
	var v2 = &Version{}

	c.Assert(Parse("A").Major(), Equals, -1)
	c.Assert(Parse("A").Minor(), Equals, -1)
	c.Assert(Parse("A").Patch(), Equals, -1)
	c.Assert(Parse("A").PreRelease(), Equals, "")
	c.Assert(Parse("A").Build(), Equals, "")

	c.Assert(Parse("").Major(), Equals, -1)
	c.Assert(Parse("").Minor(), Equals, -1)
	c.Assert(Parse("").Patch(), Equals, -1)
	c.Assert(Parse("").PreRelease(), Equals, "")
	c.Assert(Parse("").Build(), Equals, "")

	c.Assert(v1.Major(), Equals, -1)
	c.Assert(v1.Minor(), Equals, -1)
	c.Assert(v1.Patch(), Equals, -1)
	c.Assert(v1.PreRelease(), Equals, "")
	c.Assert(v1.Build(), Equals, "")
	c.Assert(v1.String(), Equals, "")

	c.Assert(v2.Major(), Equals, -1)
	c.Assert(v2.Minor(), Equals, -1)
	c.Assert(v2.Patch(), Equals, -1)
	c.Assert(v2.PreRelease(), Equals, "")
	c.Assert(v2.Build(), Equals, "")
	c.Assert(v2.String(), Equals, "")
}
