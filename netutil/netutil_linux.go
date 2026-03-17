package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"os"

	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with routes info in procfs
var procRouteFile = "/proc/net/route"

// ////////////////////////////////////////////////////////////////////////////////// //

// getDefaultRouteInterface returns the name of the default route interface
func getDefaultRouteInterface() string {
	fd, err := os.OpenFile(procRouteFile, os.O_RDONLY, 0)

	if err != nil {
		return ""
	}

	defer fd.Close()

	s := bufio.NewScanner(fd)

	var header bool

	for s.Scan() {
		if !header {
			header = true
			continue
		}

		line := s.Text()

		if strutil.ReadField(line, 1, true) == "00000000" {
			return strutil.ReadField(line, 0, true)
		}
	}

	return ""
}
