// Package fs provides KNF validators for checking file-system items
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"

	"pkg.re/essentialkaos/ek.v11/fsutil"
	"pkg.re/essentialkaos/ek.v11/knf"
	"pkg.re/essentialkaos/ek.v11/system"
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
)

// ////////////////////////////////////////////////////////////////////////////////// //

func validatePerms(config *knf.Config, prop string, value interface{}) error {
	perms := value.(string)
	target := config.GetS(prop)

	if target == "" {
		return nil
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

func validateOwner(config *knf.Config, prop string, value interface{}) error {
	target := config.GetS(prop)
	owner := value.(string)

	if target == "" {
		return nil
	}

	user, err := system.LookupUser(owner)

	if err != nil {
		return fmt.Errorf("Can't find user %s on system", owner)
	}

	uid, _, err := fsutil.GetOwner(target)

	if err != nil {
		return fmt.Errorf("Can't get owner for %s", target)
	}

	if user.UID != uid {
		return fmt.Errorf("User %s must be owner of %s", owner, target)
	}

	return nil
}

func validateOwnerGroup(config *knf.Config, prop string, value interface{}) error {
	target := config.GetS(prop)
	ownerGroup := value.(string)

	if target == "" {
		return nil
	}

	group, err := system.LookupGroup(ownerGroup)

	if err != nil {
		return fmt.Errorf("Can't find group %s on system", ownerGroup)
	}

	_, gid, err := fsutil.GetOwner(target)

	if err != nil {
		return fmt.Errorf("Can't get owner group for %s", target)
	}

	if group.GID != gid {
		return fmt.Errorf("Group %s must be owner of %s", ownerGroup, target)
	}

	return nil
}

func validateFileMode(config *knf.Config, prop string, value interface{}) error {
	perms := value.(os.FileMode)
	target := config.GetS(prop)

	if target == "" {
		return nil
	}

	targetPerms := fsutil.GetMode(target)

	if targetPerms == 0 {
		return fmt.Errorf("Can't get mode for %s", target)
	}

	if perms != targetPerms {
		return fmt.Errorf("%s has different mode (%d != %d)", target, targetPerms, perms)
	}

	return nil
}
