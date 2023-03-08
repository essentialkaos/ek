// Package knf provides methods for working with configuration files in KNF format
package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Config is basic config struct
type Config struct {
	sections []string
	props    []string
	data     map[string]string
	file     string

	mx *sync.RWMutex
}

// Validator is config property validator struct
type Validator struct {
	Property string            // Property name
	Func     PropertyValidator // Validation function
	Value    any               // Expected value
}

// PropertyValidator is default type of property validation function
type PropertyValidator func(config *Config, prop string, value any) error

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrNilConfig  = errors.New("Config is nil")
	ErrFileNotSet = errors.New("Path to config file is empty (non initialized struct?)")
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
	return readKNFFile(file)
}

// Reload reloads global configuration file
func Reload() (map[string]bool, error) {
	if global == nil {
		return nil, ErrNilConfig
	}

	return global.Reload()
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
func GetD(name string, defvals ...time.Duration) time.Duration {
	if global == nil {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	return global.GetD(name, defvals...)
}

// Is checks if given property contains given value
func Is(name string, value any) bool {
	if global == nil {
		return false
	}

	return global.Is(name, value)
}

// HasSection checks if section exist
func HasSection(section string) bool {
	if global == nil {
		return false
	}

	return global.HasSection(section)
}

// HasProp checks if property exist
func HasProp(name string) bool {
	if global == nil {
		return false
	}

	return global.HasProp(name)
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
func Validate(validators []*Validator) []error {
	if global == nil {
		return []error{ErrNilConfig}
	}

	return global.Validate(validators)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Reload reloads configuration file
func (c *Config) Reload() (map[string]bool, error) {
	if c == nil || c.mx == nil {
		return nil, ErrNilConfig
	}

	if c.file == "" {
		return nil, ErrFileNotSet
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

// GetS returns configuration value as string
func (c *Config) GetS(name string, defvals ...string) string {
	if c == nil || c.mx == nil {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	c.mx.RLock()
	val := c.data[strings.ToLower(name)]
	c.mx.RUnlock()

	if val == "" {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	return val
}

// GetI64 returns configuration value as int64
func (c *Config) GetI64(name string, defvals ...int64) int64 {
	if c == nil || c.mx == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	c.mx.RLock()
	val := c.data[strings.ToLower(name)]
	c.mx.RUnlock()

	if val == "" {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	// HEX Parsing
	if len(val) >= 3 && val[0:2] == "0x" {
		valHex, err := strconv.ParseInt(val[2:], 16, 0)

		if err != nil {
			return 0
		}

		return valHex
	}

	valInt, err := strconv.ParseInt(val, 10, 0)

	if err != nil {
		return 0
	}

	return valInt
}

// GetI returns configuration value as int
func (c *Config) GetI(name string, defvals ...int) int {
	if len(defvals) != 0 {
		return int(c.GetI64(name, int64(defvals[0])))
	}

	return int(c.GetI64(name))
}

// GetU returns configuration value as uint
func (c *Config) GetU(name string, defvals ...uint) uint {
	if len(defvals) != 0 {
		return uint(c.GetI64(name, int64(defvals[0])))
	}

	return uint(c.GetI64(name))
}

// GetU64 returns configuration value as uint64
func (c *Config) GetU64(name string, defvals ...uint64) uint64 {
	if len(defvals) != 0 {
		return uint64(c.GetI64(name, int64(defvals[0])))
	}

	return uint64(c.GetI64(name))
}

// GetF returns configuration value as floating number
func (c *Config) GetF(name string, defvals ...float64) float64 {
	if c == nil || c.mx == nil {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	c.mx.RLock()
	val := c.data[strings.ToLower(name)]
	c.mx.RUnlock()

	if val == "" {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	valFl, err := strconv.ParseFloat(val, 64)

	if err != nil {
		return 0.0
	}

	return valFl
}

// GetB returns configuration value as boolean
func (c *Config) GetB(name string, defvals ...bool) bool {
	if c == nil || c.mx == nil {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	c.mx.RLock()
	val := c.data[strings.ToLower(name)]
	c.mx.RUnlock()

	if val == "" {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	switch val {
	case "", "0", "false", "no":
		return false
	default:
		return true
	}
}

// GetM returns configuration value as file mode
func (c *Config) GetM(name string, defvals ...os.FileMode) os.FileMode {
	if c == nil || c.mx == nil {
		if len(defvals) == 0 {
			return os.FileMode(0)
		}

		return defvals[0]
	}

	c.mx.RLock()
	val := c.data[strings.ToLower(name)]
	c.mx.RUnlock()

	if val == "" {
		if len(defvals) == 0 {
			return os.FileMode(0)
		}

		return defvals[0]
	}

	valM, err := strconv.ParseUint(val, 8, 32)

	if err != nil {
		return 0
	}

	return os.FileMode(valM)
}

// GetD returns configuration value as duration
func (c *Config) GetD(name string, defvals ...time.Duration) time.Duration {
	if c == nil || c.mx == nil {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	c.mx.RLock()
	val := c.data[strings.ToLower(name)]
	c.mx.RUnlock()

	if val == "" {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	return time.Duration(c.GetI64(name)) * time.Second
}

// Is checks if given property contains given value
func (c *Config) Is(name string, value any) bool {
	if c == nil || c.mx == nil {
		return false
	}

	if !c.HasProp(name) {
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
		return c.GetD(name) == t
	}

	return false
}

// HasSection checks if section exist
func (c *Config) HasSection(section string) bool {
	if c == nil || c.mx == nil {
		return false
	}

	c.mx.RLock()
	defer c.mx.RUnlock()

	return c.data[strings.ToLower(section)] == "!"
}

// HasProp checks if property exist
func (c *Config) HasProp(name string) bool {
	if c == nil || c.mx == nil {
		return false
	}

	c.mx.RLock()
	defer c.mx.RUnlock()

	return c.data[strings.ToLower(name)] != ""
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
	var result []string

	if c == nil || !c.HasSection(section) {
		return result
	}

	// Section name + delimiter
	snLength := len(section) + 1

	c.mx.RLock()

	for _, prop := range c.props {
		if len(prop) <= snLength {
			continue
		}

		if prop[:snLength] == section+_PROP_DELIMITER {
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
func (c *Config) Validate(validators []*Validator) []error {
	if c == nil || c.mx == nil {
		return []error{ErrNilConfig}
	}

	var result []error

	c.mx.RLock()

	for _, v := range validators {
		err := v.Func(c, v.Property, v.Value)

		if err != nil {
			result = append(result, err)
		}
	}

	defer c.mx.RUnlock()

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //
