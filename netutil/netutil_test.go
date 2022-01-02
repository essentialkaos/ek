//go:build linux
// +build linux

package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type NetUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&NetUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *NetUtilSuite) TestCommon(c *C) {
	if runtime.GOOS == "linux" {
		c.Assert(getDefaultRouteInterface(), Not(Equals), "")

		procRouteFile = "/tmp/not-exist"

		c.Assert(getDefaultRouteInterface(), Equals, "")

		procRouteFile = s.CreateTestFile(c, "Iface   Destination     Gateway         Flags   RefCnt  Use     Metric  Mask            MTU     Window  IRTT\neth0    0070652E        00000000        0001    0       0       0       00F0FFFF        0       0       0")

		c.Assert(getDefaultRouteInterface(), Equals, "")
	}
}

func (s *NetUtilSuite) TestGetIP(c *C) {
	c.Assert(GetIP(), Not(Equals), "")
}

func (s *NetUtilSuite) TestGetIP6(c *C) {
	if os.Getenv("CI") == "" {
		c.Assert(GetIP6(), Not(Equals), "")
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *NetUtilSuite) CreateTestFile(c *C, data string) string {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"

	if ioutil.WriteFile(tmpFile, []byte(data), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	return tmpFile
}
