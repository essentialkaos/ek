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

	"github.com/essentialkaos/ek/v12/knf"
	"github.com/essentialkaos/ek/v12/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// User returns error if config property contains name of user or UID which not
	// present on the system
	User = validateUser

	// Group returns error if config property contains name of group or GID which not
	// present on the system
	Group = validateGroup
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validateUser(config knf.IConfig, prop string, value any) error {
	userNameOrID := config.GetS(prop)

	if userNameOrID == "" {
		return nil
	}

	if !system.IsUserExist(userNameOrID) {
		return fmt.Errorf("User %q is not present on the system", userNameOrID)
	}

	return nil
}

func validateGroup(config knf.IConfig, prop string, value any) error {
	groupNameOrID := config.GetS(prop)

	if groupNameOrID == "" {
		return nil
	}

	if !system.IsGroupExist(groupNameOrID) {
		return fmt.Errorf("Group %q is not present on the system", groupNameOrID)
	}

	return nil
}
