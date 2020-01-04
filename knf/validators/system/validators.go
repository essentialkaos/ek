// Package system provides KNF validators for checking system items (user, groups,
// network interfaces)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"pkg.re/essentialkaos/ek.v11/knf"
	"pkg.re/essentialkaos/ek.v11/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// User returns error if config property contains name of user or UID which not
	// present on the system
	User = validateUser

	// Group returns error if config property contains name of group or GID which not
	// present on the system
	Group = validateGroup

	// Interface returns error if config property contains name of network interface
	// which not present on the system
	Interface = validateInterface
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validateUser(config *knf.Config, prop string, value interface{}) error {
	userNameOrID := config.GetS(prop)

	if userNameOrID == "" {
		return nil
	}

	if !system.IsUserExist(userNameOrID) {
		return fmt.Errorf("User %s is not present on the system", userNameOrID)
	}

	return nil
}

func validateGroup(config *knf.Config, prop string, value interface{}) error {
	groupNameOrID := config.GetS(prop)

	if groupNameOrID == "" {
		return nil
	}

	if !system.IsGroupExist(groupNameOrID) {
		return fmt.Errorf("Group %s is not present on the system", groupNameOrID)
	}

	return nil
}

func validateInterface(config *knf.Config, prop string, value interface{}) error {
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
