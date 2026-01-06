package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type SystemSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SystemSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SystemSuite) TestUptime(c *C) {
	uptime, err := GetUptime()

	c.Assert(err, IsNil)
	c.Assert(uptime, Not(Equals), 0)
}

func (s *SystemSuite) TestLoadAvg(c *C) {
	la, err := GetLA()

	c.Assert(err, IsNil)
	c.Assert(la, NotNil)
}

func (s *SystemSuite) TestUser(c *C) {
	c.Assert(IsUserExist("root"), Equals, true)
	c.Assert(IsUserExist("_unknown_"), Equals, false)
	c.Assert(IsGroupExist("wheel"), Equals, true)
	c.Assert(IsGroupExist("_unknown_"), Equals, false)
}

func (s *SystemSuite) TestSystemInfo(c *C) {
	sysInfo, err := GetSystemInfo()

	c.Assert(err, IsNil)
	c.Assert(sysInfo, NotNil)

	c.Assert(sysInfo.Hostname, Not(Equals), "")
	c.Assert(sysInfo.OS, Not(Equals), "")
	c.Assert(sysInfo.ID, Not(Equals), "")
	c.Assert(sysInfo.Kernel, Not(Equals), "")
	c.Assert(sysInfo.Arch, Not(Equals), "")
	c.Assert(sysInfo.ArchName, Not(Equals), "")
	c.Assert(sysInfo.ArchBits, Not(Equals), 0)

	osInfo, err := GetOSInfo()

	c.Assert(err, IsNil)
	c.Assert(osInfo, NotNil)

	c.Assert(osInfo.Name, Equals, "macOS")
	c.Assert(osInfo.Version, Not(Equals), "")
	c.Assert(osInfo.Build, Not(Equals), "")
}

func (s *SystemSuite) TestCPUInfo(c *C) {
	info, err := GetCPUInfo()

	c.Assert(err, IsNil)
	c.Assert(info, NotNil)
}

func (s *SystemSuite) TestMemUsage(c *C) {
	mem, err := GetMemUsage()

	c.Assert(err, IsNil)
	c.Assert(mem, NotNil)
}
