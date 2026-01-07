//go:build linux || freebsd

package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type InitSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&InitSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *InitSuite) TestSystemdEnabled(c *C) {
	c.Assert(parseSystemdEnabledOutput("enabled\r\n"), Equals, true)
	c.Assert(parseSystemdEnabledOutput("enabled"), Equals, true)
	c.Assert(parseSystemdEnabledOutput(""), Equals, false)
	c.Assert(parseSystemdEnabledOutput("Failed to get unit file state for unknown.service: No such file or directory"), Equals, false)
}

func (s *InitSuite) TestUpstartEnabled(c *C) {
	data1 := `# Ignore this comment
start on runlevel [2]
stop on runlevel [5]

pre-start script

bash << "EOF"
  mkdir -p /var/log/myapp
EOF

end script`

	data2 := `# Ignore this comment

pre-start script

bash << "EOF"
  mkdir -p /var/log/myapp
EOF

end script`

	tmpDir := c.MkDir()

	os.WriteFile(tmpDir+"/1.conf", []byte(data1), 0644)
	os.WriteFile(tmpDir+"/2.conf", []byte(data2), 0644)

	ok, err := parseUpstartEnabledData(tmpDir + "/0.conf")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't read service unit file`)
	c.Assert(ok, Equals, false)

	ok, err = parseUpstartEnabledData(tmpDir + "/1.conf")

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, true)

	ok, err = parseUpstartEnabledData(tmpDir + "/2.conf")

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, false)
}

func (s *InitSuite) TestSysvEnabled(c *C) {
	d1 := "myapp         0:off 1:off 2:on  3:on  4:on  5:on  6:off\r\n"
	d2 := "myapp         0:off 1:off 2:off 3:off 4:off 5:off 6:off\r\n"
	d3 := "myapp\r\n"

	ok, err := parseSysvEnabledOutput(d1)

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, true)

	ok, err = parseSysvEnabledOutput(d2)

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, false)

	ok, err = parseSysvEnabledOutput(d3)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse chkconfig output`)
	c.Assert(ok, Equals, false)

	ok, err = parseSysvEnabledOutput("")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse chkconfig output`)
	c.Assert(ok, Equals, false)
}

func (s *InitSuite) TestSystemdStatus(c *C) {
	d := "\r\n\r\n\n\n"
	ok, err := parseSystemdStatusOutput("myapp", d)

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Can't parse systemd output")
	c.Assert(ok, Equals, false)

	d = "LoadState=not-found\nActiveState=inactive\n"
	ok, err = parseSystemdStatusOutput("myapp", d)

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Unit myapp could not be found")
	c.Assert(ok, Equals, false)

	d = "LoadState=loaded\nActiveState=failed\n"
	ok, err = parseSystemdStatusOutput("myapp", d)

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, false)

	d = "LoadState=loaded\nActiveState=failed\n"
	ok, err = parseSystemdStatusOutput("myapp", d)

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, false)

	d = "LoadState=loaded\nActiveState=inactive\n"
	ok, err = parseSystemdStatusOutput("myapp", d)

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, false)

	d = "LoadState=loaded\nActiveState=active\n"
	ok, err = parseSystemdStatusOutput("myapp", d)

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, true)
}

func (s *InitSuite) TestUpstartStatus(c *C) {
	ok, err := parseUpstartStatusOutput("\r\n")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse upstart output`)
	c.Assert(ok, Equals, false)

	ok, err = parseUpstartStatusOutput("assdas ad asd asd\r\n")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse upstart output`)
	c.Assert(ok, Equals, false)

	ok, err = parseUpstartStatusOutput("myapp stop/waiting\r\n")

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, false)

	ok, err = parseUpstartStatusOutput("myapp start/running\r\n")

	c.Assert(err, IsNil)
	c.Assert(ok, Equals, true)
}
