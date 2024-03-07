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
type Config struct {
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

var global *Config

var optionHasFunc = options.Has
var optionGetFunc = options.GetS

// ////////////////////////////////////////////////////////////////////////////////// //

// Combine applies mappings to combine knf properties, options, and environment
// variables
func Combine(mappings ...Mapping) error {
	config := &Config{
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
	if global == nil {
		return knf.GetS(name, defvals...)
	}

	val := global.getProp(name)

	if val == "" && len(defvals) != 0 {
		return defvals[0]
	}

	return val
}

// GetI returns configuration value as int
func GetI(name string, defvals ...int) int {
	if global == nil {
		return knf.GetI(name, defvals...)
	}

	return value.ParseInt(global.getProp(name), defvals...)
}

// GetI64 returns configuration value as int64
func GetI64(name string, defvals ...int64) int64 {
	if global == nil {
		return knf.GetI64(name, defvals...)
	}

	return value.ParseInt64(global.getProp(name), defvals...)
}

// GetU returns configuration value as uint
func GetU(name string, defvals ...uint) uint {
	if global == nil {
		return knf.GetU(name, defvals...)
	}

	return value.ParseUint(global.getProp(name), defvals...)
}

// GetU64 returns configuration value as uint64
func GetU64(name string, defvals ...uint64) uint64 {
	if global == nil {
		return knf.GetU64(name, defvals...)
	}

	return value.ParseUint64(global.getProp(name), defvals...)
}

// GetF returns configuration value as floating number
func GetF(name string, defvals ...float64) float64 {
	if global == nil {
		return knf.GetF(name, defvals...)
	}

	return value.ParseFloat(global.getProp(name), defvals...)
}

// GetB returns configuration value as boolean
func GetB(name string, defvals ...bool) bool {
	if global == nil {
		return knf.GetB(name, defvals...)
	}

	return value.ParseBool(global.getProp(name), defvals...)
}

// GetM returns configuration value as file mode
func GetM(name string, defvals ...os.FileMode) os.FileMode {
	if global == nil {
		return knf.GetM(name, defvals...)
	}

	return value.ParseMode(global.getProp(name), defvals...)
}

// GetD returns configuration values as duration
func GetD(name string, mod DurationMod, defvals ...time.Duration) time.Duration {
	if global == nil {
		return knf.GetD(name, mod, defvals...)
	}

	return value.ParseDuration(global.getProp(name), time.Duration(mod), defvals...)
}

// GetTD returns configuration value as time duration
func GetTD(name string, defvals ...time.Duration) time.Duration {
	if global == nil {
		return knf.GetTD(name, defvals...)
	}

	return value.ParseTimeDuration(global.getProp(name), defvals...)
}

// GetTS returns configuration timestamp value as time
func GetTS(name string, defvals ...time.Time) time.Time {
	if global == nil {
		return knf.GetTS(name, defvals...)
	}

	return value.ParseTimestamp(global.getProp(name), defvals...)
}

// GetTS returns configuration value as timezone
func GetTZ(name string, defvals ...*time.Location) *time.Location {
	if global == nil {
		return knf.GetTZ(name, defvals...)
	}

	return value.ParseTimezone(global.getProp(name), defvals...)
}

// GetL returns configuration value as list
func GetL(name string, defvals ...[]string) []string {
	if global == nil {
		return knf.GetL(name, defvals...)
	}

	return value.ParseList(global.getProp(name), defvals...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsEmpty returns true if mapping is empty
func (m Mapping) IsEmpty() bool {
	return m.Option == "" && m.Property == "" && m.Variable == ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (c *Config) getProp(name string) string {
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
