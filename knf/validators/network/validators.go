// Package network provides KNF validators for checking items related to network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"net"
	"net/url"

	"pkg.re/essentialkaos/ek.v12/knf"
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
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validateIP(config *knf.Config, prop string, value interface{}) error {
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

func validatePort(config *knf.Config, prop string, value interface{}) error {
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

func validateMAC(config *knf.Config, prop string, value interface{}) error {
	macStr := config.GetS(prop)

	if macStr == "" {
		return nil
	}

	_, err := net.ParseMAC(macStr)

	if err != nil {
		return fmt.Errorf("%s is not a valid MAC address: %v", macStr, err)
	}

	return err
}

func validateCIDR(config *knf.Config, prop string, value interface{}) error {
	cidrStr := config.GetS(prop)

	if cidrStr == "" {
		return nil
	}

	_, _, err := net.ParseCIDR(cidrStr)

	if err != nil {
		return fmt.Errorf("%s is not a valid CIDR address: %v", cidrStr, err)
	}

	return err
}

func validateURL(config *knf.Config, prop string, value interface{}) error {
	urlStr := config.GetS(prop)

	if urlStr == "" {
		return nil
	}

	_, err := url.ParseRequestURI(urlStr)

	if err != nil {
		return fmt.Errorf("%s is not a valid URL address: %v", urlStr, err)
	}

	return err
}
