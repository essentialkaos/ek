package options

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/path"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Arguments is a slice with with command argument
type Arguments []Argument

// Argument is command argument
type Argument string

// ////////////////////////////////////////////////////////////////////////////////// //

// NewArguments creates new arguments slice from given strings
func NewArguments(args ...string) Arguments {
	var result Arguments

	for _, arg := range args {
		result = append(result, Argument(arg))
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Has returns true if arguments contains argument with given index
func (a Arguments) Has(index int) bool {
	return index < len(a) && a[index] != ""
}

// Get returns argument with given index
func (a Arguments) Get(index int) Argument {
	if index >= len(a) {
		return ""
	}

	return a[index]
}

// Get returns the last argument
func (a Arguments) Last() Argument {
	if len(a) == 0 {
		return ""
	}

	return a[len(a)-1]
}

// Append adds arguments to the end of the arguments slices
func (a Arguments) Append(args ...string) Arguments {
	var result Arguments

	result = append(Arguments{}, a...)

	for _, arg := range args {
		result = append(result, Argument(arg))
	}

	return result
}

// Unshift adds arguments to the beginning of the arguments slices
func (a Arguments) Unshift(args ...string) Arguments {
	var result Arguments

	for _, arg := range args {
		result = append(result, Argument(arg))
	}

	return append(result, a...)
}

// Strings converts arguments to slice with strings
func (a Arguments) Strings() []string {
	var result []string

	for _, arg := range a {
		result = append(result, string(arg))
	}

	return result
}

// Filter filters arguments by a given glob pattern. This method works only with
// files. It means that for a given path only the last part will be checked for
// pattern matching.
func (a Arguments) Filter(pattern string) Arguments {
	var result Arguments

	for _, arg := range a {
		// Skip all unexpanded globs
		if path.IsGlob(arg.String()) {
			continue
		}

		ok, _ := arg.Base().Match(pattern)

		if ok {
			result = append(result, arg)
		}
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ToLower returns argument converted to lower case
func (a Argument) ToLower() Argument {
	return Argument(strings.ToLower(string(a)))
}

// ToUpper returns argument converted to upper case
func (a Argument) ToUpper() Argument {
	return Argument(strings.ToUpper(string(a)))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String converts argument to string
func (a Argument) String() string {
	return string(a)
}

// Is returns true if argument equals to given value
func (a Argument) Is(value any) bool {
	switch t := value.(type) {
	case string:
		return a.String() == t
	case int:
		v, err := a.Int()
		return v == t && err == nil
	case int64:
		v, err := a.Int64()
		return v == t && err == nil
	case uint64:
		v, err := a.Uint()
		return v == t && err == nil
	case float64:
		v, err := a.Float()
		return v == t && err == nil
	case bool:
		v, err := a.Bool()
		return v == t && err == nil
	}

	return false
}

// Int converts argument to int
func (a Argument) Int() (int, error) {
	return strconv.Atoi(string(a))
}

// Int64 converts argument to int64
func (a Argument) Int64() (int64, error) {
	return strconv.ParseInt(string(a), 10, 64)
}

// Uint converts argument to uint
func (a Argument) Uint() (uint64, error) {
	return strconv.ParseUint(string(a), 10, 64)
}

// Int converts argument to int
func (a Argument) Float() (float64, error) {
	return strconv.ParseFloat(string(a), 64)
}

// Int converts argument to int
func (a Argument) Bool() (bool, error) {
	switch strings.ToLower(string(a)) {
	case "true", "yes", "y", "1":
		return true, nil
	case "false", "no", "n", "0", "":
		return false, nil
	}

	return false, fmt.Errorf("Unsupported boolean value %q", a)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Base is shorthand analog of path.Base
func (a Argument) Base() Argument {
	return Argument(path.Base(string(a)))
}

// Clean is shorthand analog of path.Clean
func (a Argument) Clean() Argument {
	return Argument(path.Clean(string(a)))
}

// Dir is shorthand analog of path.Dir
func (a Argument) Dir() Argument {
	return Argument(path.Dir(string(a)))
}

// Ext is shorthand analog of path.Ext
func (a Argument) Ext() Argument {
	return Argument(path.Ext(string(a)))
}

// IsAbs is shorthand analog of path.IsAbs
func (a Argument) IsAbs() bool {
	return path.IsAbs(string(a))
}

// Match is shorthand analog of path.Match
func (a Argument) Match(pattern string) (bool, error) {
	return path.Match(pattern, string(a))
}
