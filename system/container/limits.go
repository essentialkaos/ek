package container

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// cgroupDir is process's cgroup v2 directory
var cgroupDir = "/sys/fs/cgroup/"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUQuota returns the CPU quota in cores. If cpu.max is "max",
// it returns (0, true, nil).
func GetCPUQuota() (float64, bool, error) {
	value, err := readCGroupsFile("cpu.max")

	if err != nil {
		return 0, false, err
	}

	fields := strings.Fields(value)

	if len(fields) != 2 {
		return 0, false, fmt.Errorf("invalid cpu.max: %q", value)
	}

	period, err := strconv.ParseUint(fields[1], 10, 64)

	if err != nil || period == 0 {
		return 0, false, fmt.Errorf("invalid cpu.max period: %q", fields[1])
	}

	if fields[0] == "max" {
		return 0, true, nil
	}

	quota, err := strconv.ParseUint(fields[0], 10, 64)

	if err != nil || quota == 0 {
		return 0, false, fmt.Errorf("invalid cpu.max quota: %q", fields[0])
	}

	return float64(quota) / float64(period), false, nil
}

// GetMemLimit returns the memory limit in bytes. If memory.max is "max",
// it returns (0, true, nil).
func GetMemLimit() (uint64, bool, error) {
	value, err := readCGroupsFile("memory.max")

	if err != nil {
		return 0, false, err
	}

	if value == "max" {
		return 0, true, nil
	}

	limit, err := strconv.ParseUint(value, 10, 64)

	if err != nil {
		return 0, false, fmt.Errorf("invalid memory.max: %w", err)
	}

	return limit, false, nil
}

// GetCPUSet returns a slice indexed by logical CPU ID. True means that CPU is
// in the effective cpuset. The slice ends at the highest CPU ID in the set.
func GetCPUSet() ([]bool, error) {
	value, err := readCGroupsFile("cpuset.cpus.effective")

	if err != nil {
		return nil, err
	}

	cpus := []bool{}

	if value == "" {
		return cpus, nil
	}

	for _, entry := range strings.Split(value, ",") {
		first, last, isRange := strings.Cut(entry, "-")

		if !isRange {
			last = first
		}

		start, startErr := strconv.Atoi(first)
		end, endErr := strconv.Atoi(last)

		if startErr != nil || endErr != nil || start < 0 || end < start {
			return nil, fmt.Errorf("invalid cpuset.cpus.effective entry: %q", entry)
		}

		if end == math.MaxInt {
			return nil, fmt.Errorf("CPU ID too large for slice indexing: %d", end)
		}

		size := end + 1

		if size > len(cpus) {
			cpus = append(cpus, make([]bool, size-len(cpus))...)
		}

		for cpu := start; cpu < size; cpu++ {
			cpus[cpu] = true
		}
	}

	return cpus, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// readCGroupsFile reads cgroups file data
func readCGroupsFile(name string) (string, error) {
	data, err := os.ReadFile(cgroupDir + "/" + name)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
