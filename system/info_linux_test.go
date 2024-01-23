package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

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
	origProcLoadAvgFile := procLoadAvgFile

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

	procLoadAvgFile = "/_unknown_"

	la, err = GetLA()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_unknown_: no such file or directory`)
	c.Assert(la, IsNil)

	procLoadAvgFile = s.CreateTestFile(c, "CORRUPTED")

	la, err = GetLA()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse file .*/test.file`)
	c.Assert(la, IsNil)

	procLoadAvgFile = s.CreateTestFile(c, "1.15 2.25 3.35 5+234 16354")

	la, err = GetLA()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse field 3 in .*/test.file`)
	c.Assert(la, IsNil)

	procLoadAvgFile = origProcLoadAvgFile
}

func (s *SystemSuite) TestCPUUsage(c *C) {
	origProcStatFile := procStatFile

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

	usage, err := GetCPUUsage(time.Millisecond)

	c.Assert(err, IsNil)
	c.Assert(usage, NotNil)

	c1 := &CPUStats{10, 10, 10, 2, 2, 2, 2, 0, 34, 32}
	c2 := &CPUStats{12, 12, 12, 3, 3, 3, 3, 0, 48, 32}

	cpuInfo := CalculateCPUUsage(c1, c2)

	c.Assert(cpuInfo, NotNil)
	c.Assert(cpuInfo.System, Equals, 14.285714285714285)
	c.Assert(cpuInfo.User, Equals, 14.285714285714285)
	c.Assert(cpuInfo.Nice, Equals, 14.285714285714285)
	c.Assert(cpuInfo.Wait, Equals, 7.142857142857142)
	c.Assert(cpuInfo.Idle, Equals, 7.142857142857142)
	c.Assert(cpuInfo.Average, Equals, 80.0)
	c.Assert(cpuInfo.Count, Equals, 32)

	procStatFile = "/_unknown_"

	cpu, err = GetCPUStats()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_unknown_: no such file or directory`)
	c.Assert(cpu, IsNil)

	procStatFile = s.CreateTestFile(c, "CORRUPTED")

	cpu, err = GetCPUStats()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse file .*/test.file`)
	c.Assert(cpu, IsNil)

	procStatFile = origProcStatFile
}

func (s *SystemSuite) TestCPUInfoErrors(c *C) {
	origProcStatFile := procStatFile

	for i := 1; i <= 8; i++ {
		data := strings.Replace("cpu  1 2 3 4 5 6 7 8 0", strconv.Itoa(i), "AB", -1)

		procStatFile = s.CreateTestFile(c, data)

		cpu, err := GetCPUStats()

		c.Assert(err, NotNil)
		c.Assert(err, ErrorMatches, `Can't parse field [0-9]+ as unsigned integer in .*/test.file`)
		c.Assert(cpu, IsNil)
	}

	procStatFile = origProcStatFile
}

func (s *SystemSuite) TestCPUInfo(c *C) {
	origCpuInfoFile := cpuInfoFile

	cpuInfoFile = s.CreateTestFile(c, "vendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1204.199\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1199.908\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1199.908\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1199.908\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1199.908\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1199.792\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1206.634\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1281.085\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1273.431\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1199.908\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1273.199\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1200.024\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1260.443\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1195.385\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1259.283\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1199.908\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1199.908\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1203.387\ncache size	: 15360 KB\nphysical id	: 0\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1276.330\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1305.670\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1275.750\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1253.253\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1337.677\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\nvendor_id	: GenuineIntel\nmodel name	: Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz\ncpu MHz		: 1327.008\ncache size	: 15360 KB\nphysical id	: 1\nsiblings	: 12\ncpu cores	: 6\nflags		: mmx\n\n")

	info, err := GetCPUInfo()

	c.Assert(err, IsNil)
	c.Assert(info, NotNil)
	c.Assert(info, HasLen, 2)
	c.Assert(info[0].Vendor, Equals, "GenuineIntel")
	c.Assert(info[0].Model, Equals, "Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz")
	c.Assert(info[0].Cores, Equals, 6)
	c.Assert(info[0].Siblings, Equals, 12)
	c.Assert(info[0].CacheSize, Equals, uint64(15360*1024))
	c.Assert(info[0].Speed, DeepEquals, []float64{1204.199, 1199.908, 1199.908, 1199.908, 1199.908, 1199.792, 1260.443, 1195.385, 1259.283, 1199.908, 1199.908, 1203.387})
	c.Assert(info[1].Vendor, Equals, "GenuineIntel")
	c.Assert(info[1].Model, Equals, "Intel(R) Xeon(R) CPU E5-2420 0 @ 1.90GHz")
	c.Assert(info[1].Cores, Equals, 6)
	c.Assert(info[1].Siblings, Equals, 12)
	c.Assert(info[1].CacheSize, Equals, uint64(15360*1024))
	c.Assert(info[1].Speed, DeepEquals, []float64{1206.634, 1281.085, 1273.431, 1199.908, 1273.199, 1200.024, 1276.33, 1305.67, 1275.75, 1253.253, 1337.677, 1327.008})

	cpuInfoFile = s.CreateTestFile(c, "CORRUPTED")

	info, err = GetCPUInfo()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse cpuinfo file`)
	c.Assert(info, IsNil)

	cpuInfoFile = s.CreateTestFile(c, "cpu cores	: ABCD")

	info, err = GetCPUInfo()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `strconv.Atoi: parsing "ABCD": invalid syntax`)
	c.Assert(info, IsNil)

	cpuInfoFile = "/_unknown_"

	info, err = GetCPUInfo()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_unknown_: no such file or directory`)
	c.Assert(info, IsNil)

	cpuInfoFile = origCpuInfoFile
}

