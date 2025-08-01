// Package knf provides methods for working with configuration files in KNF format
package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/essentialkaos/ek/v13/errors"
	"github.com/essentialkaos/ek/v13/knf/value"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	MILLISECOND = DurationMod(time.Millisecond)
	SECOND      = 1000 * MILLISECOND
	MINUTE      = 60 * SECOND
	HOUR        = 60 * MINUTE
	DAY         = 24 * HOUR
	WEEK        = 7 * DAY
)

// ////////////////////////////////////////////////////////////////////////////////// //

// IConfig is knf like configuration
type IConfig interface {
	// GetS returns configuration value as string
	GetS(name string, defvals ...string) string

	// GetI returns configuration value as int
	GetI(name string, defvals ...int) int

	// GetI64 returns configuration value as int64
	GetI64(name string, defvals ...int64) int64

	// GetU returns configuration value as uint
	GetU(name string, defvals ...uint) uint

	// GetU64 returns configuration value as uint64
	GetU64(name string, defvals ...uint64) uint64

	// GetF returns configuration value as floating number
	GetF(name string, defvals ...float64) float64

	// GetB returns configuration value as boolean
	GetB(name string, defvals ...bool) bool

	// GetM returns configuration value as file mode
	GetM(name string, defvals ...os.FileMode) os.FileMode

	// GetD returns configuration value as duration
	GetD(name string, mod DurationMod, defvals ...time.Duration) time.Duration

	// GetSZ returns configuration value as a size in bytes
	GetSZ(name string, defvals ...uint64) uint64

	// GetTD returns configuration value as time duration
	GetTD(name string, defvals ...time.Duration) time.Duration

	// GetTS returns configuration timestamp value as time
	GetTS(name string, defvals ...time.Time) time.Time

	// GetTS returns configuration value as timezone
	GetTZ(name string, defvals ...*time.Location) *time.Location

	// GetL returns configuration value as list
	GetL(name string, defvals ...[]string) []string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Config is basic configuration instance
type Config struct {
	sections []string
	props    []string
	data     map[string]string
	aliases  map[string]string
	file     string

	mx *sync.RWMutex
}

// Validator is configuration property validator struct
type Validator struct {
	Property string            // Property name
	Func     PropertyValidator // Validation function
	Value    any               // Expected value
}

// Validators is a slice with validators
type Validators []*Validator

// PropertyValidator is default type of property validation function
type PropertyValidator func(config IConfig, prop string, value any) error

// DurationMod is type for duration modifier
type DurationMod int64

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrNilConfig  = errors.New("Configuration is nil")
	ErrCantReload = errors.New("Can't reload configuration file: path to file is empty")
	ErrCantMerge  = errors.New("Can't merge configurations: given configuration is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// global is global configuration file
var global *Config

// ////////////////////////////////////////////////////////////////////////////////// //

// Global reads and parses configuration file
// Global instance is accessible from any part of the code
func Global(file string) error {
	config, err := Read(file)

	if err != nil {
		return err
	}

	global = config

	return nil
}

// Read reads and parses configuration file
func Read(file string) (*Config, error) {
	fd, err := os.OpenFile(path.Clean(file), os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	config, err := readData(fd)

	if err != nil {
		return nil, err
	}

	config.file = file

	return config, nil
}

// Parse parses data with KNF configuration
func Parse(data []byte) (*Config, error) {
	return readData(bytes.NewBuffer(data))
}

// Reload reloads global configuration file
func Reload() (map[string]bool, error) {
	if global == nil {
		return nil, ErrNilConfig
	}

	return global.Reload()
}

// Alias creates alias for configuration property
//
// It's useful for refactoring the configuration or for providing support for
// renamed properties
func Alias(old, new string) error {
	if global == nil {
		return ErrNilConfig
	}

	return global.Alias(old, new)
}

// GetS returns configuration value as string
func GetS(name string, defvals ...string) string {
	if global == nil {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	return global.GetS(name, defvals...)
}

// GetI returns configuration value as int
func GetI(name string, defvals ...int) int {
	if global == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return global.GetI(name, defvals...)
}

// GetI64 returns configuration value as int64
func GetI64(name string, defvals ...int64) int64 {
	if global == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return global.GetI64(name, defvals...)
}

// GetU returns configuration value as uint
func GetU(name string, defvals ...uint) uint {
	if global == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return global.GetU(name, defvals...)
}

// GetU64 returns configuration value as uint64
func GetU64(name string, defvals ...uint64) uint64 {
	if global == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return global.GetU64(name, defvals...)
}

// GetF returns configuration value as floating number
func GetF(name string, defvals ...float64) float64 {
	if global == nil {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	return global.GetF(name, defvals...)
}

// GetB returns configuration value as boolean
func GetB(name string, defvals ...bool) bool {
	if global == nil {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	return global.GetB(name, defvals...)
}

// GetM returns configuration value as file mode
func GetM(name string, defvals ...os.FileMode) os.FileMode {
	if global == nil {
		if len(defvals) == 0 {
			return os.FileMode(0)
		}

		return defvals[0]
	}

	return global.GetM(name, defvals...)
}

// GetD returns configuration values as duration
func GetD(name string, mod DurationMod, defvals ...time.Duration) time.Duration {
	if global == nil {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	return global.GetD(name, mod, defvals...)
}

// GetSZ returns configuration value as a size in bytes
func GetSZ(name string, defvals ...uint64) uint64 {
	if global == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return global.GetSZ(name, defvals...)
}

// GetTD returns configuration value as time duration
func GetTD(name string, defvals ...time.Duration) time.Duration {
	if global == nil {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	return global.GetTD(name, defvals...)
}

// GetTS returns configuration timestamp value as time
func GetTS(name string, defvals ...time.Time) time.Time {
	if global == nil {
		if len(defvals) == 0 {
			return time.Time{}
		}

		return defvals[0]
	}

	return global.GetTS(name, defvals...)
}

// GetTS returns configuration value as timezone
func GetTZ(name string, defvals ...*time.Location) *time.Location {
	if global == nil {
		if len(defvals) == 0 {
			return nil
		}

		return defvals[0]
	}

	return global.GetTZ(name, defvals...)
}

// GetL returns configuration value as list
func GetL(name string, defvals ...[]string) []string {
	if global == nil {
		if len(defvals) == 0 {
			return nil
		}

		return defvals[0]
	}

	return global.GetL(name, defvals...)
}

// Is checks if given property contains given value
func Is(name string, value any) bool {
	if global == nil {
		return false
	}

	return global.Is(name, value)
}

// HasSection checks if the section exists
func HasSection(section string) bool {
	if global == nil {
		return false
	}

	return global.HasSection(section)
}

// Has checks if the property is defined and set
func Has(name string) bool {
	if global == nil {
		return false
	}

	return global.Has(name)
}

// Sections returns slice with section names
func Sections() []string {
	if global == nil {
		return nil
	}

	return global.Sections()
}

// Props returns slice with properties names in some section
func Props(section string) []string {
	if global == nil {
		return nil
	}

	return global.Props(section)
}

// Validate executes all given validators and
// returns slice with validation errors
func Validate(validators Validators) errors.Errors {
	if global == nil {
		return errors.Errors{ErrNilConfig}
	}

	return global.Validate(validators)
}

// Q is a helper to create a valid full property name (section + delimiter
// + property name)
func Q(section, prop string) string {
	return strings.ToLower(section + _SYMBOL_DELIMITER + prop)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Add adds given validators and returns new slice
func (v Validators) Add(validators Validators) Validators {
	return append(v, validators...)
}

// AddIf adds given validators if conditional is true
func (v Validators) AddIf(cond bool, validators Validators) Validators {
	if !cond {
		return v
	}

	return v.Add(validators)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Merge merges two configurations
func (c *Config) Merge(cfg *Config) error {
	if c == nil || c.mx == nil {
		return ErrNilConfig
	}

	if cfg == nil || cfg.mx == nil {
		return ErrCantMerge
	}

	for k, v := range cfg.data {
		c.data[k] = v
	}

SECTION_LOOP:
	for _, ss := range cfg.sections {
		for _, ts := range c.sections {
			if ss == ts {
				continue SECTION_LOOP
			}
		}

		c.sections = append(c.sections, ss)
	}

PROP_LOOP:
	for _, sp := range cfg.props {
		for _, tp := range c.props {
			if sp == tp {
				continue PROP_LOOP
			}
		}

		c.props = append(c.props, sp)
	}

	return nil
}

// Reload reloads configuration file
func (c *Config) Reload() (map[string]bool, error) {
	if c == nil || c.mx == nil {
		return nil, ErrNilConfig
	}

	if c.file == "" {
		return nil, ErrCantReload
	}

	nc, err := Read(c.file)

	if err != nil {
		return nil, err
	}

	changes := make(map[string]bool)

	c.mx.RLock()

	for _, prop := range c.props {
		changes[prop] = c.GetS(prop) != nc.GetS(prop)
	}

	c.mx.RUnlock()
	c.mx.Lock()

	// Update current config data
	c.data, c.sections, c.props = nc.data, nc.sections, nc.props

	c.mx.Unlock()
	return changes, nil
}

// Alias creates alias for configuration property
//
// It's useful for refactoring the configuration or for providing support for
// renamed properties
func (c *Config) Alias(old, new string) error {
	if c == nil || c.mx == nil {
		return ErrNilConfig
	}

	switch {
	case old == "":
		return fmt.Errorf("Old property name is empty")
	case new == "":
		return fmt.Errorf("New property name is empty")
	case !isValidPropName(old):
		return fmt.Errorf("Old property name (%q) is invalid", old)
	case !isValidPropName(new):
		return fmt.Errorf("New property name (%q) is invalid", new)
	}

	c.mx.Lock()

	if c.aliases == nil {
		c.aliases = make(map[string]string)
	}

	c.aliases[strings.ToLower(new)] = strings.ToLower(old)

	c.mx.Unlock()

	return nil
}

// GetS returns configuration value as string
func (c *Config) GetS(name string, defvals ...string) string {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	val := c.getValue(name)

	if val == "" {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	return val
}

// GetI returns configuration value as int
func (c *Config) GetI(name string, defvals ...int) int {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseInt(c.getValue(name), defvals...)
}

// GetI64 returns configuration value as int64
func (c *Config) GetI64(name string, defvals ...int64) int64 {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseInt64(c.getValue(name), defvals...)
}

// GetU returns configuration value as uint
func (c *Config) GetU(name string, defvals ...uint) uint {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseUint(c.getValue(name), defvals...)
}

// GetU64 returns configuration value as uint64
func (c *Config) GetU64(name string, defvals ...uint64) uint64 {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseUint64(c.getValue(name), defvals...)
}

// GetF returns configuration value as floating number
func (c *Config) GetF(name string, defvals ...float64) float64 {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	return value.ParseFloat(c.getValue(name), defvals...)
}

// GetB returns configuration value as boolean
func (c *Config) GetB(name string, defvals ...bool) bool {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	return value.ParseBool(c.getValue(name), defvals...)
}

// GetM returns configuration value as file mode
func (c *Config) GetM(name string, defvals ...os.FileMode) os.FileMode {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return os.FileMode(0)
		}

		return defvals[0]
	}

	return value.ParseMode(c.getValue(name), defvals...)
}

// GetD returns configuration value as duration
func (c *Config) GetD(name string, mod DurationMod, defvals ...time.Duration) time.Duration {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	return value.ParseDuration(c.getValue(name), time.Duration(mod), defvals...)
}

// GetSZ returns configuration value as a size in bytes
func (c *Config) GetSZ(name string, defvals ...uint64) uint64 {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return value.ParseSize(c.getValue(name), defvals...)
}

// GetTD returns configuration value as time duration (s/m/h/d/w)
func (c *Config) GetTD(name string, defvals ...time.Duration) time.Duration {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	return value.ParseTimeDuration(c.getValue(name), defvals...)
}

// GetTS returns configuration timestamp value as time
func (c *Config) GetTS(name string, defvals ...time.Time) time.Time {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return time.Time{}
		}

		return defvals[0]
	}

	return value.ParseTimestamp(c.getValue(name), defvals...)
}

// GetTS returns configuration value as timezone
func (c *Config) GetTZ(name string, defvals ...*time.Location) *time.Location {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return nil
		}

		return defvals[0]
	}

	return value.ParseTimezone(c.getValue(name), defvals...)
}

// GetL returns configuration value as list
func (c *Config) GetL(name string, defvals ...[]string) []string {
	if c == nil || c.mx == nil || !isValidPropName(name) {
		if len(defvals) == 0 {
			return nil
		}

		return defvals[0]
	}

	return value.ParseList(c.getValue(name), defvals...)
}

// Is checks if given property contains given value
func (c *Config) Is(name string, value any) bool {
	if c == nil || c.mx == nil || !isValidPropName(name) {
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

// HasSection checks if section exists
func (c *Config) HasSection(section string) bool {
	if c == nil || c.mx == nil {
		return false
	}

	c.mx.RLock()
	defer c.mx.RUnlock()

	// The "section" variable contains an invalid name for a property, so the user
	// can't read the value as a property, but we can store information about
	// sections.
	return c.data[strings.ToLower(section)] == "!"
}

// Has checks if property is defined and set
func (c *Config) Has(name string) bool {
	if c == nil || c.mx == nil {
		return false
	}

	return c.getValue(name) != ""
}

// Sections returns slice with section names
func (c *Config) Sections() []string {
	if c == nil || c.mx == nil {
		return nil
	}

	c.mx.RLock()
	defer c.mx.RUnlock()

	return c.sections
}

// Props returns slice with properties names in some section
func (c *Config) Props(section string) []string {
	if c == nil || !c.HasSection(section) {
		return nil
	}

	var result []string

	// Section name + delimiter
	snLength := len(section) + 1

	c.mx.RLock()

	for _, prop := range c.props {
		if len(prop) <= snLength {
			continue
		}

		if prop[:snLength] == section+_SYMBOL_DELIMITER {
			result = append(result, prop[snLength:])
		}
	}

	defer c.mx.RUnlock()

	return result
}

// File returns path to configuration file
func (c *Config) File() string {
	if c == nil || c.mx == nil {
		return ""
	}

	return c.file
}

// Validate executes all given validators and
// returns slice with validation errors
func (c *Config) Validate(validators Validators) errors.Errors {
	if c == nil || c.mx == nil {
		return errors.Errors{ErrNilConfig}
	}

	var errs errors.Errors

	c.mx.RLock()

	for _, v := range validators {
		err := v.Func(c, v.Property, v.Value)

		if err != nil {
			errs = append(errs, err)
		}
	}

	defer c.mx.RUnlock()

	return errs
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getValue returns property value from the storage
func (c *Config) getValue(propName string) string {
	if c == nil || c.mx == nil {
		return ""
	}

	c.mx.RLock()
	defer c.mx.RUnlock()

	propName = strings.ToLower(propName)

	if c.aliases != nil && c.aliases[propName] != "" {
		if c.data[c.aliases[propName]] != "" {
			return c.data[c.aliases[propName]]
		}
	}

	return c.data[propName]
}

// ////////////////////////////////////////////////////////////////////////////////// //

// isValidPropName returns true if property name is valid
func isValidPropName(propName string) bool {
	section, prop, ok := strings.Cut(propName, _SYMBOL_DELIMITER)

	switch {
	case !ok,
		strings.Trim(section, " ") == "",
		strings.Trim(prop, " ") == "",
		strings.Count(propName, _SYMBOL_DELIMITER) > 1:
		return false
	}

	return true
}
