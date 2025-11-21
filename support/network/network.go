//go:build !windows

// Package network provides methods for collecting information about machine network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"slices"

	"github.com/essentialkaos/ek/v13/netutil"
	"github.com/essentialkaos/ek/v13/sortutil"

	"github.com/essentialkaos/ek/v13/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects network info
func Collect(ipResolverURL ...string) *support.NetworkInfo {
	info := &support.NetworkInfo{
		IPv4: cleanIPList(netutil.GetAllIP()),
		IPv6: cleanIPList(netutil.GetAllIP6()),
	}

	sortutil.StringsNatural(info.IPv4)
	sortutil.StringsNatural(info.IPv6)

	info.IPv4 = slices.Compact(info.IPv4)
	info.IPv6 = slices.Compact(info.IPv6)

	info.Hostname, _ = os.Hostname()

	if len(ipResolverURL) != 0 {
		info.PublicIP = resolvePublicIP(ipResolverURL[0])
	}

	return info
}

// ////////////////////////////////////////////////////////////////////////////////// //

// cleanIPList returns IP slice without local IP's
func cleanIPList(ips []string) []string {
	var result []string

	for _, ip := range ips {
		switch ip {
		case "127.0.0.1", "::1":
			continue
		default:
			result = append(result, ip)
		}
	}

	return result
}
