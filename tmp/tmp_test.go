package tmp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	. "pkg.re/check.v1"

	"pkg.re/essentialkaos/ek.v12/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type TmpSuite struct {
	TempDir string
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TmpSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (ts *TmpSuite) SetUpSuite(c *C) {
	ts.TempDir = c.MkDir()
}

func (ts *TmpSuite) TestMk(c *C) {
	t, err := NewTemp()

	c.Assert(err, IsNil)
	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, "/tmp")

	t.Clean()
}

func (ts *TmpSuite) TestErrors(c *C) {
	t, err := NewTemp("/")

	c.Assert(err, NotNil)
	c.Assert(t, IsNil)

	t, err = NewTemp("/tmpz")

	c.Assert(err, NotNil)
	c.Assert(t, IsNil)

	os.Create(ts.TempDir + "/test_")

	t, err = NewTemp(ts.TempDir + "/test_")

	c.Assert(err, NotNil)
	c.Assert(t, IsNil)

	t = &Temp{Dir: "/"}

	tmpDir, err := t.MkDir("test")

	c.Assert(tmpDir, Equals, "")
	c.Assert(err, NotNil)

	tmpFd, tmpFile, err := t.MkFile("test")

	c.Assert(tmpFd, IsNil)
	c.Assert(tmpFile, Equals, "")
	c.Assert(err, NotNil)

	var nilTemp *Temp

	_, err = nilTemp.MkDir()

	c.Assert(err, NotNil)

	_, _, err = nilTemp.MkFile()

	c.Assert(err, NotNil)
}

func (ts *TmpSuite) TestMkDir(c *C) {
	t, err := NewTemp(ts.TempDir)

	c.Assert(err, IsNil)
	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, ts.TempDir)

	tmpDir, err := t.MkDir("test")

	c.Assert(tmpDir, Not(Equals), "")
	c.Assert(err, IsNil)
	c.Assert(fsutil.IsExist(tmpDir), Equals, true)
	c.Assert(fsutil.IsDir(tmpDir), Equals, true)
	c.Assert(fsutil.IsReadable(tmpDir), Equals, true)
	c.Assert(fsutil.IsWritable(tmpDir), Equals, true)
	c.Assert(fsutil.GetMode(tmpDir), Equals, os.FileMode(0750))

	t.Clean()

	c.Assert(fsutil.IsExist(tmpDir), Equals, false)
}

func (ts *TmpSuite) TestMkFile(c *C) {
	t, err := NewTemp(ts.TempDir)

	c.Assert(err, IsNil)
	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, ts.TempDir)

	tmpFd, tmpFile, err := t.MkFile("test")

	c.Assert(tmpFd, NotNil)
	c.Assert(err, IsNil)
	c.Assert(tmpFile, Not(Equals), "")
	c.Assert(fsutil.IsExist(tmpFile), Equals, true)
	c.Assert(fsutil.IsRegular(tmpFile), Equals, true)
	c.Assert(fsutil.IsReadable(tmpFile), Equals, true)
	c.Assert(fsutil.IsWritable(tmpFile), Equals, true)
	c.Assert(fsutil.GetMode(tmpFile), Equals, os.FileMode(0640))

	t.Clean()

	c.Assert(fsutil.IsExist(tmpFile), Equals, false)
}

func (ts *TmpSuite) TestMkName(c *C) {
	t, err := NewTemp(ts.TempDir)

	c.Assert(err, IsNil)
	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, ts.TempDir)

	c.Assert(t.MkName(), Not(Equals), "")
	c.Assert(t.MkName("1234"), Not(Equals), "")

	ln := len(ts.TempDir + "/")

	c.Assert(t.MkName("1234.json")[ln+14:], Equals, "1234.json")
}
