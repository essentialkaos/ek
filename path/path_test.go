package path

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type PathUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PathUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PathUtilSuite) TestBase(c *C) {
	c.Assert(Base("/some/test/path"), Equals, "path")
	c.Assert(Clean("//some/test//path"), Equals, "/some/test/path")
	c.Assert(Dir("/some/test/path"), Equals, "/some/test")
	c.Assert(Ext("/some/test/path/file.jpg"), Equals, ".jpg")
	c.Assert(IsAbs("/some/test/path"), Equals, true)
	c.Assert(Join("/", "some", "test", "path"), Equals, "/some/test/path")

	match, err := Match("/some/test/*", "/some/test/path")

	c.Assert(err, IsNil)
	c.Assert(match, Equals, true)

	d, f := Split("/some/test/path/file.jpg")

	c.Assert(d, Equals, "/some/test/path/")
	c.Assert(f, Equals, "file.jpg")
}

func (s *PathUtilSuite) TestDirN(c *C) {
	c.Assert(DirN("", 99), Equals, "")
	c.Assert(DirN("1", 99), Equals, "1")
	c.Assert(DirN("abcde", 1), Equals, "abcde")
	c.Assert(DirN("/a/b/c/d", -1), Equals, "/a/b/c/d")
	c.Assert(DirN("/a/b/c/d", 1), Equals, "/a")
	c.Assert(DirN("a/b/c/d", 2), Equals, "a/b")
	c.Assert(DirN("/a/b/c/d", 99), Equals, "/a/b/c/d")
}

func (s *PathUtilSuite) TestEvalHome(c *C) {
	homeDir := os.Getenv("HOME")

	c.Assert(Clean("~/path"), Equals, homeDir+"/path")
	c.Assert(Clean("/path"), Equals, "/path")
}

func (s *PathUtilSuite) TestSafe(c *C) {
	c.Assert(IsSafe("/home/user/test.jpg"), Equals, true)
	c.Assert(IsSafe("/home/user"), Equals, true)
	c.Assert(IsSafe("/opt/software-1234"), Equals, true)
	c.Assert(IsSafe("/srv/my-supper-service"), Equals, true)

	c.Assert(IsSafe(""), Equals, false)
	c.Assert(IsSafe("/"), Equals, false)
	c.Assert(IsSafe("/dev/tty3"), Equals, false)
	c.Assert(IsSafe("/etc/file.conf"), Equals, false)
	c.Assert(IsSafe("/lib/some-lib"), Equals, false)
	c.Assert(IsSafe("/lib64/some-lib"), Equals, false)
	c.Assert(IsSafe("/lost+found"), Equals, false)
	c.Assert(IsSafe("/proc/19313"), Equals, false)
	c.Assert(IsSafe("/root"), Equals, false)
	c.Assert(IsSafe("/sbin/useradd"), Equals, false)
	c.Assert(IsSafe("/bin/useradd"), Equals, false)
	c.Assert(IsSafe("/selinux"), Equals, false)
	c.Assert(IsSafe("/sys/kernel"), Equals, false)
	c.Assert(IsSafe("/usr/bin/du"), Equals, false)
	c.Assert(IsSafe("/usr/sbin/chroot"), Equals, false)
	c.Assert(IsSafe("/usr/lib/some-lib"), Equals, false)
	c.Assert(IsSafe("/usr/lib64/some-lib"), Equals, false)
	c.Assert(IsSafe("/usr/libexec/gcc"), Equals, false)
	c.Assert(IsSafe("/usr/include/xlocale.h"), Equals, false)
	c.Assert(IsSafe("/var/cache/yum"), Equals, false)
	c.Assert(IsSafe("/var/db/yum"), Equals, false)
	c.Assert(IsSafe("/var/lib/pgsql"), Equals, false)
}

func (s *PathUtilSuite) TestDotfile(c *C) {
	c.Assert(IsDotfile(""), Equals, false)
	c.Assert(IsDotfile("/some/dir/abcd"), Equals, false)
	c.Assert(IsDotfile("/some/dir/"), Equals, false)
	c.Assert(IsDotfile("/"), Equals, false)
	c.Assert(IsDotfile("/////"), Equals, false)
	c.Assert(IsDotfile("   /    "), Equals, false)

	c.Assert(IsDotfile(".dotfile"), Equals, true)
	c.Assert(IsDotfile("/.dotfile"), Equals, true)
	c.Assert(IsDotfile("/some/dir/.abcd"), Equals, true)
}

func (s *PathUtilSuite) TestGlob(c *C) {
	c.Assert(IsGlob(""), Equals, false)
	c.Assert(IsGlob("ancd-1234"), Equals, false)
	c.Assert(IsGlob("[1234"), Equals, false)
	c.Assert(IsGlob("test*"), Equals, true)
	c.Assert(IsGlob("t?st"), Equals, true)
	c.Assert(IsGlob("t[a-z]st"), Equals, true)
}
