// Package network provides KNF validators for checking items related to network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"net"
	"net/url"

	"github.com/essentialkaos/ek/v12/knf"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// IP returns error if config property isn't a valid IP address
	IP = validateIP

	// Port returns error if config property isn't a valid port number
	Port = validatePort

	// MAC returns error if config property isn't a valid MAC address
	MAC = validateMAC

	// CIDR returns error if config property isn't a valid CIDR address
	CIDR = validateCIDR

	// URL returns error if config property isn't a valid URL
	URL = validateURL

	// HasIP returns error if system doesn't have interface with IP from config property
	HasIP = validateHasIP
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validateIP(config *knf.Config, prop string, value any) error {
	ipStr := config.GetS(prop)

	if ipStr == "" {
		return nil
	}

	ip := net.ParseIP(ipStr)

	if ip == nil {
		return fmt.Errorf("%s is not a valid IP address", ipStr)
	}

	return nil
}

func validatePort(config *knf.Config, prop string, value any) error {
	portStr := config.GetS(prop)

	if portStr == "" {
		return nil
	}

	portInt := config.GetI(prop)

	if portInt == 0 || portInt > 65535 {
		return fmt.Errorf("%s is not a valid port number", portStr)
	}

	return nil
}

func validateMAC(config *knf.Config, prop string, value any) error {
	macStr := config.GetS(prop)

	if macStr == "" {
		return nil
	}

	_, err := net.ParseMAC(macStr)

	if err != nil {
		return fmt.Errorf("%s is not a valid MAC address: %v", macStr, err)
	}

	return nil
}

func validateCIDR(config *knf.Config, prop string, value any) error {
	cidrStr := config.GetS(prop)

	if cidrStr == "" {
		return nil
	}

	_, _, err := net.ParseCIDR(cidrStr)

	if err != nil {
		return fmt.Errorf("%s is not a valid CIDR address: %v", cidrStr, err)
	}

	return nil
}

func validateURL(config *knf.Config, prop string, value any) error {
	urlStr := config.GetS(prop)

	if urlStr == "" {
		return nil
	}

	_, err := url.ParseRequestURI(urlStr)

	if err != nil {
		return fmt.Errorf("%s is not a valid URL address: %v", urlStr, err)
	}

	return nil
}

func validateHasIP(config *knf.Config, prop string, value any) error {
	ipStr := config.GetS(prop)

	if ipStr == "" {
		return nil
	}

	interfaces, err := net.Interfaces()

	if err != nil {
		return fmt.Errorf("Can't get interfaces info for check: %v", err)
	}

	for _, i := range interfaces {
		addr, err := i.Addrs()

		if err != nil {
			continue
		}

		for _, a := range addr {
			if ipStr == a.(*net.IPNet).IP.String() {
				return nil
			}
		}
	}

	return fmt.Errorf("The system does not have an interface with the address %s", ipStr)
}
