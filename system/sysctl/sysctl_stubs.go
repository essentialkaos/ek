//go:build !linux && !darwin

package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ All returns all kernel parameters available on the current system
func All() (Params, error) {
	panic("UNSUPPORTED")
}

// ❗ Get returns the kernel parameter with the given name. The name must be dot-separated
// (e.g. "kernel.pid_max") and must not contain spaces or slashes.
func Get(name string) (Param, error) {
	panic("UNSUPPORTED")
}
