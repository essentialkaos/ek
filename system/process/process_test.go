// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"
	"time"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ProcessSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ProcessSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ProcessSuite) TestGetTree(c *C) {
	tree, err := GetTree(66000)

	c.Assert(tree, IsNil)
	c.Assert(err, NotNil)

	tree, err = GetTree(1)

	c.Assert(tree, NotNil)
	c.Assert(err, IsNil)
}

func (s *ProcessSuite) TestGetList(c *C) {
	procs, err := GetList()

	c.Assert(procs, NotNil)
	c.Assert(procs, Not(HasLen), 0)
	c.Assert(err, IsNil)
}

func (s *ProcessSuite) TestGetInfo(c *C) {
	info, err := GetInfo(66000)

	c.Assert(info, IsNil)
	c.Assert(err, NotNil)

	info, err = GetInfo(1)

	c.Assert(info, NotNil)
	c.Assert(err, IsNil)
}

func (s *ProcessSuite) TestGetSample(c *C) {
	_, err := GetSample(66000)

	c.Assert(err, NotNil)

	_, err = GetSample(1)

	c.Assert(err, IsNil)
}

func (s *ProcessSuite) TestInfoToSample(c *C) {
	pi := &ProcInfo{UTime: 10, STime: 1, CUTime: 1, CSTime: 1}
	ps := pi.ToSample()

	c.Assert(ps, Equals, ProcSample(13))
}

func (s *ProcessSuite) TestGetMemInfo(c *C) {
	info, err := GetMemInfo(66000)

	c.Assert(info, IsNil)
	c.Assert(err, NotNil)

	info, err = GetMemInfo(1)

	c.Assert(info, NotNil)
	c.Assert(err, IsNil)
}

func (s *ProcessSuite) TestCalculateCPUUsage(c *C) {
	s1 := ProcSample(13)
	s2 := ProcSample(66)

	c.Assert(CalculateCPUUsage(s1, s2, time.Second), Equals, 53.0)
}
