//go:build linux || darwin

// Package netutil provides methods for working with network
package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetIP returns main IPv4 address
func GetIP() string {
	return getMainIP(false)
}

// GetIP6 returns main IPv6 address
func GetIP6() string {
	return getMainIP(true)
}

// GetAllIP returns all IPv4 addresses
func GetAllIP() []string {
	return getAllIP(false)
}

// GetAllIP6 returns all IPv6 addresses
func GetAllIP6() []string {
	return getAllIP(true)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getAllIP returns all IP addresses (IPv4 or IPv6)
func getAllIP(v6 bool) []string {
	interfaces, err := net.Interfaces()

	if err != nil {
		return nil
	}

	var result []string

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()

		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			isV6 := ipnet.IP.To4() == nil

			if ok && isV6 == v6 {
				result = append(result, ipnet.IP.String())
			}
		}
	}

	return result
}

// getMainIP returns main IP address (IPv4 or IPv6)
func getMainIP(v6 bool) string {
	interfaces, err := net.Interfaces()

	if err != nil {
		return ""
	}

	defaultInterface := getDefaultRouteInterface()

	for i := len(interfaces) - 1; i >= 0; i-- {
		if defaultInterface != "" && interfaces[i].Name != defaultInterface {
			continue
		}

		// Ignore TUN/TAP interfaces
		switch {
		case strings.Contains(interfaces[i].Name, "tun"),
			strings.Contains(interfaces[i].Name, "tap"):
			continue
		}

		addrs, err := interfaces[i].Addrs()

		if err != nil || len(addrs) == 0 {
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)

			if !ok {
				continue
			}

			if ipnet.IP.IsLoopback() {
				continue
			}

			isV6 := ipnet.IP.To4() == nil

			if ok && isV6 == v6 {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
