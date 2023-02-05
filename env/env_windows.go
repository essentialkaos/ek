// Package env provides methods for working with environment variables
package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Env is map with environment values
type Env map[string]string

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Get return key-value map with environment values
func Get() Env {
	panic("UNSUPPORTED")
	return Env{}
}

// ❗ Which find full path to some app
func Which(name string) string {
	panic("UNSUPPORTED")
	return ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Path return path as string slice
func (e Env) Path() []string {
	panic("UNSUPPORTED")
	return nil
}

// ❗ GetS return environment variable value as string
func (e Env) GetS(name string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ GetI return environment variable value as int
func (e Env) GetI(name string) int {
	panic("UNSUPPORTED")
	return 0
}

// ❗ GetF return environment variable value as float
func (e Env) GetF(name string) float64 {
	panic("UNSUPPORTED")
	return 0.0
}