func (s *SystemSuite) TestCPUCount(c *C) {
	origCpuPossibleFile := cpuPossibleFile
	origCpuPresentFile := cpuPresentFile
	origCpuOnlineFile := cpuOnlineFile
	origCpuOfflineFile := cpuOfflineFile

	cpuPossibleFile = s.CreateTestFile(c, "0-127\n")
	cpuPresentFile = s.CreateTestFile(c, "0-3\n")
	cpuOnlineFile = s.CreateTestFile(c, "0-3\n")
	cpuOfflineFile = s.CreateTestFile(c, "4-127\n")

	info, err := GetCPUCount()

	c.Assert(err, IsNil)

	c.Assert(info.Possible, Equals, uint32(128))
	c.Assert(info.Present, Equals, uint32(4))
	c.Assert(info.Online, Equals, uint32(4))
	c.Assert(info.Offline, Equals, uint32(124))

	cpuOfflineFile = "/_cpuOfflineFile"
	_, err = GetCPUCount()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_cpuOfflineFile: no such file or directory`)

	cpuOnlineFile = "/_cpuOnlineFile"
	_, err = GetCPUCount()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_cpuOnlineFile: no such file or directory`)

	cpuPresentFile = "/_cpuPresentFile"
	_, err = GetCPUCount()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_cpuPresentFile: no such file or directory`)

	cpuPossibleFile = "/_cpuPossibleFile"
	_, err = GetCPUCount()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_cpuPossibleFile: no such file or directory`)

	cpuPossibleFile = origCpuPossibleFile
	cpuPresentFile = origCpuPresentFile
	cpuOnlineFile = origCpuOnlineFile
	cpuOfflineFile = origCpuOfflineFile
}

