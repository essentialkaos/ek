// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// LoadAvg contains information about average system load
type LoadAvg struct {
	Min1  float64 `json:"min1"`  // LA in last 1 minute
	Min5  float64 `json:"min5"`  // LA in last 5 minutes
	Min15 float64 `json:"min15"` // LA in last 15 minutes
	RProc int     `json:"rproc"` // Number of currently runnable kernel scheduling entities
	TProc int     `json:"tproc"` // Number of kernel scheduling entities that currently exist on the system
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with LA info in procfs
var procLoadAvgFile = "/proc/loadavg"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetLA return loadavg
func GetLA() (*LoadAvg, error) {
	content, err := readFileContent(procLoadAvgFile)

	if err != nil {
		return nil, err
	}

	contentSlice := strings.Split(content[0], " ")

	if len(contentSlice) != 5 {
		return nil, errors.New("Can't parse file " + procLoadAvgFile)
	}

	procSlice := strings.Split(contentSlice[3], "/")

	la := &LoadAvg{}

	la.Min1, _ = strconv.ParseFloat(contentSlice[0], 64)
	la.Min5, _ = strconv.ParseFloat(contentSlice[1], 64)
	la.Min15, _ = strconv.ParseFloat(contentSlice[2], 64)
	la.RProc, _ = strconv.Atoi(procSlice[0])
	la.TProc, _ = strconv.Atoi(procSlice[1])

	return la, nil
}
