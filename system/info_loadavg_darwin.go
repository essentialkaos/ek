package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os/exec"
	"strings"
	"unsafe"

	"golang.org/x/sys/unix"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetLA returns loadavg
//
// Implementation borrowed from https://github.com/prometheus/node_exporter/blob/master/collector/loadavg_bsd.go
func GetLA() (*LoadAvg, error) {
	type loadavg struct {
		load  [3]uint32
		scale int
	}

	b, err := unix.SysctlRaw("vm.loadavg")

	if err != nil {
		return nil, err
	}

	rproc, tproc, err := getProcStats()

	if err != nil {
		return nil, err
	}

	load := *(*loadavg)(unsafe.Pointer((&b[0])))
	scale := float64(load.scale)

	return &LoadAvg{
		Min1:  float64(load.load[0]) / scale,
		Min5:  float64(load.load[1]) / scale,
		Min15: float64(load.load[2]) / scale,
		RProc: rproc,
		TProc: tproc,
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getProcStats() (int, int, error) {
	cmd := exec.Command("ps", "axo", "state")
	output, err := cmd.Output()

	if err != nil {
		return 0, 0, errors.New("Can't run ps command for collecting information about processes")
	}

	outputStr := string(output)
	rproc := strings.Count(outputStr, "R")
	tproc := strings.Count(outputStr, "\n")

	return rproc, tproc, nil
}
