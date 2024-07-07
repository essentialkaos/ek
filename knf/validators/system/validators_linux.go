// Package system provides KNF validators for checking system items (user, groups,
// network interfaces)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"github.com/essentialkaos/ek.v13/knf"
	"github.com/essentialkaos/ek.v13/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Interface returns error if config property contains name of network interface
	// which not present on the system
	Interface = validateInterface
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validateInterface(config knf.IConfig, prop string, value any) error {
	interfaceName := config.GetS(prop)

	if interfaceName == "" {
		return nil
	}

	stats, err := system.GetInterfacesStats()

	if err != nil {
		return fmt.Errorf("Can't get interfaces info: %v", err)
	}

	_, isPresent := stats[interfaceName]

	if !isPresent {
		return fmt.Errorf("Interface %q is not present on the system", interfaceName)
	}

	return nil
}
