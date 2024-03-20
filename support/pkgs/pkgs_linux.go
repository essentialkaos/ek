// Package pkgs provides methods for collecting information about installed packages
package pkgs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os/exec"
	"strings"

	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/strutil"

	"github.com/essentialkaos/ek/v12/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_RPM uint8 = 1
	_DEB uint8 = 2
	_APK uint8 = 3
)

// ////////////////////////////////////////////////////////////////////////////////// //

var pkgManager uint8

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collect info about packages
func Collect(pkgs ...string) []support.Pkg {
	var result []support.Pkg

	if pkgManager == 0 {
		pkgManager = getPackageManagerType()
	}

	for _, pkg := range pkgs {
		result = append(result, getPackageInfo(pkg))
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getPackageManagerType returns type of package manager rpm/deb
func getPackageManagerType() uint8 {
	switch {
	case fsutil.IsExist("/usr/bin/rpm"):
		return _RPM
	case fsutil.IsExist("/usr/bin/dpkg-query"):
		return _DEB
	case fsutil.IsExist("/sbin/apk"):
		return _APK
	}

	return 0
}

// getPackageInfo returns info about package
func getPackageInfo(names string) support.Pkg {
	var info support.Pkg
	var firstPackage string

	for _, pkgName := range strutil.Fields(names) {
		if firstPackage == "" {
			firstPackage = pkgName
		}

		switch pkgManager {
		case _RPM:
			info = getRPMPackageInfo(pkgName)
		case _DEB:
			info = getDEBPackageInfo(pkgName)
		case _APK:
			info = getAPKPackageInfo(pkgName)
		}

		if info.Version != "" {
			return info
		}
	}

	return support.Pkg{firstPackage, ""}
}

// getRPMPackageInfo returns info about RPM package
func getRPMPackageInfo(name string) support.Pkg {
	cmd := exec.Command("rpm", "-q", name)
	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return support.Pkg{name, ""}
	}

	return support.Pkg{name, strings.TrimRight(string(out), "\n\r")}
}

// getDEBPackageInfo returns info about DEB package
func getDEBPackageInfo(name string) support.Pkg {
	cmd := exec.Command("dpkg-query", "--show", "--showformat=${Package}-${Version}", name)
	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return support.Pkg{name, ""}
	}

	return support.Pkg{name, string(out)}
}

// getAPKPackageInfo returns info about APK package
func getAPKPackageInfo(name string) support.Pkg {
	cmd := exec.Command("apk", "list", "--installed", name)
	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return support.Pkg{name, ""}
	}

	return support.Pkg{
		name,
		strutil.ReadField(strings.TrimRight(string(out), "\n\r"), 0, false, ' '),
	}
}
