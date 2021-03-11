// +build !linux, !darwin, windows

package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Env is map with environment values
type Env map[string]string

// ////////////////////////////////////////////////////////////////////////////////// //

// Get return key-value map with environment values
func Get() Env {
	return Env{}
}

// Which find full path to some app
func Which(name string) string {
	return ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Path return path as string slice
func (e Env) Path() []string {
	return []string{}
}

// GetS return environment variable value as string
func (e Env) GetS(name string) string {
	return e[name]
}

// GetI return environment variable value as int
func (e Env) GetI(name string) int {
	return -1
}

// GetF return environment variable value as float
func (e Env) GetF(name string) float64 {
	return -1.0
}