func (s *SystemSuite) TestMemUsage(c *C) {
	origProcMemInfoFile := procMemInfoFile

	procMemInfoFile = s.CreateTestFile(c, "MemTotal:       32653288 kB\nMemFree:          531664 kB\nMemAvailable:   31243272 kB\nBuffers:            7912 kB\nCached:         28485940 kB\nSwapCached:          889 kB\nActive:         15379084 kB\nInactive:       13340308 kB\nActive(anon):     143408 kB\nInactive(anon):   247048 kB\nActive(file):   15235676 kB\nInactive(file): 13093260 kB\nUnevictable:          16 kB\nMlocked:              16 kB\nSwapTotal:       4194300 kB\nSwapFree:        2739454 kB\nDirty:              6744 kB\nWriteback:             0 kB\nAnonPages:        225604 kB\nMapped:            46404 kB\nShmem:            164916 kB\nSlab:            3021828 kB\nSReclaimable:    2789624 kB\nSUnreclaim:       232204 kB\nKernelStack:        3696 kB\n")

	mem, err := GetMemUsage()

	c.Assert(err, IsNil)
	c.Assert(mem, NotNil)
	c.Assert(mem.MemTotal, Equals, uint64(1024*32653288))
	c.Assert(mem.MemFree, Equals, uint64(1024*31650224))
	c.Assert(mem.MemUsed, Equals, uint64(1024*1003064))
	c.Assert(mem.Buffers, Equals, uint64(1024*7912))
	c.Assert(mem.Cached, Equals, uint64(1024*28485940))
	c.Assert(mem.Active, Equals, uint64(1024*15379084))
	c.Assert(mem.Inactive, Equals, uint64(1024*13340308))
	c.Assert(mem.SwapTotal, Equals, uint64(1024*4194300))
	c.Assert(mem.SwapFree, Equals, uint64(1024*2739454))
	c.Assert(mem.SwapUsed, Equals, uint64(1024*1454846))
	c.Assert(mem.SwapCached, Equals, uint64(1024*889))
	c.Assert(mem.Dirty, Equals, uint64(1024*6744))
	c.Assert(mem.Shmem, Equals, uint64(1024*164916))
	c.Assert(mem.Slab, Equals, uint64(1024*3021828))
	c.Assert(mem.SReclaimable, Equals, uint64(1024*2789624))

	mem.MemTotal = 32 * 1024 * 1024 * 1024
	mem.MemUsed = 8 * 1024 * 1024 * 1024
	mem.SwapTotal = 16 * 1024 * 1024 * 1024
	mem.SwapUsed = 4 * 1024 * 1024 * 1024

	c.Assert(mem.MemUsedPerc(), Equals, 25.0)
	c.Assert(mem.SwapUsedPerc(), Equals, 25.0)

	procMemInfoFile = "/_unknown_"

	mem, err = GetMemUsage()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_unknown_: no such file or directory`)
	c.Assert(mem, IsNil)

	procMemInfoFile = s.CreateTestFile(c, "")

	mem, err = GetMemUsage()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse file .*/test.file`)
	c.Assert(mem, IsNil)

	procMemInfoFile = s.CreateTestFile(c, "MemTotal: ABC! kB")

	mem, err = GetMemUsage()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse file .*/test.file`)
	c.Assert(mem, IsNil)

	procMemInfoFile = origProcMemInfoFile
}

func (s *SystemSuite) TestNet(c *C) {
	origProcNetFile := procNetFile

	net, err := GetInterfacesStats()

	c.Assert(err, IsNil)
	c.Assert(net, NotNil)

	procNetFile = s.CreateTestFile(c, "Inter-|   Receive                                                |  Transmit\n face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed\neth0: 144612532790 216320765    0    0    0     0          0         0 366397171405 154518846    0    0    0     0       0          0\n")

	net, err = GetInterfacesStats()

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
		map[string]*InterfaceStats{"eth0": {0, 0, 0, 0}},
		map[string]*InterfaceStats{"eth0": {0, 0, 0, 0}},
		time.Second,
	)

	c.Assert(in, Equals, uint64(0))
	c.Assert(out, Equals, uint64(0))

	procNetFile = "/_unknown_"

	net, err = GetInterfacesStats()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_unknown_: no such file or directory`)
	c.Assert(net, IsNil)

	procNetFile = s.CreateTestFile(c, "CORRUPTED")

	net, err = GetInterfacesStats()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse file .*/test.file`)
	c.Assert(net, IsNil)

	_, _, err = GetNetworkSpeed(time.Second)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse file .*/test.file`)

	procNetFile = origProcNetFile
}

