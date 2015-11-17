package tmp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "gopkg.in/check.v1"
	"os"
	"testing"

	"github.com/essentialkaos/ek/fsutil"
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
	t := NewTemp()

	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, "/tmp")

	t.Clean()
}

func (ts *TmpSuite) TestErrors(c *C) {
	t := NewTemp("/")

	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, "/")

	tmpDir, err := t.MkDir("test")

	c.Assert(tmpDir, Equals, "")
	c.Assert(err, NotNil)

	tmpFd, tmpFile, err := t.MkFile("test")

	c.Assert(tmpFd, IsNil)
	c.Assert(tmpFile, Equals, "")
	c.Assert(err, NotNil)
}

func (ts *TmpSuite) TestMkDir(c *C) {
	t := NewTemp(ts.TempDir)

	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, ts.TempDir)

	tmpDir, err := t.MkDir("test")

	c.Assert(tmpDir, Not(Equals), "")
	c.Assert(err, IsNil)
	c.Assert(fsutil.IsExist(tmpDir), Equals, true)
	c.Assert(fsutil.IsDir(tmpDir), Equals, true)
	c.Assert(fsutil.IsReadable(tmpDir), Equals, true)
	c.Assert(fsutil.IsWritable(tmpDir), Equals, true)
	c.Assert(fsutil.GetPerm(tmpDir), Equals, os.FileMode(0750))

	t.Clean()

	c.Assert(fsutil.IsExist(tmpDir), Equals, false)
}

func (ts *TmpSuite) TestMkFile(c *C) {
	t := NewTemp(ts.TempDir)

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
	c.Assert(fsutil.GetPerm(tmpFile), Equals, os.FileMode(0640))

	t.Clean()

	c.Assert(fsutil.IsExist(tmpFile), Equals, false)
}

func (ts *TmpSuite) TestMkName(c *C) {
	t := NewTemp(ts.TempDir)

	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, ts.TempDir)

	c.Assert(t.MkName(), Not(Equals), "")
	c.Assert(t.MkName("1234"), Not(Equals), "")

	ln := len(ts.TempDir + "/_1234")

	c.Assert(t.MkName("1234")[:ln], Equals, ts.TempDir+"/_1234")
}
