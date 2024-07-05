// Package fs provides KNF validators for checking file-system items
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"

	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/knf"
	"github.com/essentialkaos/ek/v12/path"
	"github.com/essentialkaos/ek/v12/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Perms returns error if config property contains path to object with given
	// permissions
	Perms = validatePerms

	// Owner returns error if config property contains path to object with other
	// owner
	Owner = validateOwner

	// OwnerGroup returns error if config property contains path to object with other
	// owner group
	OwnerGroup = validateOwnerGroup

	// FileMode returns error if config property contains path to object with other
	// file mode
	FileMode = validateFileMode

	// MatchPattern returns error if config property contains path which doesn't match
	// given shell pattern
	MatchPattern = validateMatchPattern
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validatePerms(config knf.IConfig, prop string, value any) error {
	target := config.GetS(prop)

	if target == "" {
		return nil
	}

	var perms string

	switch t := value.(type) {
	case string:
		if t == "" {
			return getValidatorEmptyInputError("Perms", prop)
		}

		perms = t

	default:
		return getValidatorInputError("Perms", prop, value)
	}

	if !fsutil.CheckPerms(perms, target) {
		switch perms {
		case "F":
			return fmt.Errorf("Property %s must be path to file", prop)
		case "FR":
			return fmt.Errorf("Property %s must be path to readable file", prop)
		case "FW":
			return fmt.Errorf("Property %s must be path to writable file", prop)
		case "FX":
			return fmt.Errorf("Property %s must be path to executable file", prop)
		case "FRW":
			return fmt.Errorf("Property %s must be path to readable/writable file", prop)
		case "DX":
			return fmt.Errorf("Property %s must be path to directory", prop)
		case "DRX":
			return fmt.Errorf("Property %s must be path to readable directory", prop)
		case "DWX":
			return fmt.Errorf("Property %s must be path to writable directory", prop)
		case "DRWX":
			return fmt.Errorf("Property %s must be path to readable/writable directory", prop)
		default:
			return fmt.Errorf("Property %s must be path to object with given permissions (%s)", prop, perms)
		}
	}

	return nil
}

func validateOwner(config knf.IConfig, prop string, value any) error {
	target := config.GetS(prop)

	if target == "" {
		return nil
	}

	var owner string

	switch t := value.(type) {
	case string:
		if t == "" {
			return getValidatorEmptyInputError("Owner", prop)
		}

		owner = t

	default:
		return getValidatorInputError("Owner", prop, value)
	}

	user, err := system.LookupUser(owner)

	if err != nil {
		return fmt.Errorf("Can't find user %q on system", owner)
	}

	uid, _, err := fsutil.GetOwner(target)

	if err != nil {
		return fmt.Errorf("Can't get owner for %q", target)
	}

	if user.UID != uid {
		return fmt.Errorf("User %s must be owner of %s", owner, target)
	}

	return nil
}

func validateOwnerGroup(config knf.IConfig, prop string, value any) error {
	target := config.GetS(prop)

	if target == "" {
		return nil
	}

	var ownerGroup string

	switch t := value.(type) {
	case string:
		if t == "" {
			return getValidatorEmptyInputError("OwnerGroup", prop)
		}

		ownerGroup = t

	default:
		return getValidatorInputError("OwnerGroup", prop, value)
	}

	group, err := system.LookupGroup(ownerGroup)

	if err != nil {
		return fmt.Errorf("Can't find group %q on system", ownerGroup)
	}

	_, gid, err := fsutil.GetOwner(target)

	if err != nil {
		return fmt.Errorf("Can't get owner group for %q", target)
	}

	if group.GID != gid {
		return fmt.Errorf("Group %s must be owner of %s", ownerGroup, target)
	}

	return nil
}

func validateFileMode(config knf.IConfig, prop string, value any) error {
	target := config.GetS(prop)

	if target == "" {
		return nil
	}

	var mode os.FileMode

	switch t := value.(type) {
	case os.FileMode:
		if t == 0 {
			return getValidatorEmptyInputError("FileMode", prop)
		}

		mode = t

	default:
		return getValidatorInputError("FileMode", prop, value)
	}

	targetPerms := fsutil.GetMode(target)

	if targetPerms == 0 {
		return fmt.Errorf("Can't get mode for %q", target)
	}

	if mode != targetPerms {
		return fmt.Errorf(
			"%s has different mode (%o != %o)", target, targetPerms, mode)
	}

	return nil
}

func validateMatchPattern(config knf.IConfig, prop string, value any) error {
	target := config.GetS(prop)

	if target == "" {
		return nil
	}

	var pattern string

	switch t := value.(type) {
	case string:
		if t == "" {
			return getValidatorEmptyInputError("MatchPattern", prop)
		}

		pattern = t

	default:
		return getValidatorInputError("MatchPattern", prop, value)
	}

	isMatch, err := path.Match(pattern, target)

	if err != nil {
		return fmt.Errorf("Can't parse shell pattern: %v", err)
	}

	if !isMatch {
		return fmt.Errorf("Property %s must match shell pattern %q", prop, pattern)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getValidatorInputError(validator, prop string, value any) error {
	return fmt.Errorf(
		"Validator fs.%s doesn't support input with type <%T> for checking %s property",
		validator, value, prop,
	)
}

func getValidatorEmptyInputError(validator, prop string) error {
	return fmt.Errorf(
		"Validator fs.%s requires non-empty input for checking %s property",
		validator, prop,
	)
}
