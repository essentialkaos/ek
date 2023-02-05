package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with net info in procfs
var procNetFile = "/proc/net/dev"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetInterfacesStats returns info about network interfaces
func GetInterfacesStats() (map[string]*InterfaceStats, error) {
	s, closer, err := getFileScanner(procNetFile)

	if err != nil {
		return nil, err
	}

	defer closer()

	return parseInterfacesStats(s)
}

// GetNetworkSpeed returns network input/output speed in bytes per second for
// all network interfaces
func GetNetworkSpeed(duration time.Duration) (uint64, uint64, error) {
	ii1, err := GetInterfacesStats()

	if err != nil {
		return 0, 0, err
	}

	time.Sleep(duration)

	ii2, err := GetInterfacesStats()

	if err != nil {
		return 0, 0, err
	}

	in, out := CalculateNetworkSpeed(ii1, ii2, duration)

	return in, out, nil
}

// CalculateNetworkSpeed calculates network input/output speed in bytes per second for
// all network interfaces
func CalculateNetworkSpeed(ii1, ii2 map[string]*InterfaceStats, duration time.Duration) (uint64, uint64) {
	if ii1 == nil || ii2 == nil {
		return 0, 0
	}

	rb1, tb1 := getActiveInterfacesBytes(ii1)
	rb2, tb2 := getActiveInterfacesBytes(ii2)

	if rb1+tb1 == 0 || rb2+tb2 == 0 {
		return 0, 0
	}

	durationSec := uint64(duration / time.Second)

	return (rb2 - rb1) / durationSec, (tb2 - tb1) / durationSec
}

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[LOC,ABC]

// parseInterfacesStats parses interfaces stats data
func parseInterfacesStats(s *bufio.Scanner) (map[string]*InterfaceStats, error) {
	var err error

	stats := make(map[string]*InterfaceStats)

	for s.Scan() {
		text := s.Text()

		if !strings.Contains(text, ":") {
			continue
		}

		ii := &InterfaceStats{}

		name := strings.TrimRight(strutil.ReadField(text, 0, true), ":")

		ii.ReceivedBytes, err = strconv.ParseUint(strutil.ReadField(text, 1, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 1 as unsigned integer in " + procNetFile)
		}

		ii.ReceivedPackets, err = strconv.ParseUint(strutil.ReadField(text, 2, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 2 as unsigned integer in " + procNetFile)
		}

		ii.TransmittedBytes, err = strconv.ParseUint(strutil.ReadField(text, 9, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 9 as unsigned integer in " + procNetFile)
		}

		ii.TransmittedPackets, err = strconv.ParseUint(strutil.ReadField(text, 10, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 10 as unsigned integer in " + procNetFile)
		}

		stats[name] = ii
	}

	if len(stats) == 0 {
		return nil, errors.New("Can't parse file " + procNetFile)
	}

	return stats, nil
}

// codebeat:enable[LOC,ABC]

// getActiveInterfacesBytes calculate received and transmitted bytes on all interfaces
func getActiveInterfacesBytes(is map[string]*InterfaceStats) (uint64, uint64) {
	var (
		received    uint64
		transmitted uint64
	)

	for name, info := range is {
		if strings.HasPrefix(name, "lo") || strings.HasPrefix(name, "bond") {
			continue
		}

		if info.ReceivedBytes == 0 && info.TransmittedBytes == 0 {
			continue
		}

		received += info.ReceivedBytes
		transmitted += info.TransmittedBytes
	}

	return received, transmitted
}
