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
	"strconv"
	"strings"
	"sync"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Env is map with environment values
type Env map[string]string

// Variable is environment variable for lazy reading
type Variable struct {
	key   string
	value string
	once  sync.Once
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Var creates new environment variable for lazy reading
func Var(name string) *Variable {
	return &Variable{key: name}
}

// Get returns key-value map with environment values
func Get() Env {
	env := make(Env)

	for _, ev := range os.Environ() {
		k, v, _ := strings.Cut(ev, "=")
		env[k] = v
	}

	return env
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Get returns environment variable value
func (v *Variable) Get() string {
	if v == nil || v.key == "" {
		return ""
	}

	v.once.Do(func() {
		v.value = os.Getenv(v.key)
	})

	return v.value
}

// Is returns true if environment variable value is equal to the given one
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

	v.value, v.once = "", sync.Once{}

	return v
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Path returns the PATH variable as a string slice
func (e Env) Path() []string {
	if e == nil {
		return nil
	}

	return filepath.SplitList(e["PATH"])
}

// Has checks if value with given name is present in env
func (e Env) Has(name string) bool {
	if e == nil || name == "" {
		return false
	}

	_, ok := e[name]

	return ok
}

// Get returns environment variable value as a string
func (e Env) Get(name string) string {
	if e == nil || name == "" {
		return ""
	}

	return e[name]
}

// GetI returns environment variable value as an int
func (e Env) GetI(name string) int {
	if e == nil || name == "" {
		return 0
	}

	value, _ := strconv.Atoi(e[name])

	return value
}

// GetF returns environment variable value as a float
func (e Env) GetF(name string) float64 {
	if e == nil || name == "" {
		return 0.0
	}

	value, _ := strconv.ParseFloat(e[name], 64)

	return value
}
