package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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
	if runtime.GOOS == "darwin" {
		vs, err := Get("kern.job_control")
		c.Assert(err, IsNil)
		c.Assert(vs, Equals, "1")
		vi, err := GetI("kern.job_control")
		c.Assert(err, IsNil)
		c.Assert(vi, Equals, int(1))
		vi64, err := GetI64("kern.job_control")
		c.Assert(err, IsNil)
		c.Assert(vi64, Equals, int64(1))
	} else {
		vs, err := Get("vm.swappiness")
		c.Assert(err, IsNil)
		c.Assert(vs, Not(Equals), "0")
		vi, err := GetI("vm.swappiness")
		c.Assert(err, IsNil)
		c.Assert(vi, Not(Equals), int(0))
		vi64, err := GetI64("vm.swappiness")
		c.Assert(err, IsNil)
		c.Assert(vi64, Not(Equals), int64(0))
	}
}

func (s *SysctlSuite) TestOneErrors(c *C) {
	_, err := Get("")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Kernel parameter name cannot be empty`)

	_, err = Get("test")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Invalid parameter name "test"`)

	_, err = GetI("test")
	c.Assert(err, NotNil)
	_, err = GetI64("test")
	c.Assert(err, NotNil)

	_, err = Get("te st")
	c.Assert(err, NotNil)
	_, err = Get("te/st")
	c.Assert(err, NotNil)
	_, err = Get("te\tst")
	c.Assert(err, NotNil)
	_, err = Get("te\nst")
	c.Assert(err, NotNil)

	if runtime.GOOS == "darwin" {
		_, err = GetI("kern.bootuuid")
		c.Assert(err, NotNil)
		_, err = GetI64("kern.bootuuid")
		c.Assert(err, NotNil)
	} else {
		_, err = GetI("kernel.random.boot_id")
		c.Assert(err, NotNil)
		_, err = GetI64("kernel.random.boot_id")
		c.Assert(err, NotNil)
	}

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
	p, err := All()

	c.Assert(err, IsNil)
	c.Assert(p, Not(HasLen), 0)

	if runtime.GOOS == "darwin" {
		c.Assert(p.Get("kern.job_control"), Equals, "1")
		vi, err := p.GetI("kern.job_control")
		c.Assert(err, IsNil)
		c.Assert(vi, Equals, int(1))
		vi64, err := p.GetI64("kern.job_control")
		c.Assert(err, IsNil)
		c.Assert(vi64, Equals, int64(1))
	} else {
		c.Assert(p.Get("vm.swappiness"), Not(Equals), "")
		vi, err := p.GetI("vm.swappiness")
		c.Assert(err, IsNil)
		c.Assert(vi, Not(Equals), int(0))
		vi64, err := p.GetI64("vm.swappiness")
		c.Assert(err, IsNil)
		c.Assert(vi64, Not(Equals), int64(0))
	}
}

func (s *SysctlSuite) TestAllErrors(c *C) {
	var p Params

	c.Assert(p.Get("test"), Equals, "")

	p, err := All()

	c.Assert(err, IsNil)
	c.Assert(p, Not(HasLen), 0)

	if runtime.GOOS == "darwin" {
		_, err = p.GetI("kern.bootuuid")
		c.Assert(err, NotNil)
		_, err = p.GetI64("kern.bootuuid")
		c.Assert(err, NotNil)
	} else {
		_, err = p.GetI("kernel.random.boot_id")
		c.Assert(err, NotNil)
		_, err = p.GetI64("kernel.random.boot_id")
		c.Assert(err, NotNil)
	}

	binary = "_unknown_"
	_, err = All()
	c.Assert(err, NotNil)
	binary = "sysctl"
}
