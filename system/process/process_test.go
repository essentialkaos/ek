//go:build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ProcessSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ProcessSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ProcessSuite) TestGetTree(c *C) {
	tree, err := GetTree(66000)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Process with PID \[66000\] doesn't exist`)
	c.Assert(tree, IsNil)

	tree, err = GetTree(os.Getpid())

	c.Assert(err, IsNil)
	c.Assert(tree, NotNil)

	procFS = s.CreateFakeProcFS(c, "10000", "task", "AABBCC")

	_, err = GetTree(10000)

	c.Assert(err, NotNil)

	procFS = "/proc"
}

func (s *ProcessSuite) TestGetTreeAux(c *C) {
	_, err := readProcessInfo("/_unknown_", "ABCD", map[int]string{})

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `strconv.Atoi: parsing "ABCD": invalid syntax`)

	c.Assert(isPID(""), Equals, false)

	_, err = getProcessUser(9999, map[int]string{})

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `User with name/ID 9999 does not exist`)

	p1, p2 := getParentPIDs("/_unknown_")

	c.Assert(p1, Equals, -1)
	c.Assert(p2, Equals, -1)

	processListToTree([]*ProcessInfo{&ProcessInfo{Parent: -1}}, 0)

	pidDir := s.CreateFakeProcFS(c, "10000", "status", "Tgid: \nPPid: \n")
	tree, err := readProcessInfo(pidDir, "10000", map[int]string{})

	c.Assert(tree, IsNil)
	c.Assert(err, IsNil)

	p1, p2 = getParentPIDs(pidDir + "/10000")

	c.Assert(p1, Equals, -1)
	c.Assert(p2, Equals, -1)

	pidDir = s.CreateFakeProcFS(c, "10000", "status", "Tgid: X\nPPid: X\n")

	p1, p2 = getParentPIDs(pidDir + "/10000")

	c.Assert(p1, Equals, -1)
	c.Assert(p2, Equals, -1)
}

func (s *ProcessSuite) TestGetList(c *C) {
	procs, err := GetList()

	c.Assert(err, IsNil)
	c.Assert(procs, NotNil)
	c.Assert(procs, Not(HasLen), 0)
}

func (s *ProcessSuite) TestGetInfo(c *C) {
	info, err := GetInfo(66000)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /proc/66000/stat: no such file or directory`)
	c.Assert(info, IsNil)

	info, err = GetInfo(1)

	c.Assert(err, IsNil)
	c.Assert(info, NotNil)

	procFS = s.CreateFakeProcFS(c, "10000", "stat", "AABBCC")

	info, err = GetInfo(10000)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse stat file for given process`)
	c.Assert(info, IsNil)

	procFS = "/proc"

	_, err = parseStatData("AA BB CC")
	c.Assert(err, NotNil)

	_, err = parseSampleData("0 0 0 0 0 0 0 0 0 0 0 0 0 X 0 0 0 0")
	c.Assert(err, NotNil)
}

func (s *ProcessSuite) TestGetSample(c *C) {
	_, err := GetSample(66000)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open .*/66000/stat: no such file or directory`)

	_, err = GetSample(1)

	c.Assert(err, IsNil)

	procFS = s.CreateFakeProcFS(c, "10000", "stat", "AABBCC")

	_, err = GetSample(10000)
	c.Assert(err, NotNil)

	procFS = "/proc"
}

func (s *ProcessSuite) TestInfoToSample(c *C) {
	pi := &ProcInfo{UTime: 10, STime: 1, CUTime: 1, CSTime: 1}
	ps := pi.ToSample()

	c.Assert(ps, Equals, ProcSample(13))
}

func (s *ProcessSuite) TestGetMemInfo(c *C) {
	info, err := GetMemInfo(66000)

	c.Assert(err, NotNil)
	c.Assert(info, IsNil)

	info, err = GetMemInfo(1)

	c.Assert(err, IsNil)
	c.Assert(info, NotNil)

	procFS = s.CreateFakeProcFS(c, "10000", "status", "VmPeak: AAA")

	info, err = GetMemInfo(10000)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse status file for given process`)
	c.Assert(info, IsNil)

	procFS = s.CreateFakeProcFS(c, "10000", "status", "VmPeak:         0 kB\nVmSize:         0 kB\n")

	info, err = GetMemInfo(10000)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse status file for given process`)
	c.Assert(info, IsNil)

	procFS = "/proc"
}

func (s *ProcessSuite) TestGetMountInfo(c *C) {
	info, err := GetMountInfo(66000)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open .*/66000/mountinfo: no such file or directory`)
	c.Assert(info, IsNil)

	info, err = GetMountInfo(1)

	c.Assert(err, IsNil)
	c.Assert(info, NotNil)

	_, _, err = parseStDevValue("AA:11")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse field StDevMajor: strconv.ParseUint: parsing "AA": invalid syntax`)

	_, _, err = parseStDevValue("11:AA")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse field StDevMinor: strconv.ParseUint: parsing "AA": invalid syntax`)

	_, err = parseMountInfoLine("AA AA")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse field MountID: strconv.ParseUint: parsing "AA": invalid syntax`)

	procFS = s.CreateFakeProcFS(c, "10000", "mountinfo", "AA AA AA")

	info, err = GetMountInfo(10000)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse field MountID: strconv.ParseUint: parsing "AA": invalid syntax`)
	c.Assert(info, IsNil)

	procFS = "/proc"
}

func (s *ProcessSuite) TestCalculateCPUUsage(c *C) {
	s1 := ProcSample(13)
	s2 := ProcSample(66)

	c.Assert(CalculateCPUUsage(s1, s2, time.Second), Equals, 53.0)
}

func (s *ProcessSuite) TestCPUPriority(c *C) {
	pid := os.Getpid()

	pri, ni, err := GetCPUPriority(pid)

	c.Assert(err, IsNil)
	c.Assert(pri, Equals, 20)
	c.Assert(ni, Equals, 0)

	err = SetCPUPriority(pid, 2)

	c.Assert(err, IsNil)

	pri, ni, err = GetCPUPriority(pid)

	c.Assert(err, IsNil)
	c.Assert(pri, Equals, 22)
	c.Assert(ni, Equals, 2)

	_, _, err = GetCPUPriority(12000000)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `no such process`)
}

func (s *ProcessSuite) TestIOPriority(c *C) {
	pid := os.Getpid()

	err := SetIOPriority(pid, PRIO_CLASS_NONE, 0)

	c.Assert(err, IsNil)

	class, classdata, err := GetIOPriority(pid)

	c.Assert(err, IsNil)
	c.Assert(class, Equals, PRIO_CLASS_NONE)
	c.Assert(classdata, Equals, 0)

	err = SetIOPriority(pid, PRIO_CLASS_BEST_EFFORT, 5)

	c.Assert(err, IsNil)

	class, classdata, err = GetIOPriority(pid)

	c.Assert(err, IsNil)
	c.Assert(class, Equals, PRIO_CLASS_BEST_EFFORT)
	c.Assert(classdata, Equals, 5)

	_, _, err = GetIOPriority(999999)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `no such process`)

	err = SetIOPriority(999999, PRIO_CLASS_BEST_EFFORT, 5)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `no such process`)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ProcessSuite) CreateFakeProcFS(c *C, pid, file, data string) string {
	tmpDir := c.MkDir()
	os.Mkdir(tmpDir+"/"+pid, 0755)
	pfsFile := tmpDir + "/" + pid + "/" + file

	if os.WriteFile(pfsFile, []byte(data), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	return tmpDir
}
