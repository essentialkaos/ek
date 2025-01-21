package path

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PathUtilSuite) TestBase(c *C) {
	c.Assert(Base(`C:\some\test\path`), Equals, "path")
	c.Assert(Clean(`C:\\some\test\\\path`), Equals, `C:\some\test\path`)
	c.Assert(Dir(`C:\some\test\path`), Equals, `C:\some\test`)
	c.Assert(Ext(`C:\some\test\path\file.jpg`), Equals, ".jpg")
	c.Assert(IsAbs(`C:\some\test\path`), Equals, true)
	c.Assert(Join("C:\\", "some", "test", "path"), Equals, `C:\some\test\path`)

	match, err := Match(`C:\some\test\*`, `C:\some\test\path`)

	c.Assert(err, IsNil)
	c.Assert(match, Equals, true)

	d, f := Split(`C:\some\test\path\file.jpg`)

	c.Assert(d, Equals, `C:\some\test\path\`)
	c.Assert(f, Equals, "file.jpg")
}

func (s *PathUtilSuite) TestJoinSecure(c *C) {
	p, err := JoinSecure(`C:\test`, "myapp")
	c.Assert(err, IsNil)
	c.Assert(p, Equals, `C:\test\myapp`)

	p, err = JoinSecure(`C:\test`, `myapp\config\..\global.cfg`)
	c.Assert(err, IsNil)
	c.Assert(p, Equals, `C:\test\myapp\global.cfg`)
}

func (s *PathUtilSuite) TestDirN(c *C) {
	c.Assert(DirN("", 99), Equals, "")
	c.Assert(DirN("1", 99), Equals, "1")
	c.Assert(DirN("abcde", 1), Equals, "abcde")
	c.Assert(DirN("abcde", -1), Equals, "abcde")
	c.Assert(DirN(`C:\a\b\c\d`, 0), Equals, `C:\a\b\c\d`)
	c.Assert(DirN(`C:\a\b\c\d`, -1), Equals, `C:\a\b\c`)
	c.Assert(DirN(`C:\a\b\c\d\`, -1), Equals, `C:\a\b\c`)
	c.Assert(DirN(`C:\a\b\c\d`, 1), Equals, `C:\a`)
	c.Assert(DirN(`a\b\c\d`, 2), Equals, `a\b`)
	c.Assert(DirN("/a/b/c/d", 99), Equals, "/a/b/c/d")

	c.Assert(DirN(`\\\\\\\\\`, 2), Equals, `\\`)
	c.Assert(DirN(`\\\\\\\\\`, -2), Equals, `\\\\\\`)
}

func (s *PathUtilSuite) TestSafe(c *C) {
	c.Assert(IsSafe(`C:\Users\john\Pictures\test.jpg`), Equals, true)
	c.Assert(IsSafe(`C:\Users\john`), Equals, true)
	c.Assert(IsSafe(`D:\`), Equals, true)

	c.Assert(IsSafe(`C:\Windows\System32`), Equals, false)
	c.Assert(IsSafe(`C:\Windows\System32`), Equals, false)
}

func (s *PathUtilSuite) TestDotfile(c *C) {
	c.Assert(IsDotfile(""), Equals, false)
	c.Assert(IsDotfile(`C:\some\dir\abcd`), Equals, false)
	c.Assert(IsDotfile(`.\some\dir\`), Equals, false)
	c.Assert(IsDotfile(`\\\\\\`), Equals, false)

	c.Assert(IsDotfile(`.dotfile`), Equals, true)
	c.Assert(IsDotfile(`C:\.dotfile`), Equals, true)
	c.Assert(IsDotfile(`.\some\dir\.abcd`), Equals, true)
}

func (s *PathUtilSuite) TestGlob(c *C) {
	c.Assert(IsGlob(""), Equals, false)
	c.Assert(IsGlob("ancd-1234"), Equals, false)
	c.Assert(IsGlob("[1234"), Equals, false)
	c.Assert(IsGlob("test*"), Equals, true)
	c.Assert(IsGlob("t?st"), Equals, true)
	c.Assert(IsGlob("t[a-z]st"), Equals, true)
}

func (s *PathUtilSuite) TestCompact(c *C) {
	c.Assert(Compact(""), Equals, "")
	c.Assert(Compact("test"), Equals, "test")
	c.Assert(Compact(`C:\a\b\c\d`), Equals, `C:\a\b\c\d`)
	c.Assert(Compact(`C:\my\random\directory\test.txt`), Equals, `C:\m\r\d\test.txt`)
	c.Assert(Compact(`.\my\random\directory\test.txt`), Equals, `.\m\r\d\test.txt`)
}
