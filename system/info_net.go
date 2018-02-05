// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v9/errutil"
	"pkg.re/essentialkaos/ek.v9/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with net info in procfs
var procNetFile = "/proc/net/dev"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetInterfacesInfo return info about network interfaces
func GetInterfacesInfo() (map[string]*InterfaceInfo, error) {
	fd, err := os.OpenFile(procNetFile, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	info := make(map[string]*InterfaceInfo)
	errs := errutil.NewErrors()

	for s.Scan() {
		text := s.Text()

		if !strings.Contains(text, ":") {
			continue
		}

		ii := &InterfaceInfo{}

		name := strings.TrimRight(strutil.ReadField(text, 0, true), ":")

		ii.ReceivedBytes = parseUint(strutil.ReadField(text, 1, true), errs)
		ii.ReceivedPackets = parseUint(strutil.ReadField(text, 2, true), errs)
		ii.TransmittedBytes = parseUint(strutil.ReadField(text, 9, true), errs)
		ii.TransmittedPackets = parseUint(strutil.ReadField(text, 10, true), errs)

		if errs.HasErrors() {
			return nil, errs.Last()
		}

		info[name] = ii
	}

	if len(info) == 0 {
		return nil, errors.New("Can't parse file " + procNetFile)
	}

	return info, nil
}

// GetNetworkSpeed return network input/output speed in bytes per second for
// all network interfaces
func GetNetworkSpeed(duration time.Duration) (uint64, uint64, error) {
	ii1, err := GetInterfacesInfo()

	if err != nil {
		return 0, 0, err
	}

	time.Sleep(duration)

	ii2, err := GetInterfacesInfo()

	if err != nil {
		return 0, 0, err
	}

	in, out := CalculateNetworkSpeed(ii1, ii2, duration)

	return in, out, nil
}

// CalculateNetworkSpeed calculate network input/output speed in bytes per second for
// all network interfaces
func CalculateNetworkSpeed(ii1, ii2 map[string]*InterfaceInfo, duration time.Duration) (uint64, uint64) {
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

// getActiveInterfacesBytes calculate received and transmitted bytes on all interfaces
func getActiveInterfacesBytes(is map[string]*InterfaceInfo) (uint64, uint64) {
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
