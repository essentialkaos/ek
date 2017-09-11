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

	procUptimeFile = s.CreateTestFile(c, "CORRUPTED")

	uptime, err = GetUptime()

	c.Assert(err, NotNil)
	c.Assert(uptime, Equals, uint64(0))
}

func (s *SystemSuite) TestLoadAvg(c *C) {
	la, err := GetLA()

	c.Assert(err, IsNil)
	c.Assert(la, NotNil)

	procLoadAvgFile = s.CreateTestFile(c, "1.15 2.25 3.35 5/234 16354")

	la, err = GetLA()

	c.Assert(err, IsNil)
	c.Assert(la, NotNil)

	c.Assert(la.Min1, Equals, 1.15)
	c.Assert(la.Min5, Equals, 2.25)
	c.Assert(la.Min15, Equals, 3.35)
	c.Assert(la.RProc, Equals, 5)
	c.Assert(la.TProc, Equals, 234)

	procLoadAvgFile = ""

	la, err = GetLA()

	c.Assert(err, NotNil)
	c.Assert(la, IsNil)

	procLoadAvgFile = s.CreateTestFile(c, "CORRUPTED")

	la, err = GetLA()

	c.Assert(err, NotNil)
	c.Assert(la, IsNil)

	procLoadAvgFile = s.CreateTestFile(c, "1.15 2.25 3.35 5+234 16354")

	la, err = GetLA()

	c.Assert(err, NotNil)
	c.Assert(la, IsNil)
}

func (s *SystemSuite) TestCPU(c *C) {
	cpu, err := GetCPUStats()

	c.Assert(err, IsNil)
	c.Assert(cpu, NotNil)

	procStatFile = s.CreateTestFile(c, "cpu  10 11 12 13 14 15 16 17 0\ncpu0 0 0 0 0 0 0 0 0 0\ncpu1 0 0 0 0 0 0 0 0 0\n")

	cpu, err = GetCPUStats()

	c.Assert(err, IsNil)
	c.Assert(cpu, NotNil)
	c.Assert(cpu.Count, Equals, 2)
	c.Assert(cpu.User, Equals, uint64(10))
	c.Assert(cpu.Nice, Equals, uint64(11))
	c.Assert(cpu.System, Equals, uint64(12))
	c.Assert(cpu.Idle, Equals, uint64(13))
	c.Assert(cpu.Wait, Equals, uint64(14))
	c.Assert(cpu.IRQ, Equals, uint64(15))
	c.Assert(cpu.SRQ, Equals, uint64(16))
	c.Assert(cpu.Steal, Equals, uint64(17))

	c1 := &CPUStats{10, 10, 10, 2, 2, 2, 2, 0, 34, 32}
	c2 := &CPUStats{12, 12, 12, 3, 3, 3, 3, 0, 48, 32}

	cpuInfo := CalculateCPUInfo(c1, c2)

	c.Assert(cpuInfo, NotNil)
	c.Assert(cpuInfo.System, Equals, 14.285714285714285)
	c.Assert(cpuInfo.User, Equals, 14.285714285714285)
	c.Assert(cpuInfo.Nice, Equals, 14.285714285714285)
	c.Assert(cpuInfo.Wait, Equals, 7.142857142857142)
	c.Assert(cpuInfo.Idle, Equals, 7.142857142857142)
	c.Assert(cpuInfo.Average, Equals, 80.0)
	c.Assert(cpuInfo.Count, Equals, 32)

	procStatFile = ""

	cpu, err = GetCPUStats()

	c.Assert(err, NotNil)
	c.Assert(cpu, IsNil)

	procStatFile = s.CreateTestFile(c, "CORRUPTED")

	cpu, err = GetCPUStats()

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

	procMemInfoFile = s.CreateTestFile(c, "")

	mem, err = GetMemInfo()

	c.Assert(err, NotNil)
	c.Assert(mem, IsNil)

	procMemInfoFile = s.CreateTestFile(c, "MemTotal: ABC! kB")

	mem, err = GetMemInfo()

	c.Assert(err, NotNil)
	c.Assert(mem, IsNil)
}

func (s *SystemSuite) TestNet(c *C) {
	net, err := GetInterfacesInfo()

	c.Assert(err, IsNil)
	c.Assert(net, NotNil)

	procNetFile = s.CreateTestFile(c, "Inter-|   Receive                                                |  Transmit\n face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed\neth0: 144612532790 216320765    0    0    0     0          0         0 366397171405 154518846    0    0    0     0       0          0\n")

	net, err = GetInterfacesInfo()

	c.Assert(err, IsNil)
	c.Assert(net, NotNil)
	c.Assert(net["eth0"], NotNil)
	c.Assert(net["eth0"].ReceivedBytes, Equals, uint64(144612532790))
	c.Assert(net["eth0"].ReceivedPackets, Equals, uint64(216320765))
	c.Assert(net["eth0"].TransmittedBytes, Equals, uint64(366397171405))
	c.Assert(net["eth0"].TransmittedPackets, Equals, uint64(154518846))

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

	procNetFile = s.CreateTestFile(c, "CORRUPTED")

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

	mtabFile = s.CreateTestFile(c, "/CORRUPTED")

	fs, err = GetFSInfo()

	c.Assert(err, NotNil)
	c.Assert(fs, IsNil)

	mtabFile = s.CreateTestFile(c, "/CORRUPTED 0 0 0")

	fs, err = GetFSInfo()

	c.Assert(err, NotNil)
	c.Assert(fs, IsNil)

	procDiscStatsFile = ""

	stats, err := GetIOStats()

	c.Assert(err, NotNil)
	c.Assert(stats, IsNil)

	procDiscStatsFile = s.CreateTestFile(c, "CORRUPTED")

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
	// This test can fail on Travis because workers
	// doesn't have any active sessions
	if os.Getenv("TRAVIS") != "1" {
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

		groups := extractGroupsInfo("uid=10123(john) gid=10123(john) groups=10123(john),10200(admins),10201(developers)")

		c.Assert(groups[0].Name, Equals, "john")
		c.Assert(groups[0].GID, Equals, 10123)
		c.Assert(groups[2].Name, Equals, "developers")
		c.Assert(groups[2].GID, Equals, 10201)

		group, err = parseGetentGroupOutput("developers:*:10201:bob,john")

		c.Assert(err, IsNil)
		c.Assert(group, NotNil)
		c.Assert(group.Name, Equals, "developers")
		c.Assert(group.GID, Equals, 10201)

		user, err = parseGetentPasswdOutput("bob:*:10567:10567::/home/bob:/bin/zsh")

		c.Assert(err, IsNil)
		c.Assert(user, NotNil)
		c.Assert(user.Name, Equals, "bob")
		c.Assert(user.UID, Equals, 10567)
		c.Assert(user.GID, Equals, 10567)
		c.Assert(user.Comment, Equals, "")
		c.Assert(user.HomeDir, Equals, "/home/bob")
		c.Assert(user.Shell, Equals, "/bin/zsh")

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
}

func (s *SystemSuite) TestFieldParser(c *C) {
	data := "    abc \t\t 123     \t         ABC $"

	c.Assert(readField("", 0), Equals, "")
	c.Assert(readField(data, 0), Equals, "abc")
	c.Assert(readField(data, 1), Equals, "123")
	c.Assert(readField(data, 2), Equals, "ABC")
	c.Assert(readField(data, 3), Equals, "$")
	c.Assert(readField(data, 4), Equals, "")
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
