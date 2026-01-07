// Package network provides methods for collecting information about machine network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v13/support"

	"golang.org/x/sys/windows"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects network info
func Collect(ipResolverURL ...string) *support.NetworkInfo {
	info := &support.NetworkInfo{}

	info.Hostname, _ = windows.ComputerName()

	if len(ipResolverURL) != 0 {
		info.PublicIP = resolvePublicIP(ipResolverURL[0])
	}

	if info.Hostname == "" && info.PublicIP == "" {
		return nil
	}

	return info
}
