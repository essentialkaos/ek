package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v14/strutil"
	"github.com/essentialkaos/ek/v14/system/sysctl"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetMemUsage returns current physical and swap memory usage
func GetMemUsage() (*MemUsage, error) {
	params, err := sysctl.All()

	if err != nil {
		return nil, fmt.Errorf("can't get kernel parameters: %w", err)
	}

	pagesize, err := params.Get("hw.pagesize").Int()

	if err != nil {
		return nil, fmt.Errorf("can't read page size from sysctl: %w", err)
	}

	totalMem, err := params.Get("hw.memsize_usable").Int()

	if err != nil {
		return nil, fmt.Errorf("can't read total memory from sysctl: %w", err)
	}

	info, err := calculateMemUsage(uint64(pagesize), uint64(totalMem))

	if err != nil {
		return nil, err
	}

	appendSwapUsage(info, params)

	return info, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// calculateMemUsage calculates memory usage using data from vm_stats
func calculateMemUsage(pageSize, totalMem uint64) (*MemUsage, error) {
	output, err := exec.Command("vm_stat").Output()

	if err != nil {
		return nil, fmt.Errorf("can't get output from vm_stats")
	}

	buf := bytes.NewBuffer(output)

	var free, active, inactive, speculative uint64

	for range 5 {
		line, err := buf.ReadString('\n')

		if err != nil {
			break
		}

		_, v, _ := strings.Cut(line, ": ")
		vu, _ := strconv.ParseUint(strings.Trim(v, " .\n\r"), 10, 64)
		vu *= pageSize

		switch strutil.Substr(line, 0, 10) {
		case "Pages free":
			free = vu
		case "Pages acti":
			active = vu
		case "Pages inac":
			inactive = vu
		case "Pages spec":
			speculative = vu
		}
	}

	return &MemUsage{
		MemTotal: totalMem,
		MemFree:  free + speculative + inactive,
		MemUsed:  totalMem - (free + speculative + inactive),
		Active:   active,
		Inactive: inactive,
	}, nil
}

// appendSwapUsage appends swap usage info
func appendSwapUsage(info *MemUsage, params sysctl.Params) {
	param := params.Get("vm.swapusage")

	if param.IsEmpty() {
		return
	}

	swap := strutil.SqueezeRepeats(param.Value, " ")
	swap = strings.ReplaceAll(swap, " = ", "=")
	swap = strings.ReplaceAll(swap, "M", "")

	for i := range 3 {
		l := strutil.ReadField(swap, i, false, ' ')
		n, v, _ := strings.Cut(l, "=")
		fv, _ := strconv.ParseFloat(v, 10)

		switch n {
		case "total":
			info.SwapTotal = uint64(fv * 1024 * 1024)
		case "used":
			info.SwapUsed = uint64(fv * 1024 * 1024)
		case "free":
			info.SwapFree = uint64(fv * 1024 * 1024)
		}
	}
}
