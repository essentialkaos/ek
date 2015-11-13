// +build !windows

package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strings"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Env is map with environment values
type Env map[string]string

// ////////////////////////////////////////////////////////////////////////////////// //

// Get return key-value map with environment values
func Get() Env {
	env := make(Env)

	for _, ev := range os.Environ() {
		evs := strings.Split(ev, "=")
		k, v := evs[0], evs[1]

		env[k] = v
	}

	return env
}

// Which find full path to some app
func Which(name string) string {
	paths := Get().Path()

	for _, path := range paths {
		if syscall.Access(path+"/"+name, syscall.F_OK) == nil {
			return path + "/" + name
		}
	}

	return ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Path return path as string slice
func (e Env) Path() []string {
	return strings.Split(e["PATH"], ":")
}
