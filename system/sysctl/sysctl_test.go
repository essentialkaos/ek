package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"runtime"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type SysctlSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SysctlSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SysctlSuite) TestOne(c *C) {
	var p Param
	var err error

	if runtime.GOOS == "darwin" {
		p, err = Get("kern.maxproc")
	} else {
		p, err = Get("kernel.pid_max")
	}

	c.Assert(err, IsNil)
	c.Assert(p.IsEmpty(), Equals, false)
	c.Assert(p, Not(Equals), "0")

	vi, err := p.Int()
	c.Assert(err, IsNil)
	c.Assert(vi, Not(Equals), int(0))

	vi64, err := p.Int()
	c.Assert(err, IsNil)
	c.Assert(vi64, Not(Equals), int64(0))
}

func (s *SysctlSuite) TestOneErrors(c *C) {
	_, err := Get("")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `parameter name is empty`)

	_, err = Get("test")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `invalid parameter name "test"`)

	if runtime.GOOS == "darwin" {
		binary = "_unknown_"
		_, err := Get("kern.bootuuid")
		c.Assert(err, NotNil)
		binary = "sysctl"
	} else {
		procFS = "/_unknown_"
		_, err := Get("kernel.random.boot_id")
		c.Assert(err, NotNil)
		procFS = "/proc/sys"
	}
}

func (s *SysctlSuite) TestAll(c *C) {
	var p Param

	pp, err := All()

	c.Assert(err, IsNil)
	c.Assert(pp, Not(HasLen), 0)

	if runtime.GOOS == "darwin" {
		p = pp.Get("kern.maxproc")
	} else {
		p = pp.Get("kernel.pid_max")
	}

	c.Assert(p.IsEmpty(), Equals, false)
	c.Assert(p.String(), Not(Equals), "")

	vi, err := p.Int()
	c.Assert(err, IsNil)
	c.Assert(vi, Not(Equals), int(0))

	vi64, err := p.Int64()
	c.Assert(err, IsNil)
	c.Assert(vi64, Not(Equals), int64(0))

	if runtime.GOOS == "darwin" {
		p = pp.Get("_test_")
	} else {
		p = pp.Get("_test_")
	}

	c.Assert(p.IsEmpty(), Equals, true)
}

func (s *SysctlSuite) TestAllErrors(c *C) {
	binary = "_unknown_"
	_, err := All()
	c.Assert(err, NotNil)
	binary = "sysctl"
}
