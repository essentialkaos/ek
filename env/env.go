// Package env provides methods for working with environment variables
package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Env is map with environment values
type Env map[string]string

// Variable is environment variable for lazy reading
type Variable struct {
	key    string
	value  string
	isRead bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Var creates new environment variable struct
func Var(name string) *Variable {
	return &Variable{key: name}
}

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

// ////////////////////////////////////////////////////////////////////////////////// //

// Get returns environment variable value
func (v *Variable) Get() string {
	if v == nil {
		return ""
	}

	if !v.isRead {
		v.value = os.Getenv(v.key)
		v.isRead = true
	}

	return v.value
}

// Is returns true if environment variable value is equal to given one
func (v *Variable) Is(value string) bool {
	return v.Get() == value
}

// String returns environment variable value as string
func (v *Variable) String() string {
	return v.Get()
}

// Reset resets reading state of variable
func (v *Variable) Reset() *Variable {
	if v == nil {
		return nil
	}

	v.value, v.isRead = "", false

	return v
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
