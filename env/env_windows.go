// Package env provides methods for working with environment variables
package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Env is map with environment values
type Env map[string]string

// ❗ Variable is environment variable for lazy reading
type Variable struct {
	key    string
	value  string
	isRead bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Var creates new environment variable struct
func Var(name string) *Variable {
	panic("UNSUPPORTED")
}

// ❗ Get return key-value map with environment values
func Get() Env {
	panic("UNSUPPORTED")
}

// ❗ Which find full path to some app
func Which(name string) string {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Get returns environment variable value
func (v *Variable) Get() string {
	panic("UNSUPPORTED")
}

// ❗ Is returns true if environment variable value is equal to given one
func (v *Variable) Is(value string) bool {
	panic("UNSUPPORTED")
}

// ❗ String returns environment variable value as string
func (v *Variable) String() string {
	panic("UNSUPPORTED")
}

// ❗ Reset resets reading state of variable
func (v *Variable) Reset() *Variable {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Path return path as string slice
func (e Env) Path() []string {
	panic("UNSUPPORTED")
}

// ❗ GetS return environment variable value as string
func (e Env) GetS(name string) string {
	panic("UNSUPPORTED")
}

// ❗ GetI return environment variable value as int
func (e Env) GetI(name string) int {
	panic("UNSUPPORTED")
}

// ❗ GetF return environment variable value as float
func (e Env) GetF(name string) float64 {
	panic("UNSUPPORTED")
}
