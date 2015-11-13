package tmp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"ek/fsutil"
	. "gopkg.in/check.v1"
	"os"
	"testing"
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

func (ts *TmpSuite) TestMkDir(c *C) {
	t := NewTemp(ts.TempDir)

	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, ts.TempDir)

	tn, err := t.MkDir("test")

	c.Assert(tn, Not(Equals), "")
	c.Assert(err, IsNil)
	c.Assert(fsutil.IsExist(tn), Equals, true)
	c.Assert(fsutil.IsDir(tn), Equals, true)
	c.Assert(fsutil.IsReadable(tn), Equals, true)
	c.Assert(fsutil.IsWritable(tn), Equals, true)
	c.Assert(fsutil.GetPerm(tn), Equals, os.FileMode(0750))

	t.Clean()

	c.Assert(fsutil.IsExist(tn), Equals, false)
}

func (ts *TmpSuite) TestMkFile(c *C) {
	t := NewTemp(ts.TempDir)

	c.Assert(t, NotNil)
	c.Assert(t.Dir, Equals, ts.TempDir)

	tf, tn, err := t.MkFile("test")

	c.Assert(tf, NotNil)
	c.Assert(err, IsNil)
	c.Assert(tn, Not(Equals), "")
	c.Assert(fsutil.IsExist(tn), Equals, true)
	c.Assert(fsutil.IsRegular(tn), Equals, true)
	c.Assert(fsutil.IsReadable(tn), Equals, true)
	c.Assert(fsutil.IsWritable(tn), Equals, true)
	c.Assert(fsutil.GetPerm(tn), Equals, os.FileMode(0640))

	t.Clean()

	c.Assert(fsutil.IsExist(tn), Equals, false)
}
