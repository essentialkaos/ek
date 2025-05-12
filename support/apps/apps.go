// Package apps provides methods for obtaining version information about various tools
package apps

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os/exec"
	"strings"

	"github.com/essentialkaos/ek/v13/strutil"

	"github.com/essentialkaos/ek/v13/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ExtractVersion extracts version info from given command
func ExtractVersion(cmd []string, line, field int) support.App {
	if len(cmd) == 0 {
		return support.App{}
	}

	return support.App{cmd[0], extractField(execVersionCmd(cmd...), line, field)}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Golang extracts version info from go command output
func Golang() support.App {
	info := ExtractVersion([]string{"go", "version"}, 0, 2)
	info.Version = strutil.Exclude(info.Version, "go")

	return info
}

// GCC extracts version info from gcc command output
func GCC() support.App {
	return ExtractVersion([]string{"gcc", "--version"}, 0, 2)
}

// Clang extracts version info from clang command output
func Clang() support.App {
	return ExtractVersion([]string{"clang", "--version"}, 0, 2)
}

// Python3 extracts version info from python 3.x command output
func Python3() support.App {
	return ExtractVersion([]string{"python3", "--version"}, 0, 1)
}

// Java extracts version info from java command output
func Java() support.App {
	ver := execVersionCmd("java", "-version")

	if strings.Contains(ver, "Temurin") {
		ver = extractField(ver, 1, 3)
	} else {
		ver = extractField(ver, 0, 2)
	}

	return support.App{"java", ver}
}

// Groovy extracts version info from groovy command output
func Groovy() support.App {
	return ExtractVersion([]string{"groovy", "--version"}, 0, 2)
}

// Ruby extracts version info from ruby command output
func Ruby() support.App {
	return ExtractVersion([]string{"ruby", "--version"}, 0, 1)
}

// NodeJS extracts version info from node command output
func NodeJS() support.App {
	info := ExtractVersion([]string{"node", "--version"}, 0, 2)
	info.Version = strings.TrimLeft(info.Version, "v")

	return info
}

// Rust extracts version info from rust command output
func Rust() support.App {
	return ExtractVersion([]string{"rustc", "--version"}, 0, 1)
}

// PHP extracts version info from php command output
func PHP() support.App {
	return ExtractVersion([]string{"php", "--version"}, 0, 1)
}

// Bash extracts version info from bash command output
func Bash() support.App {
	return ExtractVersion([]string{"bash", "--version"}, 0, 3)
}

// Git extracts version info from git command output
func Git() support.App {
	return ExtractVersion([]string{"git", "--version"}, 0, 2)
}

// Mercurial extracts version info from hg command output
func Mercurial() support.App {
	info := ExtractVersion([]string{"hg", "--version"}, 0, 4)
	info.Version = strings.Trim(info.Version, "()")

	return info
}

// SVN extracts version info from svn command output
func SVN() support.App {
	return ExtractVersion([]string{"svn", "--version"}, 0, 2)
}

// Docker extracts version info from Docker command output
func Docker() support.App {
	info := ExtractVersion([]string{"docker", "--version"}, 0, 2)
	info.Version = strings.TrimRight(info.Version, ",")

	return info
}

// Podman extracts version info from Podman command output
func Podman() support.App {
	return ExtractVersion([]string{"podman", "--version"}, 0, 2)
}

// LXC extracts version info from LXC command output
func LXC() support.App {
	return ExtractVersion([]string{"lxc", "--version"}, 0, 0)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// execVersionCmd execs command and returns output as a string
func execVersionCmd(cmd ...string) string {
	c := exec.Command(cmd[0], cmd[1:]...)
	output, err := c.CombinedOutput()

	if err != nil {
		return ""
	}

	return string(output)
}

// extractField extracts version data from given string
func extractField(data string, line, field int) string {
	if len(data) == 0 {
		return ""
	}

	lineData := strutil.ReadField(data, line, false, '\n')

	return strings.Trim(strutil.ReadField(lineData, field, false, ' '), "\"\n\r\t")
}
