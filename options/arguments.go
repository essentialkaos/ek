package options

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/path"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Arguments is a slice of parsed non-option command-line arguments
type Arguments []Argument

// Argument is a single non-option command-line argument
type Argument string

// ////////////////////////////////////////////////////////////////////////////////// //

// NewArguments creates an [Arguments] slice from the given strings
func NewArguments(args ...string) Arguments {
	var result Arguments

	for _, arg := range args {
		result = append(result, Argument(arg))
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Has reports whether an argument exists at the given index and is non-empty
func (a Arguments) Has(index int) bool {
	return index >= 0 && index < len(a) && a[index] != ""
}

// Get returns the argument at the given index, or an empty [Argument] if out of range
func (a Arguments) Get(index int) Argument {
	if index < 0 || index >= len(a) {
		return ""
	}

	return a[index]
}

// Last returns the last argument, or an empty Argument if the slice is empty
func (a Arguments) Last() Argument {
	if len(a) == 0 {
		return ""
	}

	return a[len(a)-1]
}

// Append returns a new [Arguments] slice with the given strings appended
func (a Arguments) Append(args ...string) Arguments {
	result := slices.Clone(a)

	for _, arg := range args {
		result = append(result, Argument(arg))
	}

	return result
}

// Unshift returns a new Arguments slice with the given strings prepended
func (a Arguments) Unshift(args ...string) Arguments {
	var result Arguments

	for _, arg := range args {
		result = append(result, Argument(arg))
	}

	return append(result, a...)
}

// Flatten joins all arguments into a single space-separated string
func (a Arguments) Flatten() string {
	if len(a) == 0 {
		return ""
	}

	var result strings.Builder

	for _, arg := range a {
		result.WriteString(string(arg))
		result.WriteRune(' ')
	}

	return result.String()[:result.Len()-1]
}

// Strings converts the [Arguments] slice to a plain []string
func (a Arguments) Strings() []string {
	var result []string

	for _, arg := range a {
		result = append(result, string(arg))
	}

	return result
}

// Filter returns arguments whose base filename matches the given glob pattern.
// Arguments that are unexpanded globs themselves are skipped.
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

// ToLower returns the argument converted to lower case
func (a Argument) ToLower() Argument {
	return Argument(strings.ToLower(string(a)))
}

// ToUpper returns the argument converted to upper case
func (a Argument) ToUpper() Argument {
	return Argument(strings.ToUpper(string(a)))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns the argument as a plain string
func (a Argument) String() string {
	return string(a)
}

// Is reports whether the argument equals the given value after type-appropriate
// conversion
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
	case uint:
		v, err := a.Uint()
		return v == t && err == nil
	case uint64:
		v, err := a.Uint64()
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

// Int converts the argument to int
func (a Argument) Int() (int, error) {
	return strconv.Atoi(string(a))
}

// Int64 converts the argument to int64
func (a Argument) Int64() (int64, error) {
	return strconv.ParseInt(string(a), 10, 64)
}

// Uint converts the argument to uint
func (a Argument) Uint() (uint, error) {
	u, err := strconv.ParseUint(string(a), 10, 64)
	return uint(min(u, math.MaxUint)), err
}

// Uint64 converts the argument to uint64
func (a Argument) Uint64() (uint64, error) {
	return strconv.ParseUint(string(a), 10, 64)
}

// Float converts the argument to float64
func (a Argument) Float() (float64, error) {
	return strconv.ParseFloat(string(a), 64)
}

// Bool converts the argument to bool, accepting "true/false", "yes/no", "y/n",
// and "1/0"
func (a Argument) Bool() (bool, error) {
	switch strings.ToLower(string(a)) {
	case "true", "yes", "y", "1":
		return true, nil
	case "false", "no", "n", "0", "":
		return false, nil
	}

	return false, fmt.Errorf("unsupported boolean value %q", a)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Base returns the last element of the argument's path; shorthand for [path.Base]
func (a Argument) Base() Argument {
	return Argument(path.Base(string(a)))
}

// Clean returns the cleaned path form of the argument; shorthand for [path.Clean]
func (a Argument) Clean() Argument {
	return Argument(path.Clean(string(a)))
}

// Dir returns the directory portion of the argument's path; shorthand for [path.Dir]
func (a Argument) Dir() Argument {
	return Argument(path.Dir(string(a)))
}

// Ext returns the file extension of the argument; shorthand for [path.Ext]
func (a Argument) Ext() Argument {
	return Argument(path.Ext(string(a)))
}

// IsAbs reports whether the argument is an absolute path
func (a Argument) IsAbs() bool {
	return path.IsAbs(string(a))
}

// Match reports whether the argument matches the given glob pattern; shorthand
// for [path.Match]
func (a Argument) Match(pattern string) (bool, error) {
	return path.Match(pattern, string(a))
}
