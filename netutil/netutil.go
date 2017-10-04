// Package netutil with network utils
package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetIP return main IPv4 address
func GetIP() string {
	return getMainIP(false)
}

// GetIP6 return main IPv6 address
func GetIP6() string {
	return getMainIP(true)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getMainIP(v6 bool) string {
	interfaces, err := net.Interfaces()

	if err != nil {
		return ""
	}

	for i := len(interfaces) - 1; i >= 0; i-- {
		// Ignore TUN/TAP interfaces
		if strings.HasPrefix(interfaces[i].Name, "t") {
			continue
		}

		addr, err := interfaces[i].Addrs()

		if err != nil || len(addr) == 0 {
			continue
		}

		for _, a := range addr {
			ipnet, ok := a.(*net.IPNet)

			if ok && strings.Contains(ipnet.IP.String(), "::") == v6 {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
