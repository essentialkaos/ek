// Package system provides KNF validators for checking system items (user, groups,
// network interfaces)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
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
