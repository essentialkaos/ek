//go:build gofuzz
// +build gofuzz

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"bytes"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func FuzzCPUStats(data []byte) int {
	r := bytes.NewReader(data)
	s := bufio.NewScanner(r)

	_, err := parseCPUStats(s)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzCPUInfo(data []byte) int {
	r := bytes.NewReader(data)
	s := bufio.NewScanner(r)

	_, err := parseCPUInfo(s)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzMemUsage(data []byte) int {
	r := bytes.NewReader(data)
	s := bufio.NewScanner(r)

	_, err := parseMemUsage(s)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzIOStats(data []byte) int {
	r := bytes.NewReader(data)
	s := bufio.NewScanner(r)

	_, err := parseIOStats(s)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzFSInfo(data []byte) int {
	r := bytes.NewReader(data)
	s := bufio.NewScanner(r)

	_, err := parseFSInfo(s, false)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzInterfacesStats(data []byte) int {
	r := bytes.NewReader(data)
	s := bufio.NewScanner(r)

	_, err := parseInterfacesStats(s)

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
