//go:build !windows
// +build !windows

// Package services provides methods for collecting information about system services
package services

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v13/initsystem"
	"github.com/essentialkaos/ek/v13/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects info about services
func Collect(services ...string) []support.Service {
	var result []support.Service

	for _, s := range services {
		service := support.Service{Name: s, Status: support.STATUS_UNKNOWN}

		if initsystem.IsPresent(s) {
			service.IsPresent = true
			service.IsEnabled, _ = initsystem.IsEnabled(s)

			isWorks, err := initsystem.IsWorks(s)

			switch {
			case err == nil && isWorks:
				service.Status = support.STATUS_WORKS
			case err == nil && !isWorks:
				service.Status = support.STATUS_STOPPED
			}
		}

		result = append(result, service)
	}

	return result
}
