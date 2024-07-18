// Package apps provides methods for obtaining version information about various tools
package apps

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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

// Golang extracts version info from go command output
func Golang() support.App {
	ver := extractField(execVersionCmd("go", "version"), 0, 2)
	ver = strutil.Exclude(ver, "go")

	return support.App{"golang", ver}
}

// GCC extracts version info from gcc command output
func GCC() support.App {
	ver := extractField(execVersionCmd("gcc", "--version"), 0, 2)
	return support.App{"gcc", ver}
}

// Clang extracts version info from clang command output
func Clang() support.App {
	ver := extractField(execVersionCmd("clang", "--version"), 0, 2)
	return support.App{"clang", ver}
}

// Python3 extracts version info from python 3.x command output
func Python3() support.App {
	ver := extractField(execVersionCmd("python3", "--version"), 0, 1)
	return support.App{"python3", ver}
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
	ver := extractField(execVersionCmd("groovy", "--version"), 0, 2)
	return support.App{"groovy", ver}
}

// Ruby extracts version info from ruby command output
func Ruby() support.App {
	ver := extractField(execVersionCmd("ruby", "--version"), 0, 1)
	return support.App{"ruby", ver}
}

// NodeJS extracts version info from node command output
func NodeJS() support.App {
	ver := extractField(execVersionCmd("node", "--version"), 0, 0)
	ver = strings.TrimLeft(ver, "v")

	return support.App{"nodejs", ver}
}

// Rust extracts version info from rust command output
func Rust() support.App {
	ver := extractField(execVersionCmd("rustc", "--version"), 0, 1)
	return support.App{"rust", ver}
}

// PHP extracts version info from php command output
func PHP() support.App {
	ver := extractField(execVersionCmd("php", "--version"), 0, 1)
	return support.App{"php", ver}
}

// Bash extracts version info from bash command output
func Bash() support.App {
	ver := extractField(execVersionCmd("bash", "--version"), 0, 3)
	return support.App{"bash", ver}
}

// Git extracts version info from git command output
func Git() support.App {
	ver := extractField(execVersionCmd("git", "--version"), 0, 2)
	return support.App{"git", ver}
}

// Mercurial extracts version info from hg command output
func Mercurial() support.App {
	ver := extractField(execVersionCmd("hg", "--version"), 0, 4)
	ver = strings.Trim(ver, "()")

	return support.App{"mercurial", ver}
}

// SVN extracts version info from svn command output
func SVN() support.App {
	ver := extractField(execVersionCmd("svn", "--version"), 0, 2)
	return support.App{"svn", ver}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// execVersionCmd execs command and returns output as a string
func execVersionCmd(name string, args ...string) string {
	c := exec.Command(name, args...)
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
