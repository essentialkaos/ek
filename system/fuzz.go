// +build gofuzz

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
)

func FuzzCPUStats(data []byte) int {
	r := bytes.NewReader(data)
	_, err := parseCPUStats(r)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzCPUInfo(data []byte) int {
	r := bytes.NewReader(data)
	_, err := parseCPUInfo(r)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzMemUsage(data []byte) int {
	r := bytes.NewReader(data)
	_, err := parseMemUsage(r)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzIOStats(data []byte) int {
	r := bytes.NewReader(data)
	_, err := parseIOStats(r)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzFSInfo(data []byte) int {
	r := bytes.NewReader(data)
	_, err := parseFSInfo(r, false)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzInterfacesStats(data []byte) int {
	r := bytes.NewReader(data)
	_, err := parseInterfacesStats(r)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzLAInfo(data []byte) int {
	_, err := parseLAInfo(string(data))

	if err != nil {
		return 1
	}

	return 0
}

func FuzzUptime(data []byte) int {
	_, err := parseUptime(string(data))

	if err != nil {
		return 1
	}

	return 0
}
