// Package network provides methods for collecting information about machine network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"

	"github.com/essentialkaos/ek/v12/netutil"
	"github.com/essentialkaos/ek/v12/req"
	"github.com/essentialkaos/ek/v12/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects network info
func Collect(ipResolverURL ...string) *support.NetworkInfo {
	info := &support.NetworkInfo{
		IPv4: netutil.GetAllIP(),
		IPv6: netutil.GetAllIP6(),
	}

	info.Hostname, _ = os.Hostname()

	if len(ipResolverURL) != 0 {
		info.PublicIP = resolvePublicIP(ipResolverURL[0])
	}

	return info
}

// ////////////////////////////////////////////////////////////////////////////////// //

// resolvePublicIP resolves public IP using IP resolver
func resolvePublicIP(resolver string) string {
	resp, err := req.Request{
		URL:         resolver,
		AutoDiscard: true,
	}.Get()

	if err != nil {
		return ""
	}

	return resp.String()
}
