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
	_AUR  uint8 = 5
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
	case fsutil.IsExist("/usr/bin/pacman"):
		return _AUR
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
		case _TDNF:
			info = getTDNFPackageInfo(pkgName)
		case _AUR:
			info = getAURPackageInfo(pkgName)
		}

		if info.Version != "" {
			return info
		}
	}

	return support.Pkg{firstPackage, ""}
}

// getRPMPackageInfo returns info about package from rpm
func getRPMPackageInfo(name string) support.Pkg {
	out := getCommandOutput("rpm", "-q", "--qf", "%{version}-%{release}.%{arch}", name)

	if len(out) == 0 {
		return support.Pkg{name, ""}
	}

	return support.Pkg{name, out}
}

// getDPKGPackageInfo returns info about package from dpkg
func getDPKGPackageInfo(name string) support.Pkg {
	out := getCommandOutput("dpkg-query", "--show", "--showformat=${Version}.${Architecture}", name)

	if len(out) == 0 {
		return support.Pkg{name, ""}
	}

	return support.Pkg{name, out}
}

// getAPKPackageInfo returns info about package from apk
func getAPKPackageInfo(name string) support.Pkg {
	out := getCommandOutput("apk", "list", "--installed", name)

	if len(out) == 0 {
		return support.Pkg{name, ""}
	}

	ver := strutil.ReadField(out, 0, false, ' ')
	ver = strings.Replace(ver, name+"-", "", 1)
	arch := strutil.ReadField(out, 1, false, ' ')

	return support.Pkg{name, ver + "." + arch}
}

// getTDNFPackageInfo returns info about package from tndf
func getTDNFPackageInfo(name string) support.Pkg {
	out := getCommandOutput("tdnf", "repoquery", "--installed", name)

	if len(out) == 0 {
		return support.Pkg{name, ""}
	}

	ver := strutil.ReadField(out, 0, false, ' ')
	ver = strings.Replace(ver, name+"-", "", 1)

	return support.Pkg{name, ver}
}

// getAURPackageInfo returns info about package from pacman
func getAURPackageInfo(name string) support.Pkg {
	out := getCommandOutput("pacman", "-Q", name)

	if len(out) == 0 {
		return support.Pkg{name, ""}
	}

	return support.Pkg{name, strutil.ReadField(out, 1, false, ' ')}
}

// getCommandOutput runs command and returns output as a string
func getCommandOutput(cmd string, args ...string) string {
	c := exec.Command(cmd, args...)
	out, err := c.Output()

	if err != nil {
		return ""
	}

	return strings.TrimRight(string(out), "\n\r")
}
