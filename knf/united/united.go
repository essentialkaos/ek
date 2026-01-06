// Package united provides KNF configuration extended by environment variables and options
package united

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/errors"
	"github.com/essentialkaos/ek/v13/knf"
	"github.com/essentialkaos/ek/v13/knf/value"
	"github.com/essentialkaos/ek/v13/options"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	MILLISECOND = knf.MILLISECOND
	SECOND      = knf.SECOND
	MINUTE      = knf.MINUTE
	HOUR        = knf.HOUR
	DAY         = knf.DAY
	WEEK        = knf.WEEK
)

// ////////////////////////////////////////////////////////////////////////////////// //

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

// united is united configuration wrapper
type united struct {
	knf      *knf.Config
	mappings map[string]Mapping
	env      map[string]string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var global *united

var optionHasFunc = options.Has
var optionGetFunc = options.GetS

// ////////////////////////////////////////////////////////////////////////////////// //

// Combine applies mappings to combine knf properties, options, and environment
// variables
//
// Note that the environment variable will be moved to config after combining (e.g.
// won't be accessible with os.Getenv)
func Combine(config *knf.Config, mappings ...Mapping) error {
	if config == nil {
		return knf.ErrNilConfig
	}

	global = &united{
		knf:      config,
		mappings: make(map[string]Mapping),
		env:      make(map[string]string),
	}

	for _, m := range mappings {
		global.mappings[m.Property] = m

		if m.Variable != "" {
			global.env[m.Variable] = os.Getenv(m.Variable)
			os.Setenv(m.Variable, "")
		}
	}

	return nil
}

// CombineSimple applies mappings to combine knf properties, options, and environment
// variables. This method creates simple mappings based on properties names.
//
// Note that the environment variable will be moved to config after combining (e.g.
// won't be accessible with os.Getenv)
func CombineSimple(config *knf.Config, props ...string) error {
	if config == nil {
		return knf.ErrNilConfig
	}

	global = &united{
		knf:      config,
		mappings: make(map[string]Mapping),
		env:      make(map[string]string),
	}

	for _, p := range props {
		m := Simple(p)

		global.mappings[m.Property] = m
		global.env[m.Variable] = os.Getenv(m.Variable)

		os.Setenv(m.Variable, "")
	}

	return nil
}

// AddOptions adds options with knf properties to map
func AddOptions(m options.Map, names ...string) {
	if m == nil {
		return
	}

	for _, n := range names {
		m.Set(ToOption(n), &options.V{})
	}
}

// GetMapping returns mapping info for given property
func GetMapping(prop string) Mapping {
	if global == nil || global.mappings == nil {
		return Mapping{}
	}

	return global.mappings[prop]
}

// Simple creates simple mapping for knf property
// section:property → --section-property + SECTION_PROPERTY
func Simple(name string) Mapping {
	return Mapping{name, O(name), E(name)}
}

// ToOption converts knf property name to option name
func ToOption(name string) string {
	return strutil.ReplaceAll(strings.ToLower(name), "_:", "-")
}

// ToEnvVar converts knf property name to environment variable name
func ToEnvVar(name string) string {
	return strutil.ReplaceAll(strings.ToUpper(name), "-:", "_")
}

// O is a shortcut for [ToOption]
func O(name string) string {
	return ToOption(name)
}

// E is a shortcut for [ToEnvVar]
func E(name string) string {
	return ToEnvVar(name)
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

// GetSZ returns configuration value as a size in bytes
func GetSZ(name string, defvals ...uint64) uint64 {
	return global.GetSZ(name, defvals...)
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

// Is checks if given property contains given value
func Is(name string, value any) bool {
	return global.Is(name, value)
}

// Has checks if the property is defined and set
func Has(name string) bool {
	return global.Has(name)
}

// Validate executes all given validators and
// returns slice with validation errors
func Validate(validators knf.Validators) errors.Errors {
	if global == nil {
		return errors.Errors{knf.ErrNilConfig}
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
func (c *united) GetS(name string, defvals ...string) string {
	if c == nil {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	val := c.getProp(name)

	if val == "" && len(defvals) != 0 {
		return defvals[0]
	}

	return val
}

// GetI returns configuration value as int
func (c *united) GetI(name string, defvals ...int) int {
	if c == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseInt(c.getProp(name), defvals...)
}

// GetI64 returns configuration value as int64
func (c *united) GetI64(name string, defvals ...int64) int64 {
	if c == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseInt64(c.getProp(name), defvals...)
}

// GetU returns configuration value as uint
func (c *united) GetU(name string, defvals ...uint) uint {
	if c == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseUint(c.getProp(name), defvals...)
}

// GetU64 returns configuration value as uint64
func (c *united) GetU64(name string, defvals ...uint64) uint64 {
	if c == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseUint64(c.getProp(name), defvals...)
}

// GetF returns configuration value as floating number
func (c *united) GetF(name string, defvals ...float64) float64 {
	if c == nil {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	return value.ParseFloat(c.getProp(name), defvals...)
}

// GetB returns configuration value as boolean
func (c *united) GetB(name string, defvals ...bool) bool {
	if c == nil {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	return value.ParseBool(c.getProp(name), defvals...)
}

// GetM returns configuration value as file mode
func (c *united) GetM(name string, defvals ...os.FileMode) os.FileMode {
	if c == nil {
		if len(defvals) == 0 {
			return os.FileMode(0)
		}

		return defvals[0]
	}

	return value.ParseMode(c.getProp(name), defvals...)
}

// GetD returns configuration values as duration
func (c *united) GetD(name string, mod DurationMod, defvals ...time.Duration) time.Duration {
	if c == nil {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	return value.ParseDuration(c.getProp(name), time.Duration(mod), defvals...)
}

// GetSZ returns configuration value as a size in bytes
func (c *united) GetSZ(name string, defvals ...uint64) uint64 {
	if c == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseSize(c.getProp(name), defvals...)
}

// GetTD returns configuration value as time duration
func (c *united) GetTD(name string, defvals ...time.Duration) time.Duration {
	if c == nil {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	return value.ParseTimeDuration(c.getProp(name), defvals...)
}

// GetTS returns configuration timestamp value as time
func (c *united) GetTS(name string, defvals ...time.Time) time.Time {
	if c == nil {
		if len(defvals) == 0 {
			return time.Time{}
		}

		return defvals[0]
	}

	return value.ParseTimestamp(c.getProp(name), defvals...)
}

// GetTS returns configuration value as timezone
func (c *united) GetTZ(name string, defvals ...*time.Location) *time.Location {
	if c == nil {
		if len(defvals) == 0 {
			return nil
		}

		return defvals[0]
	}

	return value.ParseTimezone(c.getProp(name), defvals...)
}

// GetL returns configuration value as list
func (c *united) GetL(name string, defvals ...[]string) []string {
	if c == nil {
		if len(defvals) == 0 {
			return nil
		}

		return defvals[0]
	}

	return value.ParseList(c.getProp(name), defvals...)
}

// Is checks if given property contains given value
func (c *united) Is(name string, value any) bool {
	if c == nil {
		return false
	}

	switch t := value.(type) {
	case string:
		return c.GetS(name) == t
	case int:
		return c.GetI(name) == t
	case int64:
		return c.GetI64(name) == t
	case uint:
		return c.GetU(name) == t
	case uint64:
		return c.GetU64(name) == t
	case float64:
		return c.GetF(name) == t
	case bool:
		return c.GetB(name) == t
	case os.FileMode:
		return c.GetM(name) == t
	case time.Duration:
		return c.GetD(name, SECOND) == t
	case time.Time:
		return c.GetTS(name).Unix() == t.Unix()
	case *time.Location:
		return fmt.Sprint(c.GetTZ(name)) == fmt.Sprint(t)
	case []string:
		return slices.Equal(c.GetL(name), t)
	}

	return false
}

// Has checks if property is defined and set
func (c *united) Has(name string) bool {
	if c == nil {
		return false
	}

	return c.getProp(name) != ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getProp returns property value for knf configuration, env vars, or options
func (c *united) getProp(name string) string {
	m := c.mappings[name]

	if m.IsEmpty() {
		return c.knf.GetS(name)
	}

	switch {
	case m.Option != "" && optionHasFunc(m.Option):
		return optionGetFunc(m.Option)
	case m.Variable != "" && c.env[m.Variable] != "":
		return c.env[m.Variable]
	default:
		return c.knf.GetS(name)
	}
}

// validate runs validators over configuration
func validate(validators knf.Validators) errors.Errors {
	var result errors.Errors

	for _, v := range validators {
		err := v.Func(global, v.Property, v.Value)

		if err != nil {
			result = append(result, err)
		}
	}

	return result
}
