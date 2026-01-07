package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"golang.org/x/sys/windows/registry"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// registryKey is key of registry record with TZ data
var registryKey = `SYSTEM\CurrentControlSet\Control\TimeZoneInformation`

// registryKeyName is name of registry key record
var registryKeyName = "TimeZoneKeyName"

// ////////////////////////////////////////////////////////////////////////////////// //

// LocalTimezone returns name of local timezone
func LocalTimezone() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, registryKey, registry.QUERY_VALUE)

	if err != nil {
		return "Local"
	}

	tzName, _, err := key.GetStringValue(registryKeyName)

	if err != nil {
		return "Local"
	}

	return tzName
}
