//go:build !windows
// +build !windows

// Package env provides methods for working with environment variables
package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strconv"
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

// GetS return environment variable value as string
func (e Env) GetS(name string) string {
	return e[name]
}

// GetI return environment variable value as int
func (e Env) GetI(name string) int {
	value, err := strconv.Atoi(e[name])

	if err != nil {
		return -1
	}

	return value
}

// GetF return environment variable value as float
func (e Env) GetF(name string) float64 {
	value, err := strconv.ParseFloat(e[name], 64)

	if err != nil {
		return -1.0
	}

	return value
}
