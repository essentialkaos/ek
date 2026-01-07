// Package fs provides KNF validators for checking file-system items
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"

	"github.com/essentialkaos/ek/v13/fsutil"
	"github.com/essentialkaos/ek/v13/knf"
	"github.com/essentialkaos/ek/v13/path"
	"github.com/essentialkaos/ek/v13/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// Perms returns error if configuration property contains path to object with given
	// permissions. Supported permissions: F, FR, FW, FX, FRW, DX, DRX, DWX, DRWX
	Perms = validatePerms

	// Owner returns error if configuration property contains path to object with other
	// owner
	Owner = validateOwner

	// OwnerGroup returns error if configuration property contains path to object with
	// other owner group
	OwnerGroup = validateOwnerGroup

	// FileMode returns error if configuration property contains path to object with
	// other file mode
	FileMode = validateFileMode

	// MatchPattern returns error if configuration property contains path which doesn't
	// match given shell pattern (e.g. "*.txt", "/etc/*", etc.)
	MatchPattern = validateMatchPattern
)

// ////////////////////////////////////////////////////////////////////////////////// //

// validatePerms checks if configuration property contains path to object with given
// permissions. Supported permissions: F, FR, FW, FX, FRW, DX, DRX, DWX, DRWX
func validatePerms(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
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

	if !fsutil.CheckPerms(perms, v) {
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

// validateOwner checks if configuration property contains path to object with given
// owner
func validateOwner(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
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

	uid, _, err := fsutil.GetOwner(v)

	if err != nil {
		return fmt.Errorf("Can't get owner for %q", v)
	}

	if user.UID != uid {
		return fmt.Errorf("User %s must be owner of %s", owner, v)
	}

	return nil
}

// validateOwnerGroup checks if configuration property contains path to object with
// given owner group
func validateOwnerGroup(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
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

	_, gid, err := fsutil.GetOwner(v)

	if err != nil {
		return fmt.Errorf("Can't get owner group for %q", v)
	}

	if group.GID != gid {
		return fmt.Errorf("Group %s must be owner of %s", ownerGroup, v)
	}

	return nil
}

// validateFileMode checks if configuration property contains path to object with
// given file mode
func validateFileMode(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
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

	targetPerms := fsutil.GetMode(v)

	if targetPerms == 0 {
		return fmt.Errorf("Can't get mode for %q", v)
	}

	if mode != targetPerms {
		return fmt.Errorf(
			"%s has different mode (%o != %o)", v, targetPerms, mode)
	}

	return nil
}

// validateMatchPattern checks if configuration property contains path which
// match given shell pattern (e.g. "*.txt", "/etc/*", etc.)
func validateMatchPattern(config knf.IConfig, prop string, value any) error {
	v := config.GetS(prop)

	if v == "" {
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

	isMatch, err := path.Match(pattern, v)

	if err != nil {
		return fmt.Errorf("Can't parse shell pattern: %v", err)
	}

	if !isMatch {
		return fmt.Errorf("Property %s must match shell pattern %q", prop, pattern)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getValidatorInputError returns error for unsupported input type
func getValidatorInputError(validator, prop string, value any) error {
	return fmt.Errorf(
		"Validator fs.%s doesn't support input with type <%T> for checking %s property",
		validator, value, prop,
	)
}

// getValidatorEmptyInputError returns error for empty input
func getValidatorEmptyInputError(validator, prop string) error {
	return fmt.Errorf(
		"Validator fs.%s requires non-empty input for checking %s property",
		validator, prop,
	)
}
