// Package network provides methods for collecting information about machine network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net"
	"strings"

	"github.com/essentialkaos/ek/v14/req"
	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// resolvePublicIP resolves public IP using IP resolver
//
// Possible resolvers:
// - https://www.cloudflare.com/cdn-cgi/trace
// - https://1.1.1.1/cdn-cgi/trace
// - http://eth0.me
// - http://checkip.amazonaws.com
// - http://api.ipify.org
// - https://ipv4-internet.yandex.net/api/v0/ip
func resolvePublicIP(resolvers ...string) string {
	for _, resolverURL := range resolvers {
		resp, err := req.Request{
			URL:         resolverURL,
			AutoDiscard: true,
		}.Get()

		if err != nil || resp.StatusCode != 200 {
			continue
		}

		var ip string

		if strings.HasSuffix(resolverURL, "/cdn-cgi/trace") {
			ip = extractIPFromCloudflareTrace(resp.String())
		} else {
			ip = strutil.Exclude(resp.String(), `"`)
		}

		if isValidIP(ip) {
			return ip
		}
	}

	return ""
}

// extractIPFromCloudflareTrace extracts public IP from Cloudflare trace
// response
func extractIPFromCloudflareTrace(data string) string {
	for i := range 16 {
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
