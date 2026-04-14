//go:build !windows

// Package network provides methods for collecting information about machine network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"slices"

	"github.com/essentialkaos/ek/v14/netutil"
	"github.com/essentialkaos/ek/v14/sortutil"

	"github.com/essentialkaos/ek/v14/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect gathers hostname, local IPv4/IPv6 addresses, and optionally the
// public IP using the first provided resolver URL
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
		for _, resolver := range ipResolverURL {
			info.PublicIP = resolvePublicIP(resolver)

			if info.PublicIP != "" {
				break
			}
		}
	}

	return info
}

// ////////////////////////////////////////////////////////////////////////////////// //

// cleanIPList returns IP slice without local IP's
func cleanIPList(ips []string) []string {
	return slices.DeleteFunc(ips, func(ip string) bool {
		return ip == "127.0.0.1" || ip == "::1"
	})
}
