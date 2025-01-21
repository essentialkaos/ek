// Package network provides methods for collecting information about machine network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net"
	"strings"

	"github.com/essentialkaos/ek/v13/req"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

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
