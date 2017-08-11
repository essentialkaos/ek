// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	. "pkg.re/check.v1"
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

	procUptimeFile = ""

	uptime, err = GetUptime()

	c.Assert(err, NotNil)
	c.Assert(uptime, Equals, uint64(0))

	procUptimeFile = s.CreateTestFile(c, "CORRUPT")

	uptime, err = GetUptime()

	c.Assert(err, NotNil)
	c.Assert(uptime, Equals, uint64(0))
}

func (s *SystemSuite) TestLoadAvg(c *C) {
	la, err := GetLA()

	c.Assert(err, IsNil)
	c.Assert(la, NotNil)

	procLoadAvgFile = ""

	la, err = GetLA()

	c.Assert(err, NotNil)
	c.Assert(la, IsNil)

	procLoadAvgFile = s.CreateTestFile(c, "CORRUPT")

	la, err = GetLA()

	c.Assert(err, NotNil)
	c.Assert(la, IsNil)
}

func (s *SystemSuite) TestCPU(c *C) {
	cpu, err := GetCPUInfo()

	c.Assert(err, IsNil)
	c.Assert(cpu, NotNil)

	procStatFile = ""

	cpu, err = GetCPUInfo()

	c.Assert(err, NotNil)
	c.Assert(cpu, IsNil)

	procStatFile = s.CreateTestFile(c, "CORRUPT")

	cpu, err = GetCPUInfo()

	c.Assert(err, NotNil)
	c.Assert(cpu, IsNil)

	procStatFile = "/proc/stat"
}

func (s *SystemSuite) TestMemory(c *C) {
	mem, err := GetMemInfo()

	c.Assert(err, IsNil)
	c.Assert(mem, NotNil)

	procMemInfoFile = ""

	mem, err = GetMemInfo()

	c.Assert(err, NotNil)
	c.Assert(mem, IsNil)

	procMemInfoFile = s.CreateTestFile(c, "MemTotal:")

	mem, err = GetMemInfo()

	c.Assert(err, IsNil)

	procMemInfoFile = s.CreateTestFile(c, "MemTotal: ABC! kB")

	mem, err = GetMemInfo()

	c.Assert(err, NotNil)
	c.Assert(mem, IsNil)
}

func (s *SystemSuite) TestNet(c *C) {
	net, err := GetInterfacesInfo()

	c.Assert(err, IsNil)
	c.Assert(net, NotNil)

	_, _, err = GetNetworkSpeed(time.Second)

	c.Assert(err, IsNil)

	in, out := CalculateNetworkSpeed(
		map[string]*InterfaceInfo{"eth0": {0, 0, 0, 0}},
		map[string]*InterfaceInfo{"eth0": {0, 0, 0, 0}},
		time.Second,
	)

	c.Assert(in, Equals, uint64(0))
	c.Assert(out, Equals, uint64(0))

	procNetFile = ""

	net, err = GetInterfacesInfo()

	c.Assert(err, NotNil)
	c.Assert(net, IsNil)

	procNetFile = s.CreateTestFile(c, "CORRUPT")

	net, err = GetInterfacesInfo()

	c.Assert(err, NotNil)
	c.Assert(net, IsNil)

	_, _, err = GetNetworkSpeed(time.Second)

	c.Assert(err, NotNil)
}

