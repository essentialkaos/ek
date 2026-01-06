// Package system provides KNF validators for checking system items (user, groups,
// network interfaces)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"github.com/essentialkaos/ek/v13/knf"
	"github.com/essentialkaos/ek/v13/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// User returns error if configuration property contains name of user or UID which not
	// present on the system
	User = validateUser

	// Group returns error if configuration property contains name of group or GID which
	// not present on the system
	Group = validateGroup
)

// ////////////////////////////////////////////////////////////////////////////////// //

// validateUser checks if user exists on the system
func validateUser(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v != "" && !system.IsUserExist(v) {
		return fmt.Errorf("User %q is not present on the system", v)
	}

	return nil
}

// validateGroup checks if group exists on the system
func validateGroup(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v != "" && !system.IsGroupExist(v) {
		return fmt.Errorf("Group %q is not present on the system", v)
	}

	return nil
}