func (s *SystemSuite) TestFSUsage(c *C) {
	origMtabFile := mtabFile
	origProcStatFile := procStatFile
	origProcDiskStatsFile := procDiskStatsFile

	fs, err := GetFSUsage()

	c.Assert(err, IsNil)
	c.Assert(fs, NotNil)

	util, err := GetIOUtil(time.Second)

	c.Assert(err, IsNil)
	c.Assert(util, NotNil)

	mtabFile = "/_unknown_"

	fs, err = GetFSUsage()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_unknown_: no such file or directory`)
	c.Assert(fs, IsNil)

	mtabFile = s.CreateTestFile(c, "/CORRUPTED")

	fs, err = GetFSUsage()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `no such file or directory`)
	c.Assert(fs, IsNil)

	mtabFile = s.CreateTestFile(c, "/CORRUPTED 0 0 0")

	fs, err = GetFSUsage()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `no such file or directory`)
	c.Assert(fs, IsNil)

	procDiskStatsFile = "/_unknown_"

	stats, err := GetIOStats()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_unknown_: no such file or directory`)
	c.Assert(stats, IsNil)

	procDiskStatsFile = s.CreateTestFile(c, "CORRUPTED")

	stats, err = GetIOStats()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse field 3 as unsigned integer in .*/test.file`)
	c.Assert(stats, IsNil)

	fs, err = GetFSUsage()

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `no such file or directory`)
	c.Assert(fs, IsNil)

	util, err = GetIOUtil(time.Millisecond)

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Can't parse field 3 as unsigned integer in .*/test.file`)
	c.Assert(util, IsNil)

	procStatFile = "/_unknown_"
	mtabFile = "/etc/mtab"
	procDiskStatsFile = "/proc/diskstats"

	mtabFile = s.CreateTestFile(c, "/dev/abc1 / ext4 rw 0 0")

	fs, err = GetFSUsage()

	c.Assert(err, IsNil)
	c.Assert(fs, NotNil)

	c.Assert(CalculateIOUtil(nil, nil, time.Minute), IsNil)

	util = CalculateIOUtil(
		map[string]*IOStats{"1": {IOMs: 10}, "2": {IOMs: 1}, "3": {IOMs: 11}},
		map[string]*IOStats{"1": {IOMs: 1840}, "2": {IOMs: 10000000}},
		time.Minute,
	)

	c.Assert(util, NotNil)
	c.Assert(util["1"], Equals, 3.05)

	bs := bufio.NewScanner(strings.NewReader(" 11       0 vda 0 0 0 0 0 0 0 0 0 0 0\n 1       0 ram0 0 0 0 0 0 0 0 0 0 0 0\n  7       0 loop0 0 0 0 0 0 0 0 0 0 0 0\n"))
	st, err := parseIOStats(bs)

	c.Assert(err, IsNil)
	c.Assert(st, HasLen, 1)

	procDiskStatsFile = origProcDiskStatsFile
	procStatFile = origProcStatFile
	mtabFile = origMtabFile
}

func (s *SystemSuite) TestDiskStatsParsingErrors(c *C) {
	origProcDiskStatsFile := procDiskStatsFile

	for i := 0; i < 11; i++ {
		data := "   8       0 sda X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10"
		data = strings.Replace(data, "X"+strconv.Itoa(i), "A", -1)
		data = strings.Replace(data, "X", "", -1)

		procDiskStatsFile = s.CreateTestFile(c, data)

		stats, err := GetIOStats()

		c.Assert(err, NotNil)
		c.Assert(err, ErrorMatches, `Can't parse field [0-9]+ as unsigned integer in .*/test.file`)
		c.Assert(stats, IsNil)
	}

	procDiskStatsFile = origProcDiskStatsFile
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

	// This test can fail on CI because workers
	// doesn't have any active sessions
	if os.Getenv("CI") == "" {
		sess, err := Who()

		c.Assert(err, IsNil)
		c.Assert(sess, NotNil)
	}

	user, err = LookupUser("")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `User name/ID can't be blank`)
	c.Assert(user, IsNil)

	group, err := LookupGroup("root")

	c.Assert(err, IsNil)
	c.Assert(group, NotNil)

	group, err = LookupGroup("")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Group name/ID can't be blank`)
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

	groups = extractGroupsInfo("uid=66(someone) gid=66(someone) groups=66(someone)\n\n")

	c.Assert(groups[0].Name, Equals, "someone")
	c.Assert(groups[0].GID, Equals, 66)

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
	c.Assert(err, ErrorMatches, `Path is empty`)

	_, err = getSessionInfo("ABC")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `no such file or directory`)

	n, _ := fixCount(-100, nil)

	c.Assert(n, Equals, 0)

	ptsDir = "/not_exist"

	sess, err := Who()

	c.Assert(err, IsNil)
	c.Assert(sess, HasLen, 0)

}