func (s *SystemSuite) TestFS(c *C) {
	fs, err := GetFSInfo()

	c.Assert(err, IsNil)
	c.Assert(fs, NotNil)

	util, err := GetIOUtil(time.Second)

	c.Assert(err, IsNil)
	c.Assert(util, NotNil)

	mtabFile = ""

	fs, err = GetFSInfo()

	c.Assert(err, NotNil)
	c.Assert(fs, IsNil)

	mtabFile = s.CreateTestFile(c, "/CORRUPT")

	fs, err = GetFSInfo()

	c.Assert(err, NotNil)
	c.Assert(fs, IsNil)

	mtabFile = s.CreateTestFile(c, "/CORRUPT 0 0 0")

	fs, err = GetFSInfo()

	c.Assert(err, NotNil)
	c.Assert(fs, IsNil)

	procDiscStatsFile = ""

	stats, err := GetIOStats()

	c.Assert(err, NotNil)
	c.Assert(stats, IsNil)

	procDiscStatsFile = s.CreateTestFile(c, "CORRUPT")

	stats, err = GetIOStats()

	c.Assert(err, NotNil)
	c.Assert(stats, IsNil)

	fs, err = GetFSInfo()

	c.Assert(err, NotNil)
	c.Assert(fs, IsNil)

	util, err = GetIOUtil(time.Millisecond)

	c.Assert(err, NotNil)
	c.Assert(util, IsNil)

	procStatFile = ""
	mtabFile = "/etc/mtab"
	procDiscStatsFile = "/proc/diskstats"

	mtabFile = s.CreateTestFile(c, "/dev/abc1 / ext4 rw 0 0")

	fs, err = GetFSInfo()

	c.Assert(err, IsNil)
	c.Assert(fs, NotNil)

	util = CalculateIOUtil(
		map[string]*IOStats{"abc": {IOMs: 10}},
		map[string]*IOStats{"abc": {IOMs: 1840}},
		time.Minute,
	)

	c.Assert(util, NotNil)
	c.Assert(util["abc"], Equals, 3.05)

	procStatFile = "/proc/stat"
}

func (s *SystemSuite) TestUser(c *C) {
	user, err := CurrentUser()

	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	appendRealUserInfo(user)

	c.Assert(user.IsRoot(), Equals, false)
	c.Assert(user.IsSudo(), Equals, false)
	c.Assert(user.GroupList(), NotNil)

	user, err = CurrentUser()

	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	sess, err := Who()

	c.Assert(err, IsNil)
	c.Assert(sess, NotNil)

	user, err = LookupUser("")

	c.Assert(err, NotNil)
	c.Assert(user, IsNil)

	group, err := LookupGroup("root")

	c.Assert(err, IsNil)
	c.Assert(group, NotNil)

	group, err = LookupGroup("")

	c.Assert(err, NotNil)
	c.Assert(group, IsNil)

	c.Assert(IsUserExist("root"), Equals, true)
	c.Assert(IsUserExist("_UNKNOWN_"), Equals, false)
	c.Assert(IsGroupExist("root"), Equals, true)
	c.Assert(IsGroupExist("_UNKNOWN_"), Equals, false)

	c.Assert(CurrentTTY(), Not(Equals), "")

	uid, ok := getTDOwnerID()

	c.Assert(uid, Not(Equals), -1)
	c.Assert(ok, Equals, true)

	os.Setenv("SUDO_USER", "testuser")
	os.Setenv("SUDO_UID", "1234")
	os.Setenv("SUDO_GID", "1234")

	username, uid, gid := getRealUserFromEnv()

	c.Assert(username, Equals, "testuser")
	c.Assert(uid, Equals, 1234)
	c.Assert(gid, Equals, 1234)

	_, _, err = getGroupInfo("_UNKNOWN_")

	c.Assert(err, NotNil)

	_, err = getOwner("")

	c.Assert(err, NotNil)

	_, err = getSessionInfo("ABC")

	c.Assert(err, NotNil)

	n, _ := fixCount(-100, nil)

	c.Assert(n, Equals, 0)

	ptsDir = "/not_exist"

	sess, err = Who()

	c.Assert(err, IsNil)
	c.Assert(sess, HasLen, 0)
}

func (s *SystemSuite) TestInternal(c *C) {
	tmpDir := c.MkDir()
	tmpFile1 := tmpDir + "/test1.file"
	tmpFile2 := tmpDir + "/test2.file"

	if ioutil.WriteFile(tmpFile1, []byte("TEST\n1234"), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	if ioutil.WriteFile(tmpFile2, []byte(""), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	content, err := readFileContent(tmpFile1)

	c.Assert(err, IsNil)
	c.Assert(content, NotNil)
	c.Assert(content, HasLen, 2)

	content, err = readFileContent(tmpFile2)

	c.Assert(err, NotNil)
	c.Assert(content, IsNil)

	content, err = readFileContent("/not_exist")

	c.Assert(err, NotNil)
	c.Assert(content, IsNil)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SystemSuite) CreateTestFile(c *C, data string) string {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"

	if ioutil.WriteFile(tmpFile, []byte(data), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	return tmpFile
}
