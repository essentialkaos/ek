package update

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

type UpdateSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&UpdateSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *UpdateSuite) TestGitHubChecker(c *C) {
	newVersion, releaseDate, hasUpdate := GitHubChecker("GitHubChecker", "0.9.9", "")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitHubChecker("GitHubChecker", "0.9.9", "essentialkaos/ftllister")

	c.Assert(newVersion, Equals, "1.0.0")
	c.Assert(releaseDate.Unix(), Equals, int64(1461710348))
	c.Assert(hasUpdate, Equals, true)
}
