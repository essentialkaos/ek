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
	_RPM  uint8 = 1
	_DPKG uint8 = 2
	_APK  uint8 = 3
	_TDNF uint8 = 4
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
		return _DPKG
	case fsutil.IsExist("/sbin/apk"):
		return _APK
	case fsutil.IsExist("/usr/bin/tdnf"):
		return _TDNF
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
		case _DPKG:
			info = getDPKGPackageInfo(pkgName)
		case _APK:
			info = getAPKPackageInfo(pkgName)
		}

		if info.Version != "" {
			return info
		}
	}

	return support.Pkg{firstPackage, ""}
}

// getRPMPackageInfo returns info about package from rpm
func getRPMPackageInfo(name string) support.Pkg {
	cmd := exec.Command("rpm", "-q", "--qf", "%{version}.%{release}", name)
	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return support.Pkg{name, ""}
	}

	return support.Pkg{name, strings.TrimRight(string(out), "\n\r")}
}

// getDPKGPackageInfo returns info about package from dpkg
func getDPKGPackageInfo(name string) support.Pkg {
	cmd := exec.Command("dpkg-query", "--show", "--showformat=${Version}", name)
	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return support.Pkg{name, ""}
	}

	return support.Pkg{name, string(out)}
}

// getAPKPackageInfo returns info about package from apk
func getAPKPackageInfo(name string) support.Pkg {
	cmd := exec.Command("apk", "list", "--installed", name)
	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return support.Pkg{name, ""}
	}

	ver := strutil.ReadField(strings.TrimRight(string(out), "\n\r"), 0, false, ' ')
	ver = strings.Replace(ver, name+"-", "", 1)

	return support.Pkg{name, ver}
}

// getTDNFPackageInfo returns info about package from tndf
func getTDNFPackageInfo(name string) support.Pkg {
	cmd := exec.Command("tdnf", "repoquery", "--installed", name)
	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return support.Pkg{name, ""}
	}

	ver := strutil.ReadField(strings.TrimRight(string(out), "\n\r"), 0, false, ' ')
	ver = strings.Replace(ver, name+"-", "", 1)

	return support.Pkg{name, ver}
}
