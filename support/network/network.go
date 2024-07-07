//go:build !windows
// +build !windows

// Package network provides methods for collecting information about machine network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net"
	"os"
	"strings"

	"github.com/essentialkaos/ek.v13/netutil"
	"github.com/essentialkaos/ek.v13/req"
	"github.com/essentialkaos/ek.v13/sliceutil"
	"github.com/essentialkaos/ek.v13/sortutil"
	"github.com/essentialkaos/ek.v13/strutil"

	"github.com/essentialkaos/ek.v13/support"
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

	info.IPv4 = sliceutil.Deduplicate(info.IPv4)
	info.IPv6 = sliceutil.Deduplicate(info.IPv6)

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

// resolvePublicIP resolves public IP using IP resolver
func resolvePublicIP(resolverURL string) string {
	resp, err := req.Request{
		URL:         resolverURL,
		AutoDiscard: true,
	}.Get()

	if err != nil {
		return ""
	}

	var ip string

	if strings.HasSuffix(resolverURL, "/cdn-cgi/trace") {
		ip = extractIPFromCloudflareTrace(resp.String())
	} else {
		ip = resp.String()
	}

	if isValidIP(ip) {
		return ip
	}

	return ""
}

// extractIPFromCloudflareTrace extracts public IP from Cloudflare trace
// response
func extractIPFromCloudflareTrace(data string) string {
	for i := 0; i < 16; i++ {
		f := strutil.ReadField(data, i, false, '\n')

		if f == "" {
			break
		}

		n := strutil.ReadField(f, 0, false, '=')

		if n != "ip" {
			continue
		}

		return strutil.ReadField(f, 1, false, '=')
	}

	return ""
}

// isValidIP returns if given IP is valid
func isValidIP(ip string) bool {
	return ip != "" && net.ParseIP(ip) != nil
}
