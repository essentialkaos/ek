// Package network provides KNF validators for checking items related to network
package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/essentialkaos/ek/v13/knf"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// IP returns error if configuration property isn't a valid IP address
	IP = validateIP

	// Port returns error if configuration property isn't a valid port number
	Port = validatePort

	// MAC returns error if configuration property isn't a valid MAC address
	MAC = validateMAC

	// CIDR returns error if configuration property isn't a valid CIDR address
	CIDR = validateCIDR

	// URL returns error if configuration property isn't a valid URL
	URL = validateURL

	// Mail returns error if configuration property isn't a valid email address
	Mail = validateMail

	// HasIP returns error if system doesn't have interface with IP from configuration
	// property
	HasIP = validateHasIP
)

// ////////////////////////////////////////////////////////////////////////////////// //

// validateIP returns error if configuration property isn't a valid IP address
func validateIP(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	ip := net.ParseIP(v)

	if ip == nil {
		return fmt.Errorf("%q is not a valid IP address", v)
	}

	return nil
}

// validatePort returns error if configuration property isn't a valid port number
func validatePort(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	portInt := config.GetI(prop)

	if portInt == 0 || portInt > 65535 {
		return fmt.Errorf("%q is not a valid port number", v)
	}

	return nil
}

// validateMAC returns error if configuration property isn't a valid MAC address
func validateMAC(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	_, err := net.ParseMAC(v)

	if err != nil {
		return fmt.Errorf("%q is not a valid MAC address: %v", v, err)
	}

	return nil
}

// validateCIDR returns error if configuration property isn't a valid CIDR address
func validateCIDR(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	_, _, err := net.ParseCIDR(v)

	if err != nil {
		return fmt.Errorf("%q is not a valid CIDR address: %v", v, err)
	}

	return nil
}

// validateURL returns error if configuration property isn't a valid URL
func validateURL(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	_, err := url.ParseRequestURI(v)

	if err != nil {
		return fmt.Errorf("%q is not a valid URL address: %v", v, err)
	}

	return nil
}

// validateMail returns error if configuration property isn't a valid email address
func validateMail(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
		return nil
	}

	if !strings.ContainsRune(v, '@') || !strings.ContainsRune(v, '.') {
		return fmt.Errorf("%q is not a valid email address", v)
	}

	return nil
}

// validateHasIP returns error if system doesn't have interface with IP from configuration
// property
func validateHasIP(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
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
			if v == a.(*net.IPNet).IP.String() {
				return nil
			}
		}
	}

	return fmt.Errorf("The system does not have an interface with the address %q", v)
}
