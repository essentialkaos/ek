// Package united provides KNF configuration extended by enviroment variables and options
package united

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"time"

	"github.com/essentialkaos/ek/v12/knf"
	"github.com/essentialkaos/ek/v12/knf/value"
	"github.com/essentialkaos/ek/v12/options"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Config is extended configuration
type config struct {
	mappings map[string]Mapping
}

// Mapping contains mapping [knf property] → [option] → [envvar]
type Mapping struct {
	Property string // Property from KNF configuration file
	Option   string // Command-line option
	Variable string // Environment variable
}

// ////////////////////////////////////////////////////////////////////////////////// //

// DurationMod is type for duration modificator
type DurationMod = knf.DurationMod

// ////////////////////////////////////////////////////////////////////////////////// //

var global *config

var optionHasFunc = options.Has
var optionGetFunc = options.GetS

// ////////////////////////////////////////////////////////////////////////////////// //

// Combine applies mappings to combine knf properties, options, and environment
// variables
func Combine(mappings ...Mapping) error {
	config := &config{
		mappings: make(map[string]Mapping),
	}

	for _, m := range mappings {
		config.mappings[m.Property] = m
	}

	global = config

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetS returns configuration value as string
func GetS(name string, defvals ...string) string {
	return global.GetS(name, defvals...)
}

// GetI returns configuration value as int
func GetI(name string, defvals ...int) int {
	return global.GetI(name, defvals...)
}

// GetI64 returns configuration value as int64
func GetI64(name string, defvals ...int64) int64 {
	return global.GetI64(name, defvals...)
}

// GetU returns configuration value as uint
func GetU(name string, defvals ...uint) uint {
	return global.GetU(name, defvals...)
}

// GetU64 returns configuration value as uint64
func GetU64(name string, defvals ...uint64) uint64 {
	return global.GetU64(name, defvals...)
}

// GetF returns configuration value as floating number
func GetF(name string, defvals ...float64) float64 {
	return global.GetF(name, defvals...)
}

// GetB returns configuration value as boolean
func GetB(name string, defvals ...bool) bool {
	return global.GetB(name, defvals...)
}

// GetM returns configuration value as file mode
func GetM(name string, defvals ...os.FileMode) os.FileMode {
	return global.GetM(name, defvals...)
}

// GetD returns configuration values as duration
func GetD(name string, mod DurationMod, defvals ...time.Duration) time.Duration {
	return global.GetD(name, mod, defvals...)
}

// GetTD returns configuration value as time duration
func GetTD(name string, defvals ...time.Duration) time.Duration {
	return global.GetTD(name, defvals...)
}

// GetTS returns configuration timestamp value as time
func GetTS(name string, defvals ...time.Time) time.Time {
	return global.GetTS(name, defvals...)
}

// GetTS returns configuration value as timezone
func GetTZ(name string, defvals ...*time.Location) *time.Location {
	return global.GetTZ(name, defvals...)
}

// GetL returns configuration value as list
func GetL(name string, defvals ...[]string) []string {
	return global.GetL(name, defvals...)
}

// Validate executes all given validators and
// returns slice with validation errors
func Validate(validators []*knf.Validator) []error {
	if global == nil {
		return []error{knf.ErrNilConfig}
	}

	return validate(validators)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsEmpty returns true if mapping is empty
func (m Mapping) IsEmpty() bool {
	return m.Option == "" && m.Property == "" && m.Variable == ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetS returns configuration value as string
func (c *config) GetS(name string, defvals ...string) string {
	if c == nil {
		return knf.GetS(name, defvals...)
	}

	val := c.getProp(name)

	if val == "" && len(defvals) != 0 {
		return defvals[0]
	}

	return val
}

// GetI returns configuration value as int
func (c *config) GetI(name string, defvals ...int) int {
	if c == nil {
		return knf.GetI(name, defvals...)
	}

	return value.ParseInt(c.getProp(name), defvals...)
}

// GetI64 returns configuration value as int64
func (c *config) GetI64(name string, defvals ...int64) int64 {
	if c == nil {
		return knf.GetI64(name, defvals...)
	}

	return value.ParseInt64(c.getProp(name), defvals...)
}

// GetU returns configuration value as uint
func (c *config) GetU(name string, defvals ...uint) uint {
	if c == nil {
		return knf.GetU(name, defvals...)
	}

	return value.ParseUint(c.getProp(name), defvals...)
}

// GetU64 returns configuration value as uint64
func (c *config) GetU64(name string, defvals ...uint64) uint64 {
	if c == nil {
		return knf.GetU64(name, defvals...)
	}

	return value.ParseUint64(c.getProp(name), defvals...)
}

// GetF returns configuration value as floating number
func (c *config) GetF(name string, defvals ...float64) float64 {
	if c == nil {
		return knf.GetF(name, defvals...)
	}

	return value.ParseFloat(c.getProp(name), defvals...)
}

// GetB returns configuration value as boolean
func (c *config) GetB(name string, defvals ...bool) bool {
	if c == nil {
		return knf.GetB(name, defvals...)
	}

	return value.ParseBool(c.getProp(name), defvals...)
}

// GetM returns configuration value as file mode
func (c *config) GetM(name string, defvals ...os.FileMode) os.FileMode {
	if c == nil {
		return knf.GetM(name, defvals...)
	}

	return value.ParseMode(c.getProp(name), defvals...)
}

// GetD returns configuration values as duration
func (c *config) GetD(name string, mod DurationMod, defvals ...time.Duration) time.Duration {
	if c == nil {
		return knf.GetD(name, mod, defvals...)
	}

	return value.ParseDuration(c.getProp(name), time.Duration(mod), defvals...)
}

// GetTD returns configuration value as time duration
func (c *config) GetTD(name string, defvals ...time.Duration) time.Duration {
	if c == nil {
		return knf.GetTD(name, defvals...)
	}

	return value.ParseTimeDuration(c.getProp(name), defvals...)
}

// GetTS returns configuration timestamp value as time
func (c *config) GetTS(name string, defvals ...time.Time) time.Time {
	if c == nil {
		return knf.GetTS(name, defvals...)
	}

	return value.ParseTimestamp(c.getProp(name), defvals...)
}

// GetTS returns configuration value as timezone
func (c *config) GetTZ(name string, defvals ...*time.Location) *time.Location {
	if c == nil {
		return knf.GetTZ(name, defvals...)
	}

	return value.ParseTimezone(c.getProp(name), defvals...)
}

// GetL returns configuration value as list
func (c *config) GetL(name string, defvals ...[]string) []string {
	if c == nil {
		return knf.GetL(name, defvals...)
	}

	return value.ParseList(c.getProp(name), defvals...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (c *config) getProp(name string) string {
	m := c.mappings[name]

	if m.IsEmpty() {
		return knf.GetS(name)
	}

	switch {
	case m.Option != "" && optionHasFunc(m.Option):
		return optionGetFunc(m.Option)
	case m.Variable != "" && os.Getenv(m.Variable) != "":
		return os.Getenv(m.Variable)
	default:
		return knf.GetS(name)
	}
}

func validate(validators []*knf.Validator) []error {
	var result []error

	for _, v := range validators {
		err := v.Func(global, v.Property, v.Value)

		if err != nil {
			result = append(result, err)
		}
	}

	return result
}
