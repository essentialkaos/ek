//go:build !windows

// Package env provides methods for working with environment variables
package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"path/filepath"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Which find full path to some app
func Which(name string) string {
	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		if syscall.Access(path+"/"+name, syscall.F_OK) == nil {
			return path + "/" + name
		}
	}

	return ""
}
