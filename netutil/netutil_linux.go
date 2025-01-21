package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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

func getDefaultRouteInterface() string {
	fd, err := os.OpenFile(procRouteFile, os.O_RDONLY, 0)

	if err != nil {
		return ""
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	var header bool

	for s.Scan() {
		if !header {
			header = true
			continue
		}

		if strutil.ReadField(s.Text(), 1, true) == "00000000" {
			return strutil.ReadField(s.Text(), 0, true)
		}
	}

	return ""
}