func (s *SystemSuite) TestGetInfo(c *C) {
	origOsReleaseFile := osReleaseFile

	sysInfo, err := GetSystemInfo()
	c.Assert(err, IsNil)
	c.Assert(sysInfo, NotNil)
	c.Assert(sysInfo.Hostname, Not(Equals), "")
	c.Assert(sysInfo.OS, Not(Equals), "")
	c.Assert(sysInfo.Kernel, Not(Equals), "")
	c.Assert(sysInfo.Arch, Not(Equals), "")
	c.Assert(sysInfo.ArchName, Not(Equals), "")
	c.Assert(sysInfo.ArchBits, Not(Equals), 0)

	osReleaseFile = "/_UNKNOWN_"
	osInfo, err := GetOSInfo()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open /_UNKNOWN_: no such file or directory`)
	c.Assert(osInfo, IsNil)

	osReleaseFile = s.CreateTestFile(c, `NAME="Ubuntu"
VERSION="20.10 (Groovy Gorilla)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 20.10"
VERSION_ID="20.10"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
VERSION_CODENAME=groovy
ANSI_COLOR="0;34"
VARIANT="Server"
VARIANT_ID="server"
PLATFORM_ID="platform:el8"
LOGO=fedora-logo-icon
DOCUMENTATION_URL="https://docs.project.org"
CPE_NAME="cpe:/o:oracle:linux:8:6:server"

REDHAT_SUPPORT_PRODUCT="centos"
REDHAT_SUPPORT_PRODUCT_VERSION="7"
`)

	osInfo, err = GetOSInfo()

	c.Assert(osInfo.Name, Equals, "Ubuntu")
	c.Assert(osInfo.PrettyName, Equals, "Ubuntu 20.10")
	c.Assert(osInfo.Version, Equals, "20.10 (Groovy Gorilla)")
	c.Assert(osInfo.VersionID, Equals, "20.10")
	c.Assert(osInfo.VersionCodename, Equals, "groovy")
	c.Assert(osInfo.ID, Equals, "ubuntu")
	c.Assert(osInfo.IDLike, Equals, "debian")
	c.Assert(osInfo.PlatformID, Equals, "platform:el8")
	c.Assert(osInfo.Variant, Equals, "Server")
	c.Assert(osInfo.VariantID, Equals, "server")
	c.Assert(osInfo.CPEName, Equals, "cpe:/o:oracle:linux:8:6:server")
	c.Assert(osInfo.HomeURL, Equals, "https://www.ubuntu.com/")
	c.Assert(osInfo.BugReportURL, Equals, "https://bugs.launchpad.net/ubuntu/")
	c.Assert(osInfo.SupportURL, Equals, "https://help.ubuntu.com/")
	c.Assert(osInfo.DocumentationURL, Equals, "https://docs.project.org")
	c.Assert(osInfo.Logo, Equals, "fedora-logo-icon")
	c.Assert(osInfo.SupportProduct, Equals, "centos")
	c.Assert(osInfo.SupportProductVersion, Equals, "7")
	c.Assert(osInfo.ANSIColor, Equals, "0;34")

	c.Assert(osInfo.ColoredPrettyName(), Equals, "\x1b[0;34mUbuntu 20.10\x1b[0m")
	c.Assert(osInfo.ColoredName(), Equals, "\x1b[0;34mUbuntu\x1b[0m")
	osInfo.ANSIColor = "ABCD"
	c.Assert(osInfo.ColoredPrettyName(), Equals, "Ubuntu 20.10")
	c.Assert(osInfo.ColoredName(), Equals, "Ubuntu")

	c.Assert(getArchName("i386"), Equals, "386")
	c.Assert(getArchName("i586"), Equals, "586")
	c.Assert(getArchName("i686"), Equals, "686")
	c.Assert(getArchName("x86_64"), Equals, "amd64")
	c.Assert(getArchName("aarch64"), Equals, "aarch64")

	osReleaseFile = origOsReleaseFile
}

func (s *SystemSuite) TestGetCPUArchBits(c *C) {
	origCpuInfoFile := cpuInfoFile

	cpuInfoFile = "/_UNKNOWN_"

	c.Assert(getCPUArchBits(), Equals, 0)

	cpuInfoFile = s.CreateTestFile(c, "processor       : 1\nflags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc aperfmperf eagerfpu pni pclmulqdq dtes64")

	c.Assert(getCPUArchBits(), Equals, 64)

	cpuInfoFile = s.CreateTestFile(c, "processor       : 1\nflags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc aperfmperf eagerfpu pni pclmulqdq dtes64")

	c.Assert(getCPUArchBits(), Equals, 32)

	cpuInfoFile = origCpuInfoFile
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SystemSuite) CreateTestFile(c *C, data string) string {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/test.file"

	if os.WriteFile(tmpFile, []byte(data), 0644) != nil {
		c.Fatal("Can't create temporary file")
	}

	return tmpFile
}
