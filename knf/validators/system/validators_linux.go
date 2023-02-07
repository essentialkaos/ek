// Package system provides KNF validators for checking system items (user, groups,
// network interfaces)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"github.com/essentialkaos/ek/v12/knf"
	"github.com/essentialkaos/ek/v12/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Interface returns error if config property contains name of network interface
	// which not present on the system
	Interface = validateInterface
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validateInterface(config *knf.Config, prop string, value any) error {
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
		return fmt.Errorf("Interface %s is not present on the system", interfaceName)
	}

	return nil
}
