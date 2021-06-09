// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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

	procFS = s.CreateFakeProcFS(c, "10000", "task", "AABBCC")

	_, err = GetTree(10000)
	c.Assert(err, NotNil)

	procFS = "/proc"
}

func (s *ProcessSuite) TestGetTreeAux(c *C) {
	_, err := readProcessInfo("/_UNKNOWN_", "10000", map[int]string{})
	c.Assert(err, NotNil)

	_, err = readProcessInfo("/_UNKNOWN_", "ABCD", map[int]string{})
	c.Assert(err, NotNil)

	c.Assert(isPID(""), Equals, false)

	_, err = getProcessUser(9999, map[int]string{})
	c.Assert(err, NotNil)

	p1, p2 := getParentPIDs("/_UNKNOWN_")
	c.Assert(p1, Equals, -1)
	c.Assert(p2, Equals, -1)

	processListToTree([]*ProcessInfo{
		&ProcessInfo{Parent: -1},
	}, 0)

	pidDir := s.CreateFakeProcFS(c, "10000", "status", "Tgid: \nPPid: \n")

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

	procFS = s.CreateFakeProcFS(c, "10000", "stat", "AABBCC")

	info, err = GetInfo(10000)
	c.Assert(info, IsNil)
	c.Assert(err, NotNil)

	procFS = "/proc"

	_, err = parseStatData("AA BB CC")
	c.Assert(err, NotNil)

	_, err = parseSampleData("0 0 0 0 0 0 0 0 0 0 0 0 0 X 0 0 0 0")
	c.Assert(err, NotNil)
}

func (s *ProcessSuite) TestGetSample(c *C) {
	_, err := GetSample(66000)

	c.Assert(err, NotNil)

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

	c.Assert(info, IsNil)
	c.Assert(err, NotNil)

	info, err = GetMemInfo(1)

	c.Assert(info, NotNil)
	c.Assert(err, IsNil)

	procFS = s.CreateFakeProcFS(c, "10000", "status", "VmPeak: AAA")

	info, err = GetMemInfo(10000)
	c.Assert(info, IsNil)
	c.Assert(err, NotNil)

	procFS = s.CreateFakeProcFS(c, "10000", "status", "VmPeak:         0 kB\nVmSize:         0 kB\n")

	info, err = GetMemInfo(10000)
	c.Assert(info, IsNil)
	c.Assert(err, NotNil)

	procFS = "/proc"
}

func (s *ProcessSuite) TestGetMountInfo(c *C) {
	info, err := GetMountInfo(66000)

	c.Assert(info, IsNil)
	c.Assert(err, NotNil)

	info, err = GetMountInfo(1)

	c.Assert(info, NotNil)
	c.Assert(err, IsNil)

	_, _, err = parseStDevValue("AA:11")
	c.Assert(err, NotNil)

	_, _, err = parseStDevValue("11:AA")
	c.Assert(err, NotNil)

	_, err = parseMountInfoLine("AA AA")
	c.Assert(err, NotNil)

	procFS = s.CreateFakeProcFS(c, "10000", "mountinfo", "AA AA AA")

	info, err = GetMountInfo(10000)

	c.Assert(info, IsNil)
	c.Assert(err, NotNil)

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
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ProcessSuite) CreateFakeProcFS(c *C, pid, file, data string) string {
	tmpDir := c.MkDir()
	os.Mkdir(tmpDir+"/"+pid, 0755)
	pfsFile := tmpDir + "/" + pid + "/" + file

	if ioutil.WriteFile(pfsFile, []byte(data), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	return tmpDir
}
