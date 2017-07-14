// Package knf provides methods for working with configs in KNF format
package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v9/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_COMMENT_SYMBOL   = "#"
	_SECTION_SYMBOL   = "["
	_SEPARATOR_SYMBOL = ":"
	_MACRO_SYMBOL     = "{"
	_DELIMITER        = ":"
)

const _MACRO_REGEXP = "{([a-zA-Z0-9_-]{2,}):([a-zA-Z0-9_-]{2,})}"

// ////////////////////////////////////////////////////////////////////////////////// //

// Config is basic config struct
type Config struct {
	sections []string
	props    []string
	data     map[string]string
	file     string
}

// Validator is config property validator struct
type Validator struct {
	Property string            // Property name
	Func     PropertyValidator // Validation function
	Value    interface{}       // Expected value
}

// ////////////////////////////////////////////////////////////////////////////////// //

// RegExp struct for searching and parsing macroses
var macroRE = regexp.MustCompile(_MACRO_REGEXP)

// Global config struct
var global *Config

// ////////////////////////////////////////////////////////////////////////////////// //

// Global read and parse config file
// Global config will be accessible globally from any part of the code
func Global(file string) error {
	config, err := Read(file)

	if err != nil {
		return err
	}

	global = config

	return nil
}

// Read reads and parse config file
func Read(file string) (*Config, error) {
	switch {
	case fsutil.IsExist(file) == false:
		return nil, errors.New("File " + file + " does not exist")
	case fsutil.IsReadable(file) == false:
		return nil, errors.New("File " + file + " is not readable")
	case fsutil.IsNonEmpty(file) == false:
		return nil, errors.New("File " + file + " is empty")
	}

	fd, err := os.OpenFile(path.Clean(file), os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	config := &Config{
		data: make(map[string]string),
		file: file,
	}

	err = readConfigData(config, fd, file)

	if err != nil {
		return nil, err
	}

	return config, nil
}

// Reload reads and parse global config file
func Reload() (map[string]bool, error) {
	if global == nil {
		return nil, errors.New("Global config is not loaded")
	}

	return global.Reload()
}

// GetS return global config value as string
func GetS(name string, defvals ...string) string {
	if global == nil {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	return global.GetS(name, defvals...)
}

// GetI return global config value as int
func GetI(name string, defvals ...int) int {
	if global == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return global.GetI(name, defvals...)
}

// GetI64 return global config value as int64
func GetI64(name string, defvals ...int64) int64 {
	if global == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return global.GetI64(name, defvals...)
}

// GetU return global config value as uint
func GetU(name string, defvals ...uint) uint {
	if global == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return global.GetU(name, defvals...)
}

// GetU64 return global config value as uint64
func GetU64(name string, defvals ...uint64) uint64 {
	if global == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	return global.GetU64(name, defvals...)
}

// GetF return global config value as floating number
func GetF(name string, defvals ...float64) float64 {
	if global == nil {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	return global.GetF(name, defvals...)
}

// GetB return global config value as boolean
func GetB(name string, defvals ...bool) bool {
	if global == nil {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	return global.GetB(name, defvals...)
}

// GetM return global config value as file mode
func GetM(name string, defvals ...os.FileMode) os.FileMode {
	if global == nil {
		if len(defvals) == 0 {
			return os.FileMode(0)
		}

		return defvals[0]
	}

	return global.GetM(name, defvals...)
}

// HasSection check if section exist
func HasSection(section string) bool {
	if global == nil {
		return false
	}

	return global.HasSection(section)
}

// HasProp check if property exist
func HasProp(name string) bool {
	if global == nil {
		return false
	}

	return global.HasProp(name)
}

// Sections return slice with section names
func Sections() []string {
	if global == nil {
		return []string{}
	}

	return global.Sections()
}

// Props return slice with properties names in some section
func Props(section string) []string {
	if global == nil {
		return []string{}
	}

	return global.Props(section)
}

// Validate require slice with pointers to validators and
// return slice with validation errors
func Validate(validators []*Validator) []error {
	if global == nil {
		return []error{errors.New("Global config struct is nil")}
	}

	return global.Validate(validators)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Reload read and parse config file
func (c *Config) Reload() (map[string]bool, error) {
	if c == nil {
		return nil, errors.New("Config is nil")
	}

	if c.file == "" {
		return nil, errors.New("Path to config file is empty (non initialized struct?)")
	}

	nc, err := Read(c.file)

	if err != nil {
		return nil, err
	}

	changes := make(map[string]bool)

	for prop, value := range c.data {
		changes[prop] = value != nc.data[prop]
	}

	c.data, c.sections = nc.data, nc.sections

	return changes, nil
}

// GetS return config value as string
func (c *Config) GetS(name string, defvals ...string) string {
	if c == nil {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	val := c.data[name]

	if val == "" {
		if len(defvals) == 0 {
			return ""
		}

		return defvals[0]
	}

	return val
}

// GetI64 return config value as int64
func (c *Config) GetI64(name string, defvals ...int64) int64 {
	if c == nil {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	val := c.data[name]

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

// GetI return config value as int
func (c *Config) GetI(name string, defvals ...int) int {
	if len(defvals) != 0 {
		return int(c.GetI64(name, int64(defvals[0])))
	}

	return int(c.GetI64(name))
}

// GetU return config value as uint
func (c *Config) GetU(name string, defvals ...uint) uint {
	if len(defvals) != 0 {
		return uint(c.GetI64(name, int64(defvals[0])))
	}

	return uint(c.GetI64(name))
}

// GetU64 return config value as uint64
func (c *Config) GetU64(name string, defvals ...uint64) uint64 {
	if len(defvals) != 0 {
		return uint64(c.GetI64(name, int64(defvals[0])))
	}

	return uint64(c.GetI64(name))
}

// GetF return config value as floating number
func (c *Config) GetF(name string, defvals ...float64) float64 {
	if c == nil {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	val := c.data[name]

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

// GetB return config value as boolean
func (c *Config) GetB(name string, defvals ...bool) bool {
	if c == nil {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	val := c.data[name]

	if val == "" {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	switch {
	case val == "", val == "0", val == "false":
		return false
	default:
		return true
	}
}

// GetM return config value as file mode
func (c *Config) GetM(name string, defvals ...os.FileMode) os.FileMode {
	if c == nil {
		if len(defvals) == 0 {
			return os.FileMode(0)
		}

		return defvals[0]
	}

	val := c.data[name]

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

// HasSection check if section exist
func (c *Config) HasSection(section string) bool {
	if c == nil {
		return false
	}

	return c.data[section+_DELIMITER] == "true"
}

// HasProp check if property exist
func (c *Config) HasProp(name string) bool {
	if c == nil {
		return false
	}

	return c.data[name] != ""
}

// Sections return slice with section names
func (c *Config) Sections() []string {
	if c == nil {
		return []string{}
	}

	return c.sections
}

// Props return slice with properties names in some section
func (c *Config) Props(section string) []string {
	if c == nil || !c.HasSection(section) {
		return []string{}
	}

	var result []string

	// Section name + delimiter
	snLength := len(section) + 1

	for _, prop := range c.props {
		if len(prop) <= snLength {
			continue
		}

		if prop[0:snLength] == section+_DELIMITER {
			result = append(result, prop[snLength:])
		}
	}

	return result
}

// Validate require slice with pointers to validators and
// return slice with validation errors
func (c *Config) Validate(validators []*Validator) []error {
	if c == nil {
		return []error{errors.New("Config is nil")}
	}

	var result []error

	for _, v := range validators {
		err := v.Func(c, v.Property, v.Value)

		if err != nil {
			result = append(result, err)
		}
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

func readConfigData(config *Config, fd io.Reader, file string) error {
	var sectionName = ""

	reader := bufio.NewReader(fd)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.Trim(line, " \t") == "" {
			continue
		}

		if strings.HasPrefix(strings.TrimLeft(line, " \t"), _COMMENT_SYMBOL) {
			continue
		}

		if strings.HasPrefix(strings.TrimLeft(line, " \t"), _SECTION_SYMBOL) {
			sectionName = strings.Trim(line, "[ ]")
			config.data[sectionName+_DELIMITER] = "true"
			config.sections = append(config.sections, sectionName)
			continue
		}

		if sectionName == "" {
			return errors.New("Configuration file " + file + " is malformed")
		}

		propName, propValue := parseRecord(line, config)
		fullPropName := sectionName + _DELIMITER + propName

		config.props = append(config.props, fullPropName)
		config.data[fullPropName] = propValue
	}

	return scanner.Err()
}

func parseRecord(data string, config *Config) (string, string) {
	va := strings.Split(data, _SEPARATOR_SYMBOL)

	propName := va[0]
	propValue := strings.Join(va[1:], _SEPARATOR_SYMBOL)

	propName = strings.TrimLeft(propName, " \t")
	propValue = strings.TrimLeft(propValue, " \t")
	propValue = strings.TrimRight(propValue, " ")

	if strings.Contains(propValue, _MACRO_SYMBOL) {
		macroses := macroRE.FindAllStringSubmatch(propValue, -1)

		for _, macros := range macroses {
			macroFull := macros[0]
			macroSect := macros[1]
			macroProp := macros[2]

			macroVal := config.GetS(macroSect + _DELIMITER + macroProp)

			propValue = strings.Replace(propValue, macroFull, macroVal, -1)
		}
	}

	return propName, propValue
}

// ////////////////////////////////////////////////////////////////////////////////// //
