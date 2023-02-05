package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetInterfacesStats returns info about network interfaces
func GetInterfacesStats() (map[string]*InterfaceStats, error) {
	panic("UNSUPPORTED")
	return map[string]*InterfaceStats{"eth0": {}}, nil
}

// ❗ GetNetworkSpeed returns input/output speed in bytes per second
func GetNetworkSpeed(duration time.Duration) (uint64, uint64, error) {
	panic("UNSUPPORTED")
	return 0, 0, nil
}

// ❗ CalculateNetworkSpeed calculates network input/output speed in bytes per second for
// all network interfaces
func CalculateNetworkSpeed(ii1, ii2 map[string]*InterfaceStats, duration time.Duration) (uint64, uint64) {
	panic("UNSUPPORTED")
	return 0, 0
}
