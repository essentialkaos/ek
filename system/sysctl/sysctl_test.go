package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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

func (s *SysctlSuite) TestBasic(c *C) {
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
		vs, err := Get("net.ipv4.conf.all.forwarding")
		c.Assert(err, IsNil)
		c.Assert(vs, Equals, "1")
		vi, err := GetI("net.ipv4.conf.all.forwarding")
		c.Assert(err, IsNil)
		c.Assert(vi, Equals, int(1))
		vi64, err := GetI64("net.ipv4.conf.all.forwarding")
		c.Assert(err, IsNil)
		c.Assert(vi64, Equals, int64(1))
	}
}

func (s *SysctlSuite) TestBasicErrors(c *C) {
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
}

func (s *SysctlSuite) TestOSErrors(c *C) {
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
